package mq

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/proto"
	xkafka "github.com/oofpgDLD/kafka-go"
	"github.com/txchat/dtalk/api/proto/message"
	"github.com/txchat/dtalk/app/services/pusher/internal/config"
	innerLogic "github.com/txchat/dtalk/app/services/pusher/internal/logic"
	"github.com/txchat/dtalk/app/services/pusher/internal/model"
	"github.com/txchat/dtalk/app/services/pusher/internal/svc"
	"github.com/txchat/dtalk/app/services/pusher/pusher"
	"github.com/txchat/dtalk/internal/proto/record"
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
	cfg.PushDealConsumerConfig.Topic = fmt.Sprintf("biz-%s-push", cfg.AppID)
	cfg.PushDealConsumerConfig.Group = fmt.Sprintf("biz-%s-push-pusher", cfg.AppID)
	//new batch consumer
	consumer := xkafka.NewConsumer(cfg.PushDealConsumerConfig, nil)
	logx.Info("dial kafka broker success")
	bc := xkafka.NewBatchConsumer(cfg.PushDealBatchConsumerConf, xkafka.WithHandle(s.handleFunc), consumer)
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
	ctx := context.Background()
	msg := new(record.PushMsgMQ)
	if err := proto.Unmarshal(data, msg); err != nil {
		s.Error("logic.BizMsg proto.Unmarshal error", "err", err)
		return err
	}
	if msg.AppId != s.Config.AppID {
		return model.ErrAppID
	}
	if err := s.consumerOnePush(ctx, msg); err != nil {
		//TODO redo consume message
		return err
	}
	return nil
}

func (s *Service) consumerOnePush(ctx context.Context, m *record.PushMsgMQ) error {
	switch m.GetChannel() {
	case message.Channel_Private:
		l := innerLogic.NewPushListLogic(ctx, s.svcCtx)
		_, err := l.PushList(&pusher.PushListReq{
			App:  m.GetAppId(),
			From: m.GetFrom(),
			Uid:  []string{m.GetTarget()},
			Body: m.GetMsg(),
		})
		return err
	case message.Channel_Group:
		l := innerLogic.NewPushGroupLogic(ctx, s.svcCtx)
		_, err := l.PushGroup(&pusher.PushGroupReq{
			App:  m.GetAppId(),
			Gid:  m.GetTarget(),
			Body: m.GetMsg(),
		})
		return err
	default:
		return fmt.Errorf("push type %v undefined", m.GetChannel())
	}
}
