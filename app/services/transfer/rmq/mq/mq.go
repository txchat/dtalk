package mq

import (
	"context"
	"fmt"

	"github.com/txchat/dtalk/app/services/device/deviceclient"
	"github.com/txchat/dtalk/app/services/transfer/internal/logic"
	"github.com/txchat/dtalk/app/services/transfer/transfer"

	"github.com/golang/protobuf/proto"
	"github.com/txchat/dtalk/app/services/transfer/internal/config"
	"github.com/txchat/dtalk/app/services/transfer/internal/model"
	"github.com/txchat/dtalk/app/services/transfer/internal/svc"
	"github.com/txchat/im/api/protocol"
	"github.com/txchat/im/app/logic/logicclient"
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

	//TODO protocol gateway check
	var p protocol.Proto
	body := receivedMsg.GetBody()
	err := proto.Unmarshal(body, &p)
	if err != nil {
		return err
	}

	switch receivedMsg.GetOp() {
	case protocol.Op_SendMsg:
		dev, err := s.svcCtx.DeviceClient.GetDeviceByConnId(ctx, &deviceclient.GetDeviceByConnIdRequest{
			Uid:    receivedMsg.GetFrom(),
			ConnID: key,
		})
		if err != nil {
			return err
		}
		uuid := dev.GetDeviceUUid()
		l := logic.NewCheckMessageResendLogic(ctx, s.svcCtx)
		resp, err := l.CheckMessageResend(&transfer.CheckMessageResendReq{
			From: receivedMsg.GetFrom(),
			Uuid: uuid,
			Seq:  p.GetSeq(),
		})
		if err != nil {
			return err
		}
		if resp.GetMid() != 0 {
			break
		}
		err = s.svcCtx.TransferMessage(ctx, receivedMsg.GetFrom(), &p)
		if err != nil {
			//TODO log
		}
	case protocol.Op_ReceiveMsgReply:
		msgIndex, err := s.svcCtx.Repo.GetUserClientSeqByMid(ctx, p.GetMid())
		if err != nil {
			return err
		}

		lastSeq, err := s.svcCtx.Repo.UpdateUserLatestRev(ctx, msgIndex.UUID, msgIndex.UID, receivedMsg.GetFrom(), msgIndex.Seq)
		if err != nil {
			return err
		}

		err := s.svcCtx.SignalHub.MessageReceived(ctx, item)
		if err != nil {
			s.Error("UniCastSignalReceived failed",
				"err", err,
				"ack", p.GetAck(),
			)
			return err
		}
	default:
		return model.ErrCustomNotSupport
	}
	return nil
}
