package service

import (
	"bytes"
	"context"
	"fmt"
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
	"github.com/uber/jaeger-client-go"
	"github.com/txchat/dtalk/pkg/naming"
	"github.com/txchat/dtalk/pkg/net/grpc"
	"github.com/txchat/dtalk/service/record/answer/config"
	"github.com/txchat/dtalk/service/record/answer/dao"
	"github.com/txchat/dtalk/service/record/answer/model"
	"github.com/txchat/dtalk/service/record/kafka/consumer"
	"github.com/txchat/im-pkg/trace"
	comet "github.com/txchat/im/api/comet/grpc"
	logic "github.com/txchat/im/api/logic/grpc"
	"github.com/txchat/imparse"
	"github.com/txchat/imparse/chat"
	"google.golang.org/grpc/resolver"
)

type Service struct {
	cfg         *config.Config
	log         zerolog.Logger
	dao         *dao.Dao
	consumers   map[string]*consumer.Consumer
	logicClient logic.LogicClient

	answer    imparse.Answer
	rcpAnswer imparse.Answer
	parser    chat.StandardParse

	workPool *workerpool.WorkerPool
}

func New(c *config.Config) *Service {
	s := &Service{
		cfg: c,
		log: log.Logger,
		dao: dao.New(c),
		consumers: consumer.NewKafkaConsumers(
			fmt.Sprintf("goim-%s-topic", c.AppId),
			fmt.Sprintf("goim-%s-group", c.AppId), c.MQSub.Brokers, c.MQSub.Number),
		logicClient: newLogicClient(c),
		workPool:    workerpool.New(c.MQSub.MaxWorker),
	}
	db := DB{
		dao: s.dao,
	}
	exec := Exec{
		appId:       s.cfg.AppId,
		dao:         s.dao,
		logicClient: s.logicClient,
	}
	withoutAckExec := withoutAckExec{
		appId:       s.cfg.AppId,
		dao:         s.dao,
		logicClient: s.logicClient,
	}
	trace := &Trace{}
	s.answer = imparse.NewStandardAnswer(&db, &exec, trace, db.GetFilters())
	s.rcpAnswer = imparse.NewStandardAnswer(&db, &withoutAckExec, trace, db.GetFilters())
	return s
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

func (s *Service) Shutdown(ctx context.Context) {
	down := make(chan struct{})
	go func() {
		s.workPool.StopWait()
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
	s.log.Info().Int("numbers", len(s.consumers)).Msg("start serve mq consumers")
	for _, c := range s.consumers {
		go c.Listen(s.asyncRev)
	}
}

func (s *Service) asyncRev(msg *sarama.ConsumerMessage) error {
	s.workPool.Submit(func() {
		//init trace and log
		s.log.Debug().Str("topic", msg.Topic).
			Bytes("key", msg.Key).
			Int32("partition", msg.Partition).
			Int64("offset", msg.Offset).
			Msg(fmt.Sprintf("consume %s", msg.Topic))
		err := s.consume(msg)
		if err != nil {
			//s.log.Error().Err(err).Msg("asyncRev consume error")
			return
		}
	})
	return nil
}

func (s *Service) consume(msg *sarama.ConsumerMessage) error {
	bizMsg := new(logic.BizMsg)
	if err := proto.Unmarshal(msg.Value, bizMsg); err != nil {
		s.log.Error().Err(err).Bytes("bizMsg byte", msg.Value).Msg("logic.BizMsg proto.Unmarshal error")
		return err
	}
	if bizMsg.AppId != s.cfg.AppId {
		return model.ErrAppId
	}
	switch bizMsg.GetOp() {
	case int32(comet.Op_SendMsg):
		if err := s.consumeTraceAndLogIpt(msg, bizMsg, s.DealOpSendMsg); err != nil {
			//TODO redo consume message
			return err
		}
	default:
		return model.ErrCustomNotSupport
	}
	return nil
}

func (s *Service) consumeTraceAndLogIpt(msg *sarama.ConsumerMessage, bizMsg *logic.BizMsg, handler func(ctx context.Context, m *logic.BizMsg) error) error {
	// trace
	tracer := opentracing.GlobalTracer()
	spanContext, err := trace.ExtractMQHeader(tracer, msg)
	if err != nil {
		// maybe rpc client not transform trace instance, it should be work continue
	}
	spanName := fmt.Sprintf("consume %s:%s", msg.Topic, comet.Op_SendMsg.String())
	span := tracer.StartSpan(spanName, opentracing.ChildOf(spanContext))
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
		traceLog.String("Option", comet.Op(bizMsg.GetOp()).String()),
	)
	ctx := opentracing.ContextWithSpan(context.Background(), span)
	ctx = logger.WithContext(ctx)

	err = handler(ctx, bizMsg)
	if err != nil {
		pointer := reflect.ValueOf(handler).Pointer()
		funcName := runtime.FuncForPC(pointer).Name()
		log.Error().Err(err).Msg(fmt.Sprintf("consume %s error", funcName))
		span.SetTag("ERROR", err)
		return err
	}
	return nil
}

func (s *Service) DealOpSendMsg(ctx context.Context, m *logic.BizMsg) error {
	logger := zerolog.Ctx(ctx).With().Str("AppId", m.GetAppId()).Str("Uid", m.GetFromId()).Str("ConnId", m.GetKey()).Logger()

	logger.Debug().Msg("start create frame")
	frame, err := s.parser.NewFrame(m.GetKey(), m.GetFromId(), bytes.NewReader(m.GetMsg()))
	if err != nil {
		logger.Error().Stack().Err(err).Msg("NewFrame error")
		return err
	}
	logger.Debug().Msg("start check frame")
	if err := s.answer.Check(ctx, &checker, frame); err != nil {
		logger.Error().Stack().Err(err).Msg("Check error")
		return err
	}
	logger.Debug().Msg("start frame filter")
	_, err = s.answer.Filter(ctx, frame)
	if err != nil {
		logger.Error().Stack().Err(err).Msg("Filter error")
		return err
	}
	logger.Debug().Msg("start frame transport")
	err = s.answer.Transport(ctx, frame)
	if err != nil {
		logger.Error().Stack().Err(err).Msg("Transport error")
		return err
	}
	// ACk A
	logger.Debug().Msg("start frame ack")
	_, err = s.answer.Ack(ctx, frame)
	if err != nil {
		logger.Error().Stack().Err(err).Msg("Ack error")
		return err
	}
	logger.Debug().Msg("deal msg send success")
	return nil
}
