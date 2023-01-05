package svc

import (
	"fmt"
	"time"

	"github.com/txchat/dtalk/app/services/group/groupclient"

	"github.com/txchat/dtalk/app/services/answer/internal/config"
	"github.com/txchat/dtalk/app/services/answer/internal/dao"
	"github.com/txchat/dtalk/app/services/answer/internal/msgfactory"
	"github.com/txchat/dtalk/app/services/generator/generatorclient"
	xerror "github.com/txchat/dtalk/pkg/error"
	xkafka "github.com/txchat/dtalk/pkg/mq/kafka"
	"github.com/txchat/dtalk/pkg/naming"
	"github.com/txchat/dtalk/pkg/net/grpc"
	logic "github.com/txchat/im/api/logic/grpc"
	"github.com/txchat/imparse"
	"github.com/txchat/imparse/chat"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc/resolver"
)

type ServiceContext struct {
	Config     config.Config
	Repo       dao.AnswerRepository
	IDGenRPC   generatorclient.Generator
	MsgChecker *msgfactory.Checker

	AnswerInter4rmq imparse.Answer
	AnswerInter4rpc imparse.Answer
	//need not init
	Parser chat.StandardParse

	//todo will deprecate
	LogicRPC logic.LogicClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	groupRPC := groupclient.NewGroup(zrpc.MustNewClient(c.GroupRPC,
		zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor)))
	idGenRPC := generatorclient.NewGenerator(zrpc.MustNewClient(c.IDGenRPC,
		zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor)))
	svc := &ServiceContext{
		Config:     c,
		Repo:       dao.NewAnswerRepositoryRedis(c.RedisDB),
		IDGenRPC:   idGenRPC,
		MsgChecker: msgfactory.NewChecker(),
		// will deprecate
		LogicRPC: newLogicClient(c),
	}
	// answerInter components
	msgCache := msgfactory.NewMsgCache(svc.Repo, svc.IDGenRPC)
	trace := msgfactory.NewTrace()
	fs := msgfactory.NewFilters(groupRPC)
	c.Producer.Topic = fmt.Sprintf("received-%s-topic", c.AppID)
	pub := xkafka.NewProducer(c.Producer)
	withoutAckCB := msgfactory.NewWithoutAckCallback(c.AppID, pub)
	withCometAckCB := msgfactory.NewWithCometLevelAckCallback(c.AppID, pub, svc.LogicRPC)

	svc.AnswerInter4rmq = imparse.NewStandardAnswer(msgCache, withCometAckCB, trace, fs.GetFilters())
	svc.AnswerInter4rpc = imparse.NewStandardAnswer(msgCache, withoutAckCB, trace, fs.GetFilters())
	return svc
}

func newLogicClient(cfg config.Config) logic.LogicClient {
	rb := naming.NewResolver(cfg.LogicRPCClient.RegAddrs, cfg.LogicRPCClient.Schema)
	resolver.Register(rb)

	addr := fmt.Sprintf("%s:///%s", cfg.LogicRPCClient.Schema, cfg.LogicRPCClient.SrvName) // "schema://[authority]/service"
	fmt.Println("logic rpc client call addr:", addr)

	conn, err := grpc.NewGRPCConn(addr, time.Duration(cfg.LogicRPCClient.Dial))
	if err != nil {
		panic(err)
	}
	return logic.NewLogicClient(conn)
}
