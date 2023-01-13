package svc

import (
	"fmt"
	"time"

	"github.com/txchat/dtalk/app/services/group/groupclient"
	"github.com/txchat/dtalk/internal/recordhelper"
	"github.com/txchat/dtalk/internal/signal"
	txchatSignalApi "github.com/txchat/dtalk/internal/signal/txchat"

	"github.com/txchat/dtalk/app/services/answer/answerclient"
	"github.com/txchat/dtalk/app/services/device/deviceclient"
	"github.com/txchat/dtalk/app/services/pusher/internal/config"
	"github.com/txchat/dtalk/app/services/pusher/internal/dao"
	"github.com/txchat/dtalk/app/services/pusher/internal/publish"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/naming"
	"github.com/txchat/dtalk/pkg/net/grpc"
	logic "github.com/txchat/im/api/logic/grpc"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc/resolver"
)

type ServiceContext struct {
	Config    config.Config
	Repo      dao.MessageRepository
	DeviceRPC deviceclient.Device
	GroupRPC  groupclient.Group

	SignalHub      signal.Signal
	StoragePublish *publish.Storage
	OffPushPublish *publish.OffPush
	RecordHelper   *recordhelper.RecordHelper
	// will deprecate
	LogicRPC logic.LogicClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	answerRPC := answerclient.NewAnswer(zrpc.MustNewClient(c.AnswerRPC,
		zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor), zrpc.WithNonBlock()))
	deviceRPC := deviceclient.NewDevice(zrpc.MustNewClient(c.DeviceRPC,
		zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor), zrpc.WithNonBlock()))
	repo := dao.NewMessageRepositoryRedis(c.RedisDB)
	return &ServiceContext{
		Config:    c,
		Repo:      repo,
		DeviceRPC: deviceRPC,
		GroupRPC: groupclient.NewGroup(zrpc.MustNewClient(c.GroupRPC,
			zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor), zrpc.WithNonBlock())),
		SignalHub:      txchatSignalApi.NewSignalHub(answerRPC),
		StoragePublish: publish.NewStoragePublish(c.AppID, c.ProducerStorage),
		OffPushPublish: publish.NewOffPushPublish(c.AppID, c.ProducerOffPush),
		RecordHelper:   recordhelper.NewRecordHelper(repo),
		// will deprecate
		LogicRPC: newLogicClient(c),
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
