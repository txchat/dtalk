package svc

import (
	"github.com/txchat/dtalk/app/gateway/center/internal/config"
	"github.com/txchat/dtalk/app/gateway/center/internal/logic/backenduser"
	"github.com/txchat/dtalk/app/gateway/center/internal/middleware"
	"github.com/txchat/dtalk/app/gateway/center/internal/middleware/authmock"
	"github.com/txchat/dtalk/app/services/backup/backupclient"
	"github.com/txchat/dtalk/app/services/version/versionclient"
	xerror "github.com/txchat/dtalk/pkg/error"
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
		UsersManager:             backenduser.NewUserManager(c.Backend.Users),
		AppParseHeaderMiddleware: middleware.NewAppParseHeaderMiddleware().Handle,
		AppAuthMiddleware:        middleware.NewAppAuthMiddleware(authmock.NewKVMock()).Handle,
		VersionRPC: versionclient.NewVersion(zrpc.MustNewClient(c.VersionRPC,
			zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor))),
		BackupRPC: backupclient.NewBackup(zrpc.MustNewClient(c.BackupRPC,
			zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor))),
	}
}
