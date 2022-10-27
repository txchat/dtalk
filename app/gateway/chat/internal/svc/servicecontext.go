package svc

import (
	"fmt"
	"time"

	"github.com/txchat/dtalk/app/gateway/chat/internal/config"
	"github.com/txchat/dtalk/app/gateway/chat/internal/middleware"
	"github.com/txchat/dtalk/app/gateway/chat/internal/middleware/authmock"
	"github.com/txchat/dtalk/app/services/answer/answerclient"
	"github.com/txchat/dtalk/app/services/call/callclient"
	"github.com/txchat/dtalk/app/services/storage/storageclient"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/naming"
	"github.com/txchat/dtalk/pkg/net/grpc"
	groupApi "github.com/txchat/dtalk/service/group/api"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc/resolver"
)

type ServiceContext struct {
	Config                   config.Config
	CallRPC                  callclient.Call
	AnswerRPC                answerclient.Answer
	StorageRPC               storageclient.Storage
	GroupRPC                 groupApi.GroupClient
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
		AnswerRPC: answerclient.NewAnswer(zrpc.MustNewClient(c.AnswerRPC,
			zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor))),
		StorageRPC: storageclient.NewStorage(zrpc.MustNewClient(c.StorageRPC,
			zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor))),
		GroupRPC: newGroupClient(c),
	}
}

func newGroupClient(cfg config.Config) groupApi.GroupClient {
	rb := naming.NewResolver(cfg.GroupRPCClient.RegAddrs, cfg.GroupRPCClient.Schema)
	resolver.Register(rb)

	addr := fmt.Sprintf("%s:///%s", cfg.GroupRPCClient.Schema, cfg.GroupRPCClient.SrvName) // "schema://[authority]/service"
	fmt.Println("rpc client call addr:", addr)

	conn, err := grpc.NewGRPCConn(addr, time.Duration(cfg.GroupRPCClient.Dial))
	if err != nil {
		panic(err)
	}
	return groupApi.NewGroupClient(conn)
}
