package svc

import (
	"fmt"
	"time"

	"github.com/txchat/dtalk/app/services/storage/internal/exec"

	"github.com/txchat/dtalk/app/services/device/deviceclient"
	"github.com/txchat/dtalk/app/services/pusher/pusherclient"
	"github.com/txchat/dtalk/app/services/storage/internal/config"
	"github.com/txchat/dtalk/app/services/storage/internal/dao"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/naming"
	"github.com/txchat/dtalk/pkg/net/grpc"
	groupApi "github.com/txchat/dtalk/service/group/api"
	"github.com/txchat/imparse"
	"github.com/txchat/imparse/chat"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc/resolver"
)

type ServiceContext struct {
	Config config.Config
	Repo   dao.StorageRepository
	//need not init
	Parser      chat.StandardParse
	StorageExec imparse.Storage
	DeviceRPC   deviceclient.Device
	PusherRPC   pusherclient.Pusher
	GroupRPC    groupApi.GroupClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	deviceRPC := deviceclient.NewDevice(zrpc.MustNewClient(c.DeviceRPC,
		zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor)))
	pusherRPC := pusherclient.NewPusher(zrpc.MustNewClient(c.PusherRPC,
		zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor)))
	s := &ServiceContext{
		Config:    c,
		Repo:      dao.NewUniRepository(c.RedisDB, c.MySQL),
		DeviceRPC: deviceRPC,
		PusherRPC: pusherRPC,
		GroupRPC:  newGroupClient(c),
	}
	s.StorageExec = imparse.NewStandardStorage(exec.NewStorageExec(s))
	return s
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
