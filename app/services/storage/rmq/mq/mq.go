package mq

import (
	"context"
	"fmt"

	"github.com/txchat/dtalk/internal/proto/record"

	"github.com/golang/protobuf/proto"
	"github.com/txchat/dtalk/api/proto/chat"
	"github.com/txchat/dtalk/api/proto/message"
	"github.com/txchat/dtalk/api/proto/signal"
	"github.com/txchat/dtalk/app/services/storage/internal/config"
	"github.com/txchat/dtalk/app/services/storage/internal/model"
	"github.com/txchat/dtalk/app/services/storage/internal/svc"
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
	cfg.StoreDealConsumerConfig.Topic = fmt.Sprintf("biz-%s-store", cfg.AppID)
	cfg.StoreDealConsumerConfig.Group = fmt.Sprintf("biz-%s-store-storage", cfg.AppID)
	//new batch consumer
	consumer := xkafka.NewConsumer(cfg.StoreDealConsumerConfig, nil)
	logx.Info("dial kafka broker success")
	bc := xkafka.NewBatchConsumer(cfg.StoreDealBatchConsumerConf, xkafka.WithHandle(s.handleFunc), consumer)
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
	msg := new(record.StoreMsgMQ)
	if err := proto.Unmarshal(data, msg); err != nil {
		s.Error("logic.BizMsg proto.Unmarshal error", "err", err)
		return err
	}
	if msg.AppId != s.Config.AppID {
		return model.ErrAppID
	}
	if err := s.DealStore(ctx, msg); err != nil {
		//TODO redo consume message
		return err
	}
	return nil
}

func (s *Service) DealStore(ctx context.Context, m *record.StoreMsgMQ) error {
	chatProto := m.GetChat()
	switch chatProto.GetType() {
	case chat.Chat_message:
		var msg *message.Message
		err := proto.Unmarshal(chatProto.GetBody(), msg)
		if err != nil {
			return err
		}
		switch msg.GetChannelType() {
		case message.Channel_Private:
			err = s.svcCtx.StorePrivateMessage(msg)
			if err != nil {
				return err
			}
		case message.Channel_Group:
			err = s.svcCtx.StoreGroupMessage(m.GetTarget(), msg)
			if err != nil {
				return err
			}
		}
	case chat.Chat_signal:
		var sig *signal.Signal
		err := proto.Unmarshal(chatProto.GetBody(), sig)
		if err != nil {
			return err
		}
		err = s.svcCtx.StoreSignal(m.GetTarget(), chatProto.GetSeq(), sig)
		if err != nil {
			return err
		}
	}
	return nil
}
