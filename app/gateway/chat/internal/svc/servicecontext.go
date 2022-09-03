package svc

import (
	"github.com/txchat/dtalk/app/gateway/chat/internal/config"
	"github.com/txchat/dtalk/app/gateway/chat/internal/middleware"
	"github.com/txchat/dtalk/app/gateway/chat/internal/middleware/authmock"
	"github.com/txchat/dtalk/app/services/call/callclient"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config                   config.Config
	CallRPC                  callclient.Call
	AppParseHeaderMiddleware rest.Middleware
	AppAuthMiddleware        rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                   c,
		AppParseHeaderMiddleware: middleware.NewAppParseHeaderMiddleware().Handle,
		AppAuthMiddleware:        middleware.NewAppAuthMiddleware(authmock.NewKVMock()).Handle,
		CallRPC: callclient.NewCall(zrpc.MustNewClient(c.CallRPC,
			zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor))),
	}
}
