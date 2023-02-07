package ios

import pusher "github.com/txchat/dtalk/pkg/third-part-push"

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
