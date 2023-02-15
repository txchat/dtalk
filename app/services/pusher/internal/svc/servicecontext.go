package svc

import (
	"context"
	"time"

	"github.com/golang/protobuf/proto"
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
	"github.com/txchat/dtalk/proto/offline"
	"github.com/txchat/im/app/logic/logicclient"
	"github.com/txchat/imparse/proto/auth"
	"github.com/txchat/imparse/proto/common"
	"github.com/zeromicro/go-zero/core/logx"
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

func (s *ServiceContext) PushOffline(ctx context.Context, app, from string, targets []string) {
	resp, err := s.DeviceRPC.GetUserAllDevices(ctx, &deviceclient.GetUserAllDevicesRequest{
		Uid: from,
	})
	if err != nil || resp == nil || len(resp.Devices) == 0 {
		log := logx.WithContext(ctx)
		log.Error("GetAllDevices failed", "from", from, "err", err)
		return
	}
	nickname := resp.Devices[0].Username

	//offline push
	for _, mid := range targets {
		err = s.pushAllDevice(ctx, app, from, nickname, mid)
		if err != nil {
			continue
		}
	}
}

func (s *ServiceContext) pushAllDevice(ctx context.Context, app, from, nickname, mid string) error {
	log := logx.WithContext(ctx)
	resp, err := s.DeviceRPC.GetUserAllDevices(ctx, &deviceclient.GetUserAllDevicesRequest{
		Uid: mid,
	})
	if err != nil {
		log.Error("GetAllDevices failed", "mid", mid, "err", err)
		return err
	}
	if resp == nil {
		return nil
	}
	for _, dev := range resp.Devices {
		if dev.IsEnabled && dev.DTUid == dev.Uid {
			//需要推送
			pushMsg := &offline.OffPushMsg{
				AppId:       app,
				Device:      auth.Device(dev.DeviceType),
				Title:       nickname,
				Content:     "[你收到一条消息]",
				Token:       dev.DeviceToken,
				ChannelType: int32(common.Channel_ToUser),
				Target:      from,
				Timeout:     time.Now().Add(time.Minute * 7).Unix(),
			}
			var msg []byte
			msg, err = proto.Marshal(pushMsg)
			if err != nil {
				log.Error("Marshal pushMsg failed", "err", err, "from", from, "appId", app, "toId", mid)
				continue
			}
			err = s.OffPushPublish.PublishOfflineMsg(ctx, mid, msg)
			if err != nil {
				log.Error("PublishOfflineMsg failed", "err", err, "from", from, "appId", app, "toId", mid)
			}
		}
	}
	return nil
}
