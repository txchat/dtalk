package mq

import (
	"context"
	"fmt"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/txchat/dtalk/app/services/offline/internal/config"
	"github.com/txchat/dtalk/app/services/offline/internal/model"
	"github.com/txchat/dtalk/app/services/offline/internal/svc"
	pusher "github.com/txchat/dtalk/pkg/third-part-push"
	"github.com/txchat/dtalk/pkg/third-part-push/android"
	"github.com/txchat/dtalk/pkg/third-part-push/ios"
	"github.com/txchat/dtalk/proto/offline"
	"github.com/txchat/imparse/proto/auth"
	xkafka "github.com/txchat/pkg/mq/kafka"
	"github.com/zeromicro/go-zero/core/logx"
)

type Service struct {
	logx.Logger
	Config        config.Config
	svcCtx        *svc.ServiceContext
	batchConsumer *xkafka.BatchConsumer
	pushers       map[auth.Device]pusher.IPusher
}

func NewService(cfg config.Config, svcCtx *svc.ServiceContext) *Service {
	s := &Service{
		Logger:  logx.WithContext(context.TODO()),
		Config:  cfg,
		svcCtx:  svcCtx,
		pushers: make(map[auth.Device]pusher.IPusher),
	}
	s.loadPushers()
	//topic config
	cfg.DealConsumerConfig.Topic = fmt.Sprintf("offpush-%s-topic", cfg.AppID)
	cfg.DealConsumerConfig.Group = fmt.Sprintf("offpush-%s-group", cfg.AppID)
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
	bizMsg := new(offline.OffPushMsg)
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

func (s *Service) DealPush(ctx context.Context, m *offline.OffPushMsg) error {
	p, ok := s.pushers[m.Device]
	if !ok {
		s.Error("pusher exec not find", "deviceType", m.Device.String(), "pushers", s.pushers)
		return model.ErrCustomNotSupport
	}
	if tm := time.Now().Unix(); tm > m.Timeout {
		s.Info("message offline push timeout",
			"appId", m.GetAppId(),
			"deviceType", m.GetDevice().String(),
			"deviceToken", m.GetToken(),
			"channelType", m.GetChannelType(),
			"target", m.GetTarget(),
			"timeout time", m.GetTimeout(),
			"now time", tm,
		)
		return nil
	}
	return p.SinglePush(m.Token, m.Title, m.Content, &pusher.Extra{
		Address:     m.Target,
		ChannelType: m.ChannelType,
		TimeOutTime: m.Timeout,
	})
}

func (s *Service) loadPushers() {
	androidCreator, err := pusher.Load(android.Name)
	if err != nil {
		panic(err)
	}
	iOSCreator, err := pusher.Load(ios.Name)
	if err != nil {
		panic(err)
	}
	s.pushers[auth.Device_Android] = androidCreator(pusher.Config{
		AppKey:          s.svcCtx.Config.Pushers[android.Name].AppKey,
		AppMasterSecret: s.svcCtx.Config.Pushers[android.Name].AppMasterSecret,
		MiActivity:      s.svcCtx.Config.Pushers[android.Name].MiActivity,
		Environment:     s.svcCtx.Config.Pushers[android.Name].Env,
	})
	s.pushers[auth.Device_IOS] = iOSCreator(pusher.Config{
		AppKey:          s.svcCtx.Config.Pushers[ios.Name].AppKey,
		AppMasterSecret: s.svcCtx.Config.Pushers[ios.Name].AppMasterSecret,
		MiActivity:      s.svcCtx.Config.Pushers[ios.Name].MiActivity,
		Environment:     s.svcCtx.Config.Pushers[ios.Name].Env,
	})
}
