package svc

import (
	"fmt"
	"time"

	"github.com/txchat/dtalk/app/services/answer/answerclient"
	"github.com/txchat/dtalk/app/services/device/deviceclient"
	"github.com/txchat/dtalk/app/services/pusher/internal/config"
	"github.com/txchat/dtalk/app/services/pusher/internal/dao"
	"github.com/txchat/dtalk/app/services/pusher/internal/publish"
	"github.com/txchat/dtalk/app/services/pusher/internal/recordhelper"
	"github.com/txchat/dtalk/app/services/pusher/internal/signal"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/naming"
	"github.com/txchat/dtalk/pkg/net/grpc"
	groupApi "github.com/txchat/dtalk/service/group/api"
	logic "github.com/txchat/im/api/logic/grpc"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc/resolver"
)

type ServiceContext struct {
	Config    config.Config
	Repo      dao.MessageRepository
	DeviceRPC deviceclient.Device
	AnswerRPC answerclient.Answer

	SignalNotice   *signal.Signal
	StoragePublish *publish.Storage
	OffPushPublish *publish.OffPush
	RecordHelper   *recordhelper.RecordHelper
	// will deprecate
	LogicRPC logic.LogicClient
	GroupRPC groupApi.GroupClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	answerRPC := answerclient.NewAnswer(zrpc.MustNewClient(c.AnswerRPC,
		zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor)))
	deviceRPC := deviceclient.NewDevice(zrpc.MustNewClient(c.DeviceRPC,
		zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor)))
	repo := dao.NewMessageRepositoryRedis(c.RedisDB)
	return &ServiceContext{
		Config:         c,
		Repo:           repo,
		DeviceRPC:      deviceRPC,
		AnswerRPC:      answerRPC,
		SignalNotice:   signal.NewSignal(answerRPC),
		StoragePublish: publish.NewStoragePublish(c.AppID, c.ProducerStorage),
		OffPushPublish: publish.NewOffPushPublish(c.AppID, c.ProducerOffPush),
		RecordHelper:   recordhelper.NewRecordHelper(repo),
		// will deprecate
		LogicRPC: newLogicClient(c),
		GroupRPC: newGroupClient(c),
	}
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
