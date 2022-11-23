package svc

import (
	"github.com/txchat/dtalk/app/gateway/center/internal/config"
	"github.com/txchat/dtalk/app/gateway/center/internal/logic/backenduser"
	"github.com/txchat/dtalk/app/gateway/center/internal/middleware"
	"github.com/txchat/dtalk/app/gateway/center/internal/middleware/authmock"
	"github.com/txchat/dtalk/app/services/backup/backupclient"
	"github.com/txchat/dtalk/app/services/version/versionclient"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/notify"
	"github.com/txchat/dtalk/pkg/notify/debug"
	"github.com/txchat/dtalk/pkg/notify/email"
	"github.com/txchat/dtalk/pkg/notify/sms"
	"github.com/txchat/dtalk/pkg/notify/whitelist"
	"github.com/txchat/dtalk/service/backup/model"
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

	SmsValidate   notify.Validate
	EmailValidate notify.Validate
}

func NewServiceContext(c config.Config) *ServiceContext {
	var smsValidate model.Validate
	smsValidate = sms.NewSMS(c.SMS.Surl, c.SMS.AppKey, c.SMS.SecretKey, c.SMS.Msg)
	if c.SMS.Env == model.Debug {
		smsValidate = debug.NewDebugValidate(debug.GetMockCode(c.SMS.Msg), smsValidate)
	}

	var emailValidate model.Validate
	emailValidate = email.NewEmail(c.Email.Surl, c.Email.AppKey, c.Email.SecretKey, c.Email.Msg)
	if c.Email.Env == model.Debug {
		emailValidate = debug.NewDebugValidate(debug.GetMockCode(c.Email.Msg), emailValidate)
	}

	return &ServiceContext{
		Config:                   c,
		UsersManager:             backenduser.NewUserManager(c.Backend.Users),
		AppParseHeaderMiddleware: middleware.NewAppParseHeaderMiddleware().Handle,
		AppAuthMiddleware:        middleware.NewAppAuthMiddleware(authmock.NewKVMock()).Handle,
		VersionRPC: versionclient.NewVersion(zrpc.MustNewClient(c.VersionRPC,
			zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor))),
		BackupRPC: backupclient.NewBackup(zrpc.MustNewClient(c.BackupRPC,
			zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor))),
		SmsValidate:   whitelist.NewWhitelistValidate(c.Whitelist, smsValidate),
		EmailValidate: whitelist.NewWhitelistValidate(c.Whitelist, emailValidate),
	}
}
