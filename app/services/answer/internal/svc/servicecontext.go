package svc

import (
	"github.com/txchat/dtalk/app/services/answer/internal/config"
	"github.com/txchat/dtalk/app/services/answer/internal/dao"
	"github.com/txchat/dtalk/app/services/answer/internal/msgfactory"
	"github.com/txchat/dtalk/app/services/generator/generatorclient"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/imparse"
	"github.com/txchat/imparse/chat"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Repo   dao.AnswerRepository

	IDGenRPC generatorclient.Generator

	Parser          chat.StandardParse
	MsgChecker      *msgfactory.Checker
	AnswerInter4rmq imparse.Answer
	AnswerInter4rpc imparse.Answer
}

func NewServiceContext(c config.Config) *ServiceContext {
	idGenRPC := generatorclient.NewGenerator(zrpc.MustNewClient(c.IDGenRPC,
		zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor)))
	svc := &ServiceContext{
		Config:   c,
		Repo:     dao.NewAnswerRepositoryRedis(c.RedisDB),
		IDGenRPC: idGenRPC,

		MsgChecker: msgfactory.NewChecker(),
	}
	msgCache := msgfactory.NewMsgCache(svc.Repo, svc.IDGenRPC)
	trace := msgfactory.NewTrace()
	fs := msgfactory.NewFilters()
	withoutAckCB := msgfactory.NewWithoutAckCallback()
	withCometAckCB := msgfactory.NewWithCometLevelAckCallback()

	svc.AnswerInter4rmq = imparse.NewStandardAnswer(msgCache, withCometAckCB, trace, fs.GetFilters())
	svc.AnswerInter4rpc = imparse.NewStandardAnswer(msgCache, withoutAckCB, trace, fs.GetFilters())
	return svc
}
