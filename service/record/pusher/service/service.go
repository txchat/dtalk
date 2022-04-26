package service

import (
	"context"
	"errors"
	"fmt"
	device "github.com/txchat/dtalk/service/device/api"
	"reflect"
	"runtime"
	"time"

	"github.com/Shopify/sarama"
	"github.com/gammazero/workerpool"
	"github.com/golang/protobuf/proto"
	"github.com/opentracing/opentracing-go"
	traceLog "github.com/opentracing/opentracing-go/log"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/txchat/dtalk/pkg/naming"
	"github.com/txchat/dtalk/pkg/net/grpc"
	"github.com/txchat/dtalk/pkg/util"
	answer "github.com/txchat/dtalk/service/record/answer/api"
	"github.com/txchat/dtalk/service/record/kafka/consumer"
	record "github.com/txchat/dtalk/service/record/proto"
	"github.com/txchat/dtalk/service/record/pusher/config"
	"github.com/txchat/dtalk/service/record/pusher/dao"
	"github.com/txchat/dtalk/service/record/pusher/logH"
	"github.com/txchat/dtalk/service/record/pusher/model"
	"github.com/txchat/im-pkg/trace"
	comet "github.com/txchat/im/api/comet/grpc"
	logic "github.com/txchat/im/api/logic/grpc"
	"github.com/txchat/imparse"
	"github.com/txchat/imparse/chat"
	xproto "github.com/txchat/imparse/proto"
	"github.com/uber/jaeger-client-go"
	"google.golang.org/grpc/resolver"
)

type Service struct {
	log               zerolog.Logger
	dao               *dao.Dao
	lh                *logH.LogHelper
	cfg               *config.Config
	connConsumers     map[string]*consumer.Consumer
	receivedConsumers map[string]*consumer.Consumer
	logicClient       logic.LogicClient

	pusher Pusher

	answerClient *answer.Client
	deviceClient *device.Client

	workPoolConn *workerpool.WorkerPool
	workPoolSend *workerpool.WorkerPool
}

func New(c *config.Config) *Service {
	s := &Service{
		log: log.Logger,
		dao: dao.New(c),
		lh:  logH.NewLogHelper(dao.New(c)),
		cfg: c,
		connConsumers: consumer.NewKafkaConsumers(
			fmt.Sprintf("goim-%s-topic", c.AppId),
			fmt.Sprintf("goim-%s-conn-group", c.AppId), c.IMSub.Brokers, c.IMSub.Number),
		receivedConsumers: consumer.NewKafkaConsumers(
			fmt.Sprintf("received-%s-topic", c.AppId),
			fmt.Sprintf("received-%s-group", c.AppId), c.RevSub.Brokers, c.RevSub.Number),
		logicClient:  newLogicClient(c),
		answerClient: answer.New(c.AnswerRPCClient.RegAddrs, c.AnswerRPCClient.Schema, c.AnswerRPCClient.SrvName, time.Duration(c.AnswerRPCClient.Dial)),
		deviceClient: device.New(c.DeviceRPCClient.RegAddrs, c.DeviceRPCClient.Schema, c.DeviceRPCClient.SrvName, time.Duration(c.DeviceRPCClient.Dial)),
		workPoolConn: workerpool.New(c.IMSub.MaxWorker),
		workPoolSend: workerpool.New(c.RevSub.MaxWorker),
	}
	s.pusher = Pusher{s: s}
	return s
}

func (s *Service) ListenMQ() {
	s.log.Info().Int("numbers", len(s.connConsumers)).Msg("start serve mq consumers")
	s.log.Info().Int("numbers", len(s.receivedConsumers)).Msg("start serve mq receivedConsumers")
	for _, c := range s.connConsumers {
		go c.Listen(s.asyncConn)
	}
	for _, c := range s.receivedConsumers {
		go c.Listen(s.asyncSend)
	}
}

func (s *Service) Shutdown(ctx context.Context) {
	down := make(chan struct{})
	go func() {
		s.workPoolConn.StopWait()
		s.workPoolSend.StopWait()
		close(down)
	}()

	select {
	case <-ctx.Done():
		return
	case <-down:
		return
	}
}

func (s *Service) asyncConn(msg *sarama.ConsumerMessage) error {
	s.workPoolConn.Submit(func() {
		s.log.Debug().Str("topic", msg.Topic).
			Bytes("key", msg.Key).
			Int32("partition", msg.Partition).
			Int64("offset", msg.Offset).
			Msg("consume process")
		err := s.consumeConn(msg)
		if err != nil {
			//s.log.Error().Err(err).Interface("msg", msg).Msg("proto.Unmarshal error")
		}
	})
	return nil
}

func (s *Service) asyncSend(msg *sarama.ConsumerMessage) error {
	s.workPoolSend.Submit(func() {
		s.log.Debug().Str("topic", msg.Topic).
			Bytes("key", msg.Key).
			Int32("partition", msg.Partition).
			Int64("offset", msg.Offset).
			Msg("consume process")
		err := s.consumePush(msg)
		if err != nil {
			//s.log.Error().Err(err).Interface("msg", msg).Msg("proto.Unmarshal error")
		}
	})
	return nil
}

func newLogicClient(cfg *config.Config) logic.LogicClient {
	rb := naming.NewResolver(cfg.LogicRPCClient.RegAddrs, cfg.LogicRPCClient.Schema)
	resolver.Register(rb)

	addr := fmt.Sprintf("%s:///%s", cfg.LogicRPCClient.Schema, cfg.LogicRPCClient.SrvName) // "schema://[authority]/service"
	fmt.Println("logic rpc client call addr:", addr)

	conn, err := grpc.NewGRPCConn(addr, time.Duration(cfg.LogicRPCClient.Dial))
	if err != nil {
		panic(err)
	}
	return logic.NewLogicClient(conn)
}

func (s *Service) consumeConn(msg *sarama.ConsumerMessage) (err error) {
	bizMsg := new(logic.BizMsg)
	if err = proto.Unmarshal(msg.Value, bizMsg); err != nil {
		s.log.Error().Err(err).Bytes("bizMsg byte", msg.Value).Msg("logic.BizMsg proto.Unmarshal error")
		return err
	}
	if bizMsg.AppId != s.cfg.AppId {
		return model.ErrAppId
	}
	switch bizMsg.GetOp() {
	case int32(comet.Op_Auth), int32(comet.Op_Disconnect), int32(comet.Op_ReceiveMsgReply), int32(comet.Op_SyncMsgReq):
		spanName := fmt.Sprintf("consume %s:%s", msg.Topic, comet.Op(bizMsg.GetOp()).String())
		if err := s.consumeConnTraceAndLogIpt(spanName, bizMsg, s.DealConn); err != nil {
			//TODO redo consume message
			return err
		}
	default:
		return model.ErrCustomNotSupport
	}
	return err
}

func (s *Service) consumePush(msg *sarama.ConsumerMessage) (err error) {
	bizMsg := new(record.PushMsg)
	if err = proto.Unmarshal(msg.Value, bizMsg); err != nil {
		s.log.Error().Err(err).Bytes("bizMsg byte", msg.Value).Msg("record.PushMsg proto.Unmarshal error")
		return err
	}
	if bizMsg.AppId != s.cfg.AppId {
		return model.ErrAppId
	}
	spanName := fmt.Sprintf("consume %s:%s", msg.Topic, "Push")
	if err := s.consumePushTraceAndLogIpt(spanName, bizMsg, s.DealPush); err != nil {
		//TODO redo consume message
		return err
	}
	return nil
}

func (s *Service) consumeConnTraceAndLogIpt(spanName string, bizMsg *logic.BizMsg, handler func(ctx context.Context, m *logic.BizMsg) error) error {
	// trace
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan(spanName)
	defer span.Finish()

	traceId := ""
	if jgSpan, ok := span.Context().(jaeger.SpanContext); ok {
		traceId = jgSpan.TraceID().String()
	}
	//get global log instance
	logger := s.log.With().Str(trace.TraceIDLogKey, traceId).Logger()
	span.LogFields(
		traceLog.String("AppId", bizMsg.GetAppId()),
		traceLog.String("Uid", bizMsg.GetFromId()),
		traceLog.String("ConnId", bizMsg.GetKey()),
	)
	ctx := opentracing.ContextWithSpan(context.Background(), span)
	ctx = logger.WithContext(ctx)

	err := handler(ctx, bizMsg)
	if err != nil {
		pointer := reflect.ValueOf(handler).Pointer()
		funcName := runtime.FuncForPC(pointer).Name()
		log.Error().Err(err).Msg(fmt.Sprintf("consume %s error", funcName))
		span.SetTag("ERROR", err)
		return err
	}
	return nil
}

func (s *Service) consumePushTraceAndLogIpt(spanName string, bizMsg *record.PushMsg, handler func(ctx context.Context, m *record.PushMsg) error) error {
	// trace
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan(spanName)
	defer span.Finish()

	traceId := ""
	if jgSpan, ok := span.Context().(jaeger.SpanContext); ok {
		traceId = jgSpan.TraceID().String()
	}
	//get global log instance
	logger := s.log.With().Str(trace.TraceIDLogKey, traceId).Logger()
	span.LogFields(
		traceLog.String("AppId", bizMsg.GetAppId()),
		traceLog.String("Uid", bizMsg.GetFromId()),
		traceLog.String("ConnId", bizMsg.GetKey()),
		traceLog.Int64("Mid", bizMsg.GetMid()),
		traceLog.String("Target", bizMsg.GetTarget()),
		traceLog.String("Channel Type", imparse.Channel(bizMsg.GetType()).String()),
		traceLog.String("Frame Type", bizMsg.GetFrameType()),
	)
	ctx := opentracing.ContextWithSpan(context.Background(), span)
	ctx = logger.WithContext(ctx)

	err := handler(ctx, bizMsg)
	if err != nil {
		pointer := reflect.ValueOf(handler).Pointer()
		funcName := runtime.FuncForPC(pointer).Name()
		log.Error().Err(err).Msg(fmt.Sprintf("consume %s error", funcName))
		span.SetTag("ERROR", err)
		return err
	}
	return nil
}

func (s *Service) parseDevice(m *logic.BizMsg) (*xproto.Login, error) {
	var p comet.Proto
	err := proto.Unmarshal(m.Msg, &p)
	if err != nil {
		return nil, err
	}
	var auth comet.AuthMsg
	err = proto.Unmarshal(p.Body, &auth)
	if err != nil {
		return nil, err
	}
	if len(auth.Ext) == 0 {
		return nil, errors.New("ext is nil")
	}
	var device xproto.Login
	err = proto.Unmarshal(auth.Ext, &device)
	if err != nil {
		return nil, err
	}
	return &device, nil
}

func (s *Service) DealConn(ctx context.Context, m *logic.BizMsg) error {
	if m.AppId != s.cfg.AppId {
		return model.ErrAppId
	}
	logger := zerolog.Ctx(ctx).With().Str("AppId", m.GetAppId()).Str("Uid", m.GetFromId()).Str("ConnId", m.GetKey()).Logger()

	switch m.Op {
	case int32(comet.Op_Auth):
		logger.Info().Msg("user login with key")
		//将用户设备信息存入缓存
		dev, err := s.parseDevice(m)
		if err != nil {
			logger.Error().Err(err).Msg("parseDevice failed")
		} else {
			now := uint64(util.TimeNowUnixNano() / int64(time.Millisecond))
			err = s.deviceClient.AddDevice(ctx, &device.Device{
				Uid:         m.FromId,
				ConnectId:   m.Key,
				DeviceUUid:  dev.GetUuid(),
				DeviceType:  int32(dev.Device),
				DeviceName:  dev.GetDeviceName(),
				Username:    dev.Username,
				DeviceToken: dev.DeviceToken,
				IsEnabled:   false,
				AddTime:     now,
			})
			if err != nil {
				logger.Error().Err(err).Msg("AddDeviceInfo failed")
			}
			//发送登录通知
			err = s.UniCastSignalEndpointLogin(ctx, m.GetFromId(), &xproto.SignalEndpointLogin{
				Uuid:       dev.GetUuid(),
				Device:     dev.Device,
				DeviceName: dev.GetDeviceName(),
				Datetime:   now,
			})
			if err != nil {
				logger.Error().Err(err).Msg("UniCastSignalEndpointLogin failed")
				return err
			}
		}
		if dev != nil && (dev.Device == xproto.Device_Android || dev.Device == xproto.Device_IOS) {
			err = s.dao.BatchPushPublish(ctx, m.Key, m.GetFromId())
			if err != nil {
				logger.Error().Err(err).Msg("BatchPushPublish failed")
			}
		}
		//连接群聊
		err = s.JoinGroups(ctx, m.GetFromId(), m.GetKey())
		if err != nil {
			logger.Error().Err(err).Msg("JoinGroups failed")
		}
	case int32(comet.Op_Disconnect):
		logger.Info().Msg("user logout with key")
		err := s.dao.ClearConnSeq(m.Key)
		if err != nil {
			logger.Error().Err(err).Msg("ClearConnSeq failed")
		}
		err = s.deviceClient.EnableThreadPushDevice(ctx, &device.EnableThreadPushDeviceRequest{
			Uid:    m.GetFromId(),
			ConnId: m.Key,
		})
		if err != nil {
			logger.Error().Err(err).Msg("EnableDeviceInfo failed")
		}
	case int32(comet.Op_ReceiveMsgReply):
		var p comet.Proto
		err := proto.Unmarshal(m.Msg, &p)
		if err != nil {
			logger.Error().Err(err).Msg("unmarshal proto error")
			return err
		}
		item, err := s.lh.GetLogsIndex(m.Key, p.Ack)
		if err != nil {
			logger.Error().Err(err).Int32("ack", p.Ack).Msg("GetConnSeqIndex failed")
			return err
		}
		if item == nil {
			//TODO print err
			return nil
		}
		dev, err := s.deviceClient.GetDeviceByConnID(ctx, &device.GetDeviceByConnIdRequest{
			Uid:    m.GetFromId(),
			ConnID: m.GetKey(),
		})
		if err != nil || dev == nil {
			logger.Error().Err(err).Int32("ack", p.Ack).Str("uid", m.GetFromId()).Str("connID", m.GetKey()).Msg("GetDeviceByConnID failed")
		}
		if xproto.Device(dev.GetDeviceType()) == xproto.Device_Android || xproto.Device(dev.GetDeviceType()) == xproto.Device_IOS {
			err = s.dao.MarkReadPublish(ctx, m.Key, m.GetFromId(), imparse.FrameType(item.Type), item.Logs)
			if err != nil {
				logger.Error().Err(err).Int32("ack", p.Ack).Msg("MarkReadPublish failed")
			}
			switch item.Type {
			case string(chat.PrivateFrameType):
				err := s.UniCastSignalReceived(ctx, item)
				if err != nil {
					logger.Error().Err(err).Int32("ack", p.Ack).Msg("UniCastSignalReceived failed")
					return err
				}
			default:
			}
		}
	case int32(comet.Op_SyncMsgReq):
		//TODO disabled
		return nil
		var p comet.Proto
		var pro comet.SyncMsg
		err := proto.Unmarshal(m.Msg, &p)
		if err != nil {
			logger.Error().Err(err).Msg("unmarshal proto error")
			return err
		}
		err = proto.Unmarshal(p.Body, &pro)
		if err != nil {
			logger.Error().Err(err).Interface("option", comet.Op_SyncMsgReq).Msg("Unmarshal failed")
			break
		}
		err = s.dao.SyncPublish(ctx, m.Key, m.FromId, pro.LogId)
		if err != nil {
			logger.Error().Err(err).Int64("start", pro.LogId).Msg("SyncMsg failed")
			break
		}
	default:
		return model.ErrCustomNotSupport
	}
	return nil
}

func (s *Service) DealPush(ctx context.Context, m *record.PushMsg) error {
	logger := zerolog.Ctx(ctx).With().Str("AppId", m.GetAppId()).Str("Uid", m.GetFromId()).Str("ConnId", m.GetKey()).
		Int64("Mid", m.GetMid()).Str("Target", m.GetTarget()).Int32("Push Type", m.GetType()).Str("Frame Type", m.GetFrameType()).Logger()
	//sendB
	var err error
	switch m.GetType() {
	case int32(imparse.UniCast):
		err = s.pusher.UniCast(ctx, m)
	case int32(imparse.GroupCast):
		err = s.pusher.GroupCast(ctx, m)
	default:
		err = fmt.Errorf("push type %v undefined", m.GetType())
	}
	if err != nil {
		logger.Error().Stack().Err(err).Msg("DealPush error")
	}
	logger.Debug().Msg("pass DealSend")
	return nil
}
