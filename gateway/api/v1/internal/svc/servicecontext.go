package svc

import (
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/interceptor/trace"
	device "github.com/txchat/dtalk/service/device/api"
	group "github.com/txchat/dtalk/service/group/api"
	store "github.com/txchat/dtalk/service/record/store/api"
	"google.golang.org/grpc"
	"sync"
	"time"

	"github.com/txchat/dtalk/gateway/api/v1/internal/config"
	answer "github.com/txchat/dtalk/service/record/answer/api"
)

// ServiceContext 服务上下文
type ServiceContext struct {
	m sync.RWMutex
	c config.Config

	AnswerClient *answer.Client
	StoreClient  *store.Client
	GroupClient  *group.Client
	DeviceClient *device.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	sc := &ServiceContext{
		c:            c,
		AnswerClient: answer.New(c.AnswerRPCClient.RegAddrs, c.AnswerRPCClient.Schema, c.AnswerRPCClient.SrvName, time.Duration(c.AnswerRPCClient.Dial)),
		StoreClient:  store.New(c.StoreRPCClient.RegAddrs, c.StoreRPCClient.Schema, c.StoreRPCClient.SrvName, time.Duration(c.StoreRPCClient.Dial)),
		DeviceClient: device.New(c.DeviceRPCClient.RegAddrs, c.DeviceRPCClient.Schema, c.DeviceRPCClient.SrvName, time.Duration(c.DeviceRPCClient.Dial)),
		GroupClient: group.New(c.GroupRPCClient.RegAddrs,
			c.GroupRPCClient.Schema,
			c.GroupRPCClient.SrvName,
			time.Duration(c.GroupRPCClient.Dial),
			grpc.WithChainUnaryInterceptor(xerror.ErrClientInterceptor, trace.UnaryClientInterceptor),
		),
	}

	return sc
}

// Config 获取全局配置
func (sc *ServiceContext) Config() config.Config {
	sc.m.RLock()
	defer sc.m.RUnlock()

	return sc.c
}
