package mq

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/txchat/dtalk/app/services/offline/internal/config"
	"github.com/txchat/dtalk/app/services/offline/internal/model"
	"github.com/txchat/dtalk/app/services/offline/internal/svc"
	"github.com/txchat/dtalk/internal/proto/offline"
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
	cfg.DealConsumerConfig.Topic = fmt.Sprintf("biz-%s-offlinepush", cfg.AppID)
	cfg.DealConsumerConfig.Group = fmt.Sprintf("biz-%s-offlinepush-offline", cfg.AppID)
	//new batch consumer
	consumer := xkafka.NewConsumer(cfg.DealConsumerConfig, nil)
	logx.Info("dial kafka broker success")
	bc := xkafka.NewBatchConsumer(cfg.DealBatchConsumerConf, xkafka.WithHandle(s.handleFunc), consumer)
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
	bizMsg := new(offline.ThirdPartyPushMQ)
	if err := proto.Unmarshal(data, bizMsg); err != nil {
		s.Error("logic.BizMsg proto.Unmarshal error", "err", err)
		return err
	}
	if bizMsg.AppId != s.Config.AppID {
		return model.ErrAppID
	}
	if err := s.svcCtx.PushOffline(context.Background(), bizMsg); err != nil {
		//TODO redo consume message
		return err
	}
	return nil
}
