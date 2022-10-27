package ios

import "github.com/txchat/dtalk/service/offline-push/pusher"

const Name = "iOS"

func init() {
	pusher.Register(Name, New)
}

func New(cfg pusher.Config) pusher.IPusher {
	return &iOSPusher{
		AppKey:          cfg.AppKey,
		AppMasterSecret: cfg.AppMasterSecret,
		MiActivity:      cfg.MiActivity,
		environment:     cfg.Environment,
	}
}
