package svc

import (
	"context"
	"time"

	"github.com/txchat/dtalk/api/proto/auth"
	"github.com/txchat/dtalk/api/proto/message"
	"github.com/txchat/dtalk/app/services/device/deviceclient"
	"github.com/txchat/dtalk/app/services/group/groupclient"
	"github.com/txchat/dtalk/app/services/offline/internal/config"
	"github.com/txchat/dtalk/app/services/offline/internal/model"
	"github.com/txchat/dtalk/internal/proto/offline"
	pusher "github.com/txchat/dtalk/pkg/third-part-push"
	"github.com/txchat/dtalk/pkg/third-part-push/android"
	"github.com/txchat/dtalk/pkg/third-part-push/ios"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/zeromicro/go-zero/core/logx"
)

type ServiceContext struct {
	Config    config.Config
	DeviceRPC deviceclient.Device
	GroupRPC  groupclient.Group
	pushers   map[auth.Device]pusher.IPusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	svc := &ServiceContext{
		Config:  c,
		pushers: make(map[auth.Device]pusher.IPusher),
	}
	svc.loadPushers()
	return svc
}

func (s *ServiceContext) loadPushers() {
	androidCreator, err := pusher.Load(android.Name)
	if err != nil {
		panic(err)
	}
	iOSCreator, err := pusher.Load(ios.Name)
	if err != nil {
		panic(err)
	}
	s.pushers[auth.Device_Android] = androidCreator(pusher.Config{
		AppKey:          s.Config.Pushers[android.Name].AppKey,
		AppMasterSecret: s.Config.Pushers[android.Name].AppMasterSecret,
		MiActivity:      s.Config.Pushers[android.Name].MiActivity,
		Environment:     s.Config.Pushers[android.Name].Env,
	})
	s.pushers[auth.Device_IOS] = iOSCreator(pusher.Config{
		AppKey:          s.Config.Pushers[ios.Name].AppKey,
		AppMasterSecret: s.Config.Pushers[ios.Name].AppMasterSecret,
		MiActivity:      s.Config.Pushers[ios.Name].MiActivity,
		Environment:     s.Config.Pushers[ios.Name].Env,
	})
}

func (s *ServiceContext) PushOffline(ctx context.Context, msg *offline.ThirdPartyPushMQ) error {
	notification, err := s.sessionInformation(ctx, msg)
	if err != nil {
		return err
	}
	//offline push
	for _, mid := range msg.GetTarget() {
		err = s.pushAllDevice(ctx, mid, msg, notification)
		if err != nil {
			continue
		}
	}
	return nil
}

func (s *ServiceContext) sessionInformation(ctx context.Context, msg *offline.ThirdPartyPushMQ) (*pusher.Notification, error) {
	var notification pusher.Notification
	switch msg.GetChannelType() {
	case message.Channel_Private:
		resp, err := s.DeviceRPC.GetUserAllDevices(ctx, &deviceclient.GetUserAllDevicesRequest{
			Uid: msg.GetFrom(),
		})
		if err != nil {
			return &notification, err
		}
		if resp == nil || len(resp.GetDevices()) == 0 {
			return &notification, model.ErrDeviceInfoNotFound
		}
		nickname := resp.GetDevices()[0].GetUsername()
		notification.Title = nickname
	case message.Channel_Group:
		gid := util.MustToInt64(msg.GetSession())
		groupInfoResp, err := s.GroupRPC.GroupInfo(ctx, &groupclient.GroupInfoReq{
			Gid: gid,
		})
		if err != nil {
			return &notification, err
		}
		notification.Title = groupInfoResp.GetGroup().GetName()

		memberInfoResp, err := s.GroupRPC.MemberInfo(ctx, &groupclient.MemberInfoReq{
			Gid: gid,
			Uid: msg.GetFrom(),
		})
		if err != nil {
			return &notification, err
		}
		notification.Subtitle = memberInfoResp.GetMember().GetNickname()
	}
	notification.Body = msg.GetContent()
	return &notification, nil
}

func (s *ServiceContext) pushAllDevice(ctx context.Context, mid string, msg *offline.ThirdPartyPushMQ, notification *pusher.Notification) error {
	log := logx.WithContext(ctx)
	resp, err := s.DeviceRPC.GetUserAllDevices(ctx, &deviceclient.GetUserAllDevicesRequest{
		Uid: mid,
	})
	if err != nil {
		return err
	}
	if resp == nil {
		return model.ErrDeviceInfoNotFound
	}
	for _, dev := range resp.Devices {
		if dev.IsEnabled && dev.DTUid == dev.Uid {
			//需要推送
			p, ok := s.pushers[auth.Device(dev.DeviceType)]
			if !ok {
				log.Errorf("pusher exec %s not find", auth.Device(dev.DeviceType).String())
				continue
			}
			deadline := util.UnixToTime(msg.GetDatetime()).Add(s.Config.Timeout)
			if time.Now().After(deadline) {
				continue
			}
			err = p.SinglePush(dev.DeviceToken, *notification, &pusher.Extra{
				SessionKey:  msg.GetSession(),
				ChannelType: int32(msg.GetChannelType()),
				TimeOutTime: deadline.UnixMilli(),
			})
			if err != nil {
				log.Error(err)
				continue
			}
		}
	}
	return nil
}
