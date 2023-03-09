package mq

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/proto"
	xkafka "github.com/oofpgDLD/kafka-go"
	"github.com/txchat/dtalk/api/proto/chat"
	"github.com/txchat/dtalk/api/proto/message"
	"github.com/txchat/dtalk/app/services/transfer/internal/config"
	"github.com/txchat/dtalk/app/services/transfer/internal/model"
	"github.com/txchat/dtalk/app/services/transfer/internal/svc"
	"github.com/txchat/im/api/protocol"
	"github.com/txchat/im/app/logic/logicclient"
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
	cfg.ConsumerConfig.Topic = fmt.Sprintf("goim-%s-receive", cfg.AppID)
	cfg.ConsumerConfig.Group = fmt.Sprintf("goim-%s-receive-transfer", cfg.AppID)
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
	ctx := context.Background()
	receivedMsg := new(logicclient.ReceivedMessage)
	if err := proto.Unmarshal(data, receivedMsg); err != nil {
		s.Error("proto.Unmarshal receivedMessage error", "err", err)
		return err
	}
	if receivedMsg.GetAppId() != s.Config.AppID {
		s.Error(model.ErrAppID.Error())
		return model.ErrAppID
	}

	switch receivedMsg.GetOp() {
	case protocol.Op_Message:
		var chatProto *chat.Chat
		err := proto.Unmarshal(receivedMsg.GetBody(), chatProto)
		if err != nil {
			return err
		}

		var msg *message.Message
		err = proto.Unmarshal(chatProto.GetBody(), msg)
		if err != nil {
			return err
		}
		msg.From = receivedMsg.GetFrom()
		err = s.svcCtx.TransferMessage(ctx, msg.GetChannelType(), msg.GetTarget(), receivedMsg.GetFrom(), chatProto)
		if err != nil {
			//TODO log
			return err
		}
	default:
		return model.ErrCustomNotSupport
	}
	return nil
}
