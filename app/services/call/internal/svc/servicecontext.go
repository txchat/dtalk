package svc

import (
	"github.com/txchat/dtalk/app/services/call/internal/config"
	"github.com/txchat/dtalk/app/services/call/internal/dao"
	"github.com/txchat/dtalk/app/services/generator/generatorclient"
	xcall "github.com/txchat/dtalk/pkg/call"
	"github.com/txchat/dtalk/pkg/call/roomidgen"
	"github.com/txchat/dtalk/pkg/call/rpcidgen"
	"github.com/txchat/dtalk/pkg/call/rpcnotify"
	"github.com/txchat/dtalk/pkg/call/sign"
	"github.com/txchat/dtalk/pkg/call/sign/tencentyun"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	Repo           dao.CallRepository
	SessionCreator *xcall.SessionCreator
	SignalNotify   xcall.SignalNotify
	RTC            sign.TLSSig
	IDGenRPC       generatorclient.Generator
}

func NewServiceContext(c config.Config) *ServiceContext {
	idGenRPC := generatorclient.NewGenerator(zrpc.MustNewClient(c.IDGenRPC,
		zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor)))
	return &ServiceContext{
		Config:         c,
		Repo:           dao.NewCallRepositoryRedis(c.RedisDB),
		SessionCreator: xcall.NewSessionCreator(c.RTC.CallingTimeout, rpcidgen.NewIDGenerator(idGenRPC), roomidgen.NewRoomIDGen(0)), // TODO roomid gen object
		SignalNotify:   rpcnotify.NewCallNotifyClient(nil),                                                                          // TODO answer rpc object
		RTC:            tencentyun.NewTCTLSSig(c.RTC.SDKAppId, c.RTC.SecretKey, c.RTC.Expire),
		IDGenRPC:       idGenRPC,
	}
}
