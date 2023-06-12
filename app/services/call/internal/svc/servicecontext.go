package svc

import (
	"github.com/txchat/dtalk/app/services/call/internal/config"
	"github.com/txchat/dtalk/app/services/call/internal/dao"
	"github.com/txchat/dtalk/app/services/generator/generatorclient"
	"github.com/txchat/dtalk/app/services/transfer/transferclient"
	xcall "github.com/txchat/dtalk/internal/call"
	"github.com/txchat/dtalk/internal/call/roomidgen"
	"github.com/txchat/dtalk/internal/call/rpcidgen"
	"github.com/txchat/dtalk/internal/call/sign"
	"github.com/txchat/dtalk/internal/call/sign/tencentyun"
	"github.com/txchat/dtalk/internal/signal"
	txchatSignalApi "github.com/txchat/dtalk/internal/signal/txchat"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	Repo           dao.CallRepository
	SessionCreator *xcall.SessionCreator
	SignalHub      signal.Signal
	TicketCreator  xcall.TicketCreator
}

func NewServiceContext(c config.Config) *ServiceContext {
	idGenRPC := generatorclient.NewGenerator(zrpc.MustNewClient(c.IDGenRPC,
		zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor), zrpc.WithNonBlock()))
	transferRPC := transferclient.NewTransfer(zrpc.MustNewClient(c.TransferRPC,
		zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor), zrpc.WithNonBlock()))
	return &ServiceContext{
		Config:         c,
		Repo:           dao.NewCallRepositoryRedis(c.RedisDB),
		SessionCreator: xcall.NewSessionCreator(c.RTC.CallingTimeout, rpcidgen.NewIDGenerator(idGenRPC), roomidgen.NewRoomIDGen(0)),
		SignalHub:      txchatSignalApi.NewSignalHub(transferRPC),
		TicketCreator:  sign.NewCloudSDK(tencentyun.NewTCTLSSig(c.RTC.SDKAppId, c.RTC.SecretKey, c.RTC.Expire)).GetTicket,
	}
}
