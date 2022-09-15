package connmq

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/txchat/dtalk/app/services/pusher/internal/config"
	innerLogic "github.com/txchat/dtalk/app/services/pusher/internal/logic"
	"github.com/txchat/dtalk/app/services/pusher/internal/model"
	"github.com/txchat/dtalk/app/services/pusher/internal/svc"
	xkafka "github.com/txchat/dtalk/pkg/mq/kafka"
	record "github.com/txchat/dtalk/service/record/proto"
	"github.com/txchat/imparse"
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
	cfg.PushDealConsumerConfig.Topic = fmt.Sprintf("received-%s-topic", cfg.AppID)
	cfg.PushDealConsumerConfig.Group = fmt.Sprintf("received-%s-group", cfg.AppID)
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
	bizMsg := new(record.PushMsg)
	if err := proto.Unmarshal(data, bizMsg); err != nil {
		s.Error("logic.BizMsg proto.Unmarshal error", "err", err)
		return err
	}
	if bizMsg.AppId != s.Config.AppID {
		return model.ErrAppID
	}
	if err := s.DealPush(context.TODO(), bizMsg); err != nil {
		//TODO redo consume message
		return err
	}
	return nil
}

func (s *Service) DealPush(ctx context.Context, m *record.PushMsg) error {
	//sendB
	var err error
	switch m.GetType() {
	case int32(imparse.UniCast):
		l := innerLogic.NewPusherLogic(ctx, s.svcCtx)
		err = l.UniCast(m)
	case int32(imparse.GroupCast):
		l := innerLogic.NewPusherLogic(ctx, s.svcCtx)
		err = l.GroupCast(m)
	default:
		err = fmt.Errorf("push type %v undefined", m.GetType())
	}
	if err != nil {
		s.Error("DealPush error", "err", err)
	}
	return nil
}
