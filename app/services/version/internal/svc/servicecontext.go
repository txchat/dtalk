package svc

import (
	"github.com/txchat/dtalk/app/services/version/internal/config"
	"github.com/txchat/dtalk/app/services/version/internal/dao"
)

type ServiceContext struct {
	Config config.Config
	Repo   dao.DeviceRepository
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Repo:   dao.NewVersionRepositoryMysql(c.Mode, c.MySQL),
	}
}
