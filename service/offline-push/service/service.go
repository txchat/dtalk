package service

import (
	"fmt"
	"time"

	"github.com/inconshreveable/log15"
	offlinepush "github.com/txchat/dtalk/service/offline-push/api"
	"github.com/txchat/dtalk/service/offline-push/config"
	"github.com/txchat/dtalk/service/offline-push/model"
	"github.com/txchat/dtalk/service/offline-push/pusher"
	"github.com/txchat/dtalk/service/offline-push/pusher/android"
	"github.com/txchat/dtalk/service/offline-push/pusher/ios"
	"github.com/txchat/dtalk/service/offline-push/service/kafka"
)

type Service struct {
	log       log15.Logger
	cfg       *config.Config
	consumers map[string]*kafka.Consumer
	pushers   map[offlinepush.Device]pusher.IPusher
}

func New(c *config.Config) *Service {
	s := &Service{
		log:       log15.New("module", "offline-push/service"),
		cfg:       c,
		consumers: kafka.NewKafkaConsumers(c.AppId, c.MQSub, 0),
		pushers:   make(map[offlinepush.Device]pusher.IPusher),
	}
	s.loadPushers()
	return s
}

func (s Service) Config() *config.Config {
	return s.cfg
}

func (s *Service) ListenMQ() {
	for i, c := range s.consumers {
		s.log.Debug(fmt.Sprintf("accept %v", i))
		go c.Listen(s)
	}
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
	s.pushers[offlinepush.Device_Android] = androidCreator(pusher.Config{
		AppKey:          s.cfg.Pushers[android.Name].AppKey,
		AppMasterSecret: s.cfg.Pushers[android.Name].AppMasterSecret,
		MiActivity:      s.cfg.Pushers[android.Name].MiActivity,
		Environment:     s.cfg.Pushers[android.Name].Env,
	})
	s.pushers[offlinepush.Device_IOS] = iOSCreator(pusher.Config{
		AppKey:          s.cfg.Pushers[ios.Name].AppKey,
		AppMasterSecret: s.cfg.Pushers[ios.Name].AppMasterSecret,
		MiActivity:      s.cfg.Pushers[ios.Name].MiActivity,
		Environment:     s.cfg.Pushers[ios.Name].Env,
	})
}

func (s *Service) Deal(m *offlinepush.OffPushMsg) error {
	if m.AppId != s.cfg.AppId {
		return model.ErrAppId
	}

	p, ok := s.pushers[m.Device]
	if !ok {
		s.log.Error("pusher exec not find", "deviceType", m.Device.String(), "pushers", s.pushers)
		return model.ErrCustomNotSupport
	}
	if tm := time.Now().Unix(); tm > m.Timeout {
		s.log.Info("message offline push timeout",
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
