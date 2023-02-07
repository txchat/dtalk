package svc

import (
	"github.com/txchat/dtalk/app/services/device/deviceclient"
	"github.com/txchat/dtalk/app/services/group/groupclient"
	"github.com/txchat/dtalk/app/services/pusher/pusherclient"
	"github.com/txchat/dtalk/app/services/storage/internal/config"
	"github.com/txchat/dtalk/app/services/storage/internal/dao"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Repo   dao.StorageRepository
	//need not init
	DeviceRPC deviceclient.Device
	PusherRPC pusherclient.Pusher
	GroupRPC  groupclient.Group
}

func NewServiceContext(c config.Config) *ServiceContext {
	s := &ServiceContext{
		Config: c,
		Repo:   dao.NewUniRepository(c.RedisDB, c.MySQL),
		DeviceRPC: deviceclient.NewDevice(zrpc.MustNewClient(c.DeviceRPC,
			zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor), zrpc.WithNonBlock())),
		PusherRPC: pusherclient.NewPusher(zrpc.MustNewClient(c.PusherRPC,
			zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor), zrpc.WithNonBlock())),
		GroupRPC: groupclient.NewGroup(zrpc.MustNewClient(c.GroupRPC,
			zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor), zrpc.WithNonBlock())),
	}
	return s
}
