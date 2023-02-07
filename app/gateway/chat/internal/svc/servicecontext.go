package svc

import (
	"github.com/txchat/dtalk/app/gateway/chat/internal/config"
	"github.com/txchat/dtalk/app/gateway/chat/internal/middleware"
	"github.com/txchat/dtalk/app/gateway/chat/internal/middleware/authmock"
	"github.com/txchat/dtalk/app/services/answer/answerclient"
	"github.com/txchat/dtalk/app/services/call/callclient"
	"github.com/txchat/dtalk/app/services/device/deviceclient"
	"github.com/txchat/dtalk/app/services/group/groupclient"
	"github.com/txchat/dtalk/app/services/oss/ossclient"
	"github.com/txchat/dtalk/app/services/storage/storageclient"
	"github.com/txchat/dtalk/internal/signal"
	txchatSignalApi "github.com/txchat/dtalk/internal/signal/txchat"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config                   config.Config
	CallRPC                  callclient.Call
	AnswerRPC                answerclient.Answer
	StorageRPC               storageclient.Storage
	GroupRPC                 groupclient.Group
	OssRPC                   ossclient.Oss
	DeviceRPC                deviceclient.Device
	AppParseHeaderMiddleware rest.Middleware
	AppAuthMiddleware        rest.Middleware
	SignalHub                signal.Signal
}

func NewServiceContext(c config.Config) *ServiceContext {
	answerRPC := answerclient.NewAnswer(zrpc.MustNewClient(c.AnswerRPC,
		zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor), zrpc.WithNonBlock()))
	return &ServiceContext{
		Config: c,
		CallRPC: callclient.NewCall(zrpc.MustNewClient(c.CallRPC,
			zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor), zrpc.WithNonBlock())),
		AnswerRPC: answerRPC,
		StorageRPC: storageclient.NewStorage(zrpc.MustNewClient(c.StorageRPC,
			zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor), zrpc.WithNonBlock())),
		GroupRPC: groupclient.NewGroup(zrpc.MustNewClient(c.GroupRPC,
			zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor), zrpc.WithNonBlock())),
		OssRPC: ossclient.NewOss(zrpc.MustNewClient(c.OssRPC,
			zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor), zrpc.WithNonBlock())),
		DeviceRPC: deviceclient.NewDevice(zrpc.MustNewClient(c.DeviceRPC,
			zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor), zrpc.WithNonBlock())),
		AppParseHeaderMiddleware: middleware.NewAppParseHeaderMiddleware().Handle,
		AppAuthMiddleware:        middleware.NewAppAuthMiddleware(authmock.NewKVMock()).Handle,
		SignalHub:                txchatSignalApi.NewSignalHub(answerRPC),
	}
}
