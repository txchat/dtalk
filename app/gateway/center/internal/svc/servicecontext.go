package svc

import (
	"github.com/txchat/dtalk/app/gateway/center/internal/config"
	"github.com/txchat/dtalk/app/gateway/center/internal/logic/backenduser"
	"github.com/txchat/dtalk/app/gateway/center/internal/middleware"
	"github.com/txchat/dtalk/app/services/backup/backupclient"
	"github.com/txchat/dtalk/app/services/version/versionclient"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config                   config.Config
	VersionRPC               versionclient.Version
	BackupRPC                backupclient.Backup
	UsersManager             *backenduser.UserManager
	AppParseHeaderMiddleware rest.Middleware
	AppAuthMiddleware        rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                   c,
		VersionRPC:               versionclient.NewVersion(zrpc.MustNewClient(c.VersionRPC)),
		BackupRPC:                backupclient.NewBackup(zrpc.MustNewClient(c.BackupRPC)),
		UsersManager:             backenduser.NewUserManager(c.Backend.Users),
		AppParseHeaderMiddleware: middleware.NewAppParseHeaderMiddleware().Handle,
		AppAuthMiddleware:        middleware.NewAppAuthMiddleware().Handle,
	}
}
