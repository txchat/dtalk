package mq

import (
	"context"
	"fmt"

	"github.com/txchat/im/api/protocol"
	"github.com/txchat/im/app/logic/logicclient"

	"github.com/golang/protobuf/proto"
	"github.com/txchat/dtalk/app/services/answer/internal/config"
	innerLogic "github.com/txchat/dtalk/app/services/answer/internal/logic"
	"github.com/txchat/dtalk/app/services/answer/internal/model"
	"github.com/txchat/dtalk/app/services/answer/internal/svc"
	xkafka "github.com/txchat/pkg/mq/kafka"
	"github.com/zeromicro/go-zero/core/logx"
)

type Service struct {
	logx.Logger
	Config        config.Config
	svcCtx        *svc.ServiceContext
	batchConsumer *xkafka.BatchConsumer
}

func NewService(cfg config.Config, svcCtx *svc.ServiceContext) *Service {
	s := &Service{
		Logger: logx.WithContext(context.TODO()),
		Config: cfg,
		svcCtx: svcCtx,
	}
	//topic config
	cfg.ConsumerConfig.Topic = fmt.Sprintf("goim-%s-topic", cfg.AppID)
	cfg.ConsumerConfig.Group = fmt.Sprintf("goim-%s-answer-group", cfg.AppID)
	//new batch consumer
	consumer := xkafka.NewConsumer(cfg.ConsumerConfig, nil)
	logx.Info("dial kafka broker success")
	bc := xkafka.NewBatchConsumer(cfg.BatchConsumerConf, xkafka.WithHandle(s.handleFunc), consumer)
	s.batchConsumer = bc
	return s
}

func (s *Service) Serve() {
	s.batchConsumer.Start()
}

func (s *Service) Shutdown(ctx context.Context) {
	s.batchConsumer.GracefulStop(ctx)
}

func (s *Service) handleFunc(key string, data []byte) error {
	bizMsg := new(logicclient.BizMsg)
	if err := proto.Unmarshal(data, bizMsg); err != nil {
		s.Error("logic.BizMsg proto.Unmarshal error", "err", err)
		return err
	}
	if bizMsg.AppId != s.Config.AppID {
		return model.ErrAppID
	}
	switch bizMsg.GetOp() {
	case int32(protocol.Op_SendMsg):
		if err := s.DealOpSendMsg(context.TODO(), bizMsg); err != nil {
			//TODO redo consume message
			return err
		}
	default:
		return model.ErrCustomNotSupport
	}
	return nil
}

func (s *Service) DealOpSendMsg(ctx context.Context, m *logicclient.BizMsg) error {
	s.Slow("start create frame")
	l := innerLogic.NewInnerPushLogic(ctx, s.svcCtx)
	_, _, err := l.PushToClient(s.svcCtx.AnswerInter4rmq, m.GetKey(), m.GetFromId(), m.GetMsg())
	if err != nil {
		return err
	}
	s.Slow("deal msg send success")
	return nil
}
