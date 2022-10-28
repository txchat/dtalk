package svc

import (
	"github.com/txchat/dtalk/app/services/answer/answerclient"
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
	IDGenRPC       generatorclient.Generator
	AnswerRPC      answerclient.Answer
	SessionCreator *xcall.SessionCreator
	SignalNotify   xcall.SignalNotify
	TicketCreator  xcall.TicketCreator
}

func NewServiceContext(c config.Config) *ServiceContext {
	idGenRPC := generatorclient.NewGenerator(zrpc.MustNewClient(c.IDGenRPC,
		zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor)))
	answerRPC := answerclient.NewAnswer(zrpc.MustNewClient(c.AnswerRPC,
		zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor)))
	return &ServiceContext{
		Config:         c,
		Repo:           dao.NewCallRepositoryRedis(c.RedisDB),
		IDGenRPC:       idGenRPC,
		SessionCreator: xcall.NewSessionCreator(c.RTC.CallingTimeout, rpcidgen.NewIDGenerator(idGenRPC), roomidgen.NewRoomIDGen(0)), // TODO roomid gen object
		SignalNotify:   rpcnotify.NewCallNotifyClient(answerRPC),                                                                    // TODO answer rpc object
		TicketCreator:  sign.NewCloudSDK(tencentyun.NewTCTLSSig(c.RTC.SDKAppId, c.RTC.SecretKey, c.RTC.Expire)).GetTicket,
	}
}
