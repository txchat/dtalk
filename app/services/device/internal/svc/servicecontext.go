package svc

import (
	"github.com/txchat/dtalk/app/services/device/internal/config"
	"github.com/txchat/dtalk/app/services/device/internal/dao"
)

type ServiceContext struct {
	Config config.Config
	Repo   dao.DeviceRepository
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Repo:   dao.NewDeviceRepositoryRedis(c.RedisDB),
	}
}
