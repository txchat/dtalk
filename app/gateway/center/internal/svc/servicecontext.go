package svc

import (
	"github.com/txchat/dtalk/app/gateway/center/internal/config"
	"github.com/txchat/dtalk/app/gateway/center/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config             config.Config
	HTTPRespMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:             c,
		HTTPRespMiddleware: middleware.NewResponseMiddleware().Handle,
	}
}
