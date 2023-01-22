package svc

import (
	"github.com/txchat/dtalk/app/services/answer/answerclient"
	"github.com/txchat/dtalk/app/services/device/deviceclient"
	"github.com/txchat/dtalk/app/services/group/groupclient"
	"github.com/txchat/dtalk/app/services/pusher/internal/config"
	"github.com/txchat/dtalk/app/services/pusher/internal/dao"
	"github.com/txchat/dtalk/app/services/pusher/internal/publish"
	"github.com/txchat/dtalk/internal/recordhelper"
	"github.com/txchat/dtalk/internal/signal"
	txchatSignalApi "github.com/txchat/dtalk/internal/signal/txchat"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/im/app/logic/logicclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	Repo      dao.MessageRepository
	DeviceRPC deviceclient.Device
	GroupRPC  groupclient.Group
	LogicRPC  logicclient.Logic

	SignalHub      signal.Signal
	StoragePublish *publish.Storage
	OffPushPublish *publish.OffPush
	RecordHelper   *recordhelper.RecordHelper
}

func NewServiceContext(c config.Config) *ServiceContext {
	answerRPC := answerclient.NewAnswer(zrpc.MustNewClient(c.AnswerRPC,
		zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor), zrpc.WithNonBlock()))
	repo := dao.NewMessageRepositoryRedis(c.RedisDB)
	return &ServiceContext{
		Config: c,
		Repo:   repo,
		DeviceRPC: deviceclient.NewDevice(zrpc.MustNewClient(c.DeviceRPC,
			zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor), zrpc.WithNonBlock())),
		GroupRPC: groupclient.NewGroup(zrpc.MustNewClient(c.GroupRPC,
			zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor), zrpc.WithNonBlock())),
		LogicRPC: logicclient.NewLogic(zrpc.MustNewClient(c.LogicRPC,
			zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor), zrpc.WithNonBlock())),
		SignalHub:      txchatSignalApi.NewSignalHub(answerRPC),
		StoragePublish: publish.NewStoragePublish(c.AppID, c.ProducerStorage),
		OffPushPublish: publish.NewOffPushPublish(c.AppID, c.ProducerOffPush),
		RecordHelper:   recordhelper.NewRecordHelper(repo),
	}
}
