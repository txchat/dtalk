package svc

import (
	"context"
	"strconv"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/txchat/dtalk/api/proto/chat"
	"github.com/txchat/dtalk/api/proto/signal"
	"github.com/txchat/dtalk/app/services/device/internal/config"
	"github.com/txchat/dtalk/app/services/device/internal/dao"
	"github.com/txchat/dtalk/app/services/group/groupclient"
	"github.com/txchat/dtalk/app/services/pusher/pusherclient"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/im/app/logic/logicclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Repo   dao.DeviceRepository

	LogicRPC     logicclient.Logic
	GroupRPC     groupclient.Group
	pusherClient pusherclient.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Repo:   dao.NewDeviceRepositoryRedis(c.RedisDB),
		GroupRPC: groupclient.NewGroup(zrpc.MustNewClient(c.GroupRPC,
			zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor), zrpc.WithNonBlock())),
		LogicRPC: logicclient.NewLogic(zrpc.MustNewClient(c.LogicRPC,
			zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor), zrpc.WithNonBlock())),
		pusherClient: pusherclient.NewPusher(zrpc.MustNewClient(c.PusherRPC,
			zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor), zrpc.WithNonBlock())),
	}
}

func (s *ServiceContext) getAllJoinedGroups(ctx context.Context, uid string) (groups []int64, err error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*15))
	defer cancel()
	reply, err := s.GroupRPC.JoinedGroups(ctx, &groupclient.JoinedGroupsReq{
		Uid: uid,
	})
	if err != nil {
		return
	}

	return reply.GetGid(), nil
}

func (s *ServiceContext) JoinGroups(ctx context.Context, uid, key string) error {
	groups, err := s.getAllJoinedGroups(ctx, uid)
	if err != nil {
		return err
	}

	if len(groups) == 0 {
		return nil
	}

	var gid = make([]string, len(groups))
	for i, group := range groups {
		gid[i] = strconv.FormatInt(group, 10)
	}
	_, err = s.LogicRPC.JoinGroupByKey(ctx, &logicclient.JoinGroupByKeyReq{
		AppId: s.Config.AppID,
		Key:   []string{key},
		Gid:   gid,
	})
	return err
}

func (s *ServiceContext) EndpointLogin(ctx context.Context, uid string, action *signal.SignalEndpointLogin) error {
	body, err := proto.Marshal(action)
	if err != nil {
		return err
	}
	sig := &signal.Signal{
		Type: signal.SignalType_EndpointLogin,
		Body: body,
	}
	sigData, err := proto.Marshal(sig)
	if err != nil {
		return err
	}
	signalEvent := &chat.Chat{
		Type: chat.Chat_signal,
		Seq:  0,
		Body: sigData,
	}
	data, err := proto.Marshal(signalEvent)
	if err != nil {
		return err
	}

	_, err = s.pusherClient.PushList(ctx, &pusherclient.PushListReq{
		App:  s.Config.AppID,
		From: uid,
		Uid:  []string{uid},
		Body: data,
	})
	return err
}
