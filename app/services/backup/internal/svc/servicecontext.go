package svc

import (
	"github.com/txchat/dtalk/app/services/backup/internal/config"
	"github.com/txchat/dtalk/app/services/backup/internal/dao"
)

type ServiceContext struct {
	Config config.Config
	Repo   dao.BackupRepository
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Repo:   dao.NewBackupRepositoryMysql(dao.NewDefaultConn(c.Mode, c.MySQL)),
	}
}
