package android

import pusher "github.com/txchat/dtalk/pkg/third-part-push"

const Name = "Android"

func init() {
	pusher.Register(Name, New)
}

func New(cfg pusher.Config) pusher.IPusher {
	return &androidPusher{
		AppKey:          cfg.AppKey,
		AppMasterSecret: cfg.AppMasterSecret,
		MiActivity:      cfg.MiActivity,
		environment:     cfg.Environment,
	}
}
