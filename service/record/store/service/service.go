package service

import (
	"bytes"
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/gammazero/workerpool"
	"github.com/golang/protobuf/proto"
	"github.com/opentracing/opentracing-go"
	traceLog "github.com/opentracing/opentracing-go/log"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
	device "github.com/txchat/dtalk/service/device/api"
	"github.com/txchat/dtalk/service/record/kafka/consumer"
	record "github.com/txchat/dtalk/service/record/proto"
	"github.com/txchat/dtalk/service/record/store/config"
	"github.com/txchat/dtalk/service/record/store/dao"
	"github.com/txchat/dtalk/service/record/store/model"
	"github.com/txchat/imparse"
	"github.com/txchat/imparse/chat"
	xproto "github.com/txchat/imparse/proto"
	"time"
)

type Service struct {
	log               zerolog.Logger
	dao               *dao.Dao
	cfg               *config.Config
	receivedConsumers map[string]*consumer.Consumer
	storeConsumers    map[string]*consumer.Consumer

	parser  chat.StandardParse
	storage imparse.Storage

	deviceClient *device.Client

	workPoolRev   *workerpool.WorkerPool
	workPoolStore *workerpool.WorkerPool
}

func New(c *config.Config) *Service {
	s := &Service{
		log: zlog.Logger,
		dao: dao.New(c),
		cfg: c,
		receivedConsumers: consumer.NewKafkaConsumers(
			fmt.Sprintf("received-%s-topic", c.AppId),
			fmt.Sprintf("received-%s-store-group", c.AppId), c.RevSub.Brokers, c.RevSub.Number),
		storeConsumers: consumer.NewKafkaConsumers(
			fmt.Sprintf("store-%s-topic", c.AppId),
			fmt.Sprintf("store-%s-group", c.AppId), c.StoreSub.Brokers, c.StoreSub.Number),
		deviceClient:  device.New(c.DeviceRPCClient.RegAddrs, c.DeviceRPCClient.Schema, c.DeviceRPCClient.SrvName, time.Duration(c.DeviceRPCClient.Dial)),
		workPoolRev:   workerpool.New(c.RevSub.MaxWorker),
		workPoolStore: workerpool.New(c.StoreSub.MaxWorker),
	}
	db := &DB{s: s}
	s.storage = imparse.NewStandardStorage(db)
	return s
}

func (s Service) Config() *config.Config {
	return s.cfg
}

func (s *Service) Shutdown(ctx context.Context) {
	down := make(chan struct{})
	go func() {
		s.workPoolRev.StopWait()
		s.workPoolStore.StopWait()
		close(down)
	}()

	select {
	case <-ctx.Done():
		return
	case <-down:
		return
	}
}

func (s *Service) ListenMQ() {
	s.log.Info().Int("numbers", len(s.receivedConsumers)).Msg("start serve mq receivedConsumers")
	s.log.Info().Int("numbers", len(s.storeConsumers)).Msg("start serve mq storeConsumers")
	for _, c := range s.receivedConsumers {
		go c.Listen(s.asyncRev)
	}
	for _, c := range s.storeConsumers {
		go c.Listen(s.asyncStore)
	}
}

func (s *Service) asyncRev(msg *sarama.ConsumerMessage) error {
	s.workPoolRev.Submit(func() {
		err := s.consumeStore(msg)
		if err != nil {
			s.log.Error().Err(err).Interface("msg", msg).Msg("proto.Unmarshal error")
		}
	})
	return nil
}

func (s *Service) asyncStore(msg *sarama.ConsumerMessage) error {
	s.workPoolStore.Submit(func() {
		err := s.consume(msg)
		if err != nil {
			s.log.Error().Err(err).Interface("msg", msg).Msg("proto.Unmarshal error")
		}
	})
	return nil
}

func (s *Service) consume(msg *sarama.ConsumerMessage) (err error) {
	// trace
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan(fmt.Sprintf("consume %s", msg.Topic))
	defer func() {
		if err != nil {
			span.SetTag("ERROR", err)
		}
		span.Finish()
	}()
	ctx := opentracing.ContextWithSpan(context.Background(), span)

	bizMsg := new(record.RecordDeal)
	if err = proto.Unmarshal(msg.Value, bizMsg); err != nil {
		s.log.Error().Err(err).Interface("msg", msg).Msg("proto.Unmarshal error")
		return err
	}
	span.LogFields(
		traceLog.String("AppId", bizMsg.GetAppId()),
		traceLog.String("Uid", bizMsg.GetFromId()),
		traceLog.String("ConnId", bizMsg.GetKey()),
		traceLog.String("Option", bizMsg.GetOpt().String()),
	)
	s.log.Debug().Str("topic", msg.Topic).
		Bytes("key", msg.Key).
		Int32("partition", msg.Partition).
		Int64("offset", msg.Offset).
		Msg("consume process")
	err = s.DealConn(ctx, bizMsg)
	if err == model.ErrConsumeRedo {
		//TODO redo consume message
	}
	return err
}

func (s *Service) consumeStore(msg *sarama.ConsumerMessage) (err error) {
	// trace
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan(fmt.Sprintf("consume %s", msg.Topic))
	defer func() {
		if err != nil {
			span.SetTag("ERROR", err)
		}
		span.Finish()
	}()
	ctx := opentracing.ContextWithSpan(context.Background(), span)

	bizMsg := new(record.PushMsg)
	if err = proto.Unmarshal(msg.Value, bizMsg); err != nil {
		s.log.Error().Err(err).Interface("msg", msg).Msg("proto.Unmarshal error")
		return err
	}
	span.LogFields(
		traceLog.String("AppId", bizMsg.GetAppId()),
		traceLog.String("Uid", bizMsg.GetFromId()),
		traceLog.String("ConnId", bizMsg.GetKey()),
		traceLog.Int64("Mid", bizMsg.GetMid()),
		traceLog.String("Target", bizMsg.GetTarget()),
		traceLog.String("Channel Type", imparse.Channel(bizMsg.GetType()).String()),
		traceLog.String("Frame Type", bizMsg.GetFrameType()),
	)
	s.log.Debug().Str("topic", msg.Topic).
		Bytes("key", msg.Key).
		Int32("partition", msg.Partition).
		Int64("offset", msg.Offset).
		Msg("consume process")
	err = s.DealStore(ctx, bizMsg)
	if err == model.ErrConsumeRedo {
		//TODO redo consume message
	}
	return err
}

func (s *Service) DealConn(ctx context.Context, m *record.RecordDeal) error {
	if m.AppId != s.cfg.AppId {
		return model.ErrAppId
	}

	switch m.Opt {
	case record.Operation_BatchPush:
		s.log.Info().Str("uid", m.FromId).Str("key", m.Key).Msg("user batch push")
		//检查用户ws是否保持连接
		if b, err := s.dao.CheckOnline(context.TODO(), m.Key); !b {
			if err != nil {
				s.log.Error().Err(err).Str("uid", m.FromId).Str("key", m.Key).Msg("CheckOnline failed")
			}
			break
		}
		err := s.SendUnReadMsg(m.Key, m.FromId)
		if err != nil {
			s.log.Error().Err(err).Str("uid", m.FromId).Str("key", m.Key).Msg("SendUnReadMsg failed")
			break
		}
	case record.Operation_MarkRead:
		var mk record.Marked
		err := proto.Unmarshal(m.Msg, &mk)
		if err != nil {
			s.log.Error().Err(err).Str("uid", m.FromId).Str("key", m.Key).Msg("unmarshal proto error")
			return err
		}
		for _, mid := range mk.GetMids() {
			s.log.Debug().Str("uid", m.FromId).Str("key", m.Key).Int64("mid", mid).Msg("msg received")
			err = s.MarkReceived(mk.Type, m.GetFromId(), mid)
			if err != nil {
				s.log.Error().Err(err).Str("uid", m.FromId).Str("key", m.Key).Int64("mid", mid).Msg("MarkMsgReceived failed")
			}
		}
	case record.Operation_SyncMsg:
		//TODO disabled
		return nil
		var sy record.Sync
		err := proto.Unmarshal(m.Msg, &sy)
		if err != nil {
			s.log.Error().Err(err).Str("uid", m.FromId).Str("key", m.Key).Msg("unmarshal proto error")
			return err
		}
		err = s.SendSyncMsg(m.Key, m.FromId, sy.Mid)
		if err != nil {
			s.log.Error().Err(err).Str("uid", m.FromId).Str("key", m.Key).Msg("SyncMsg failed")
			break
		}
	default:
		return model.ErrCustomNotSupport
	}
	return nil
}

func (s *Service) DealStore(ctx context.Context, m *record.PushMsg) error {
	s.log.Info().Str("appId", m.AppId).Str("key", m.Key).Str("from", m.FromId).Msg("do store")
	if m.AppId != s.cfg.AppId {
		return model.ErrAppId
	}

	frame, err := s.parser.NewFrame(m.GetKey(), m.GetFromId(), bytes.NewReader(m.GetMsg()), chat.WithMid(m.GetMid()), chat.WithTarget(m.GetTarget()), chat.WithTransmissionMethod(imparse.Channel(m.GetType())))
	if err != nil {
		s.log.Error().Stack().Err(err).Str("key", m.Key).Str("from", m.FromId).Msg("NewFrame error")
		return err
	}
	//TODO 暂时处理，之后device信息统一通过answer服务传递
	dev, err := s.deviceClient.GetDeviceByConnID(ctx, &device.GetDeviceByConnIdRequest{
		Uid:    m.GetFromId(),
		ConnID: m.GetKey(),
	})
	if err != nil || dev == nil {
		s.log.Error().Err(err).Str("uid", m.GetFromId()).Str("connID", m.GetKey()).Msg("GetDeviceByConnID failed")
	}
	if err := s.storage.SaveMsg(ctx, xproto.Device(dev.GetDeviceType()), frame); err != nil {
		s.log.Error().Stack().Err(err).Str("key", m.Key).Str("from", m.FromId).Msg("Store error")
		return err
	}
	s.log.Debug().Str("appId", m.AppId).Str("key", m.Key).Str("from", m.FromId).Msg("pass Store")
	return nil
}
