package mq

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/txchat/dtalk/app/services/device/internal/config"
	"github.com/txchat/dtalk/app/services/device/internal/model"
	"github.com/txchat/dtalk/app/services/device/internal/svc"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/im/api/protocol"
	"github.com/txchat/im/app/logic/logicclient"
	"github.com/txchat/imparse/proto/auth"
	"github.com/txchat/imparse/proto/signal"
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
	cfg.ConnDealConsumerConfig.Topic = fmt.Sprintf("goim-%s-connection", cfg.AppID)
	cfg.ConnDealConsumerConfig.Group = fmt.Sprintf("goim-%s-connection-device", cfg.AppID)
	//new batch consumer
	consumer := xkafka.NewConsumer(cfg.ConnDealConsumerConfig, nil)
	logx.Info("dial kafka broker success")
	bc := xkafka.NewBatchConsumer(cfg.ConnDealBatchConsumerConf, xkafka.WithHandle(s.handleFunc), consumer)
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
		s.Error("logic.BizMsg proto.Unmarshal error", "err", err)
		return err
	}
	if receivedMsg.AppId != s.Config.AppID {
		return model.ErrAppID
	}
	return s.consumerOneConnect(ctx, receivedMsg)
}

func (s *Service) consumerOneConnect(ctx context.Context, m *logicclient.ReceivedMessage) error {
	switch m.Op {
	case protocol.Op_Auth:
		s.Info("user login with key")
		//将用户设备信息存入缓存
		dev, err := parseDevice(m)
		if err != nil {
			s.Error("parseDevice failed", "err", err)
		}
		if dev != nil {
			now := uint64(util.TimeNowUnixNano() / int64(time.Millisecond))
			err = s.svcCtx.Repo.AddDeviceInfo(&model.Device{
				Uid:         m.GetFrom(),
				ConnectId:   m.GetKey(),
				DeviceUuid:  dev.GetUuid(),
				DeviceType:  int32(dev.Device),
				DeviceName:  dev.GetDeviceName(),
				Username:    dev.Username,
				DeviceToken: dev.DeviceToken,
				IsEnabled:   false,
				AddTime:     now,
			})
			if err != nil {
				s.Error("AddDeviceInfo failed", "err", err)
			}
			//发送登录通知
			err = s.svcCtx.SignalHub.EndpointLogin(ctx, m.GetFrom(), &signal.SignalEndpointLogin{
				Uuid:       dev.GetUuid(),
				Device:     dev.Device,
				DeviceName: dev.GetDeviceName(),
				Datetime:   now,
			})
			if err != nil {
				s.Error("UniCastSignalEndpointLogin failed", "err", err)
			}
		}
		//连接群聊
		err = s.svcCtx.JoinGroups(ctx, m.GetFrom(), m.GetKey())
		if err != nil {
			s.Error("JoinGroups failed", "err", err)
		}
	case protocol.Op_Disconnect:
		s.Info("user logout with key")
		err := s.svcCtx.Repo.EnableDevice(m.GetFrom(), m.GetKey())
		if err != nil {
			s.Error("EnableDeviceInfo failed", "err", err)
		}
	default:
		return model.ErrCustomNotSupport
	}
	return nil
}

func parseDevice(m *logicclient.ReceivedMessage) (*auth.Login, error) {
	var p protocol.Proto
	err := proto.Unmarshal(m.GetBody(), &p)
	if err != nil {
		return nil, err
	}
	var authMsg protocol.AuthBody
	err = proto.Unmarshal(p.Body, &authMsg)
	if err != nil {
		return nil, err
	}
	if len(authMsg.Ext) == 0 {
		return nil, errors.New("ext is nil")
	}
	var device auth.Login
	err = proto.Unmarshal(authMsg.Ext, &device)
	if err != nil {
		return nil, err
	}
	return &device, nil
}
