package svc

import (
	"github.com/txchat/im/app/logic/logicclient"

	"github.com/txchat/dtalk/app/services/group/groupclient"

	"github.com/txchat/dtalk/app/services/answer/internal/config"
	"github.com/txchat/dtalk/app/services/answer/internal/dao"
	"github.com/txchat/dtalk/app/services/answer/internal/msgfactory"
	"github.com/txchat/dtalk/app/services/generator/generatorclient"
	"github.com/txchat/dtalk/internal/bizproto"
	xerror "github.com/txchat/dtalk/pkg/error"
	xkafka "github.com/txchat/dtalk/pkg/mq/kafka"
	"github.com/txchat/imparse"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	Repo       dao.AnswerRepository
	MsgChecker *msgfactory.Checker

	AnswerInter4rmq imparse.Answer
	AnswerInter4rpc imparse.Answer
	//need not init
	Parser bizproto.StandardParse

	IDGenRPC generatorclient.Generator
	LogicRPC logicclient.Logic
}

func NewServiceContext(c config.Config) *ServiceContext {
	groupRPC := groupclient.NewGroup(zrpc.MustNewClient(c.GroupRPC,
		zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor), zrpc.WithNonBlock()))
	svc := &ServiceContext{
		Config:     c,
		Repo:       dao.NewAnswerRepositoryRedis(c.RedisDB),
		MsgChecker: msgfactory.NewChecker(),
		IDGenRPC: generatorclient.NewGenerator(zrpc.MustNewClient(c.IDGenRPC,
			zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor), zrpc.WithNonBlock())),
		LogicRPC: logicclient.NewLogic(zrpc.MustNewClient(c.LogicRPC,
			zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor), zrpc.WithNonBlock())),
	}
	// answerInter components
	msgCache := msgfactory.NewMsgCache(svc.Repo, svc.IDGenRPC)
	trace := msgfactory.NewTrace()
	fs := msgfactory.NewFilters(groupRPC)
	pub := xkafka.NewProducer(c.Producer)
	withoutAckCB := msgfactory.NewWithoutAckCallback(c.AppID, pub)
	withCometAckCB := msgfactory.NewWithCometLevelAckCallback(c.AppID, pub, svc.LogicRPC)

	svc.AnswerInter4rmq = imparse.NewStandardAnswer(msgCache, withCometAckCB, trace, fs.GetFilters())
	svc.AnswerInter4rpc = imparse.NewStandardAnswer(msgCache, withoutAckCB, trace, fs.GetFilters())
	return svc
}
