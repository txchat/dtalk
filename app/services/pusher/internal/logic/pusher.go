package logic

import (
	"context"
	"fmt"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/txchat/dtalk/app/services/device/deviceclient"
	"github.com/txchat/dtalk/app/services/pusher/internal/svc"
	"github.com/txchat/dtalk/internal/recordhelper"
	"github.com/txchat/dtalk/proto/record"
	offlinepush "github.com/txchat/dtalk/service/offline-push/api"
	comet "github.com/txchat/im/api/comet/grpc"
	logic "github.com/txchat/im/api/logic/grpc"
	"github.com/txchat/imparse/proto/common"
	"github.com/zeromicro/go-zero/core/logx"
)

type PusherLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPusherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PusherLogic {
	return &PusherLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PusherLogic) UniCastDevices(m *record.PushMsg) error {
	keysMsg := &logic.KeysMsg{
		AppId:  m.GetAppId(),
		ToKeys: []string{m.GetTarget()},
		Msg:    m.GetMsg(),
	}

	reply, err := l.svcCtx.LogicRPC.PushByKeys(l.ctx, keysMsg)
	if err != nil {
		return err
	}

	index := comet.PushMsgReply{}
	err = proto.Unmarshal(reply.Msg, &index)
	if err != nil {
		return fmt.Errorf("unmarshal PushMsgReply failed: %v", err)
	}
	item := &recordhelper.ConnSeqItem{
		Type:   m.GetFrameType(),
		Sender: m.GetFromId(),
		Client: m.GetKey(),
		Logs:   []int64{m.GetMid()},
	}
	for cid, seq := range index.Index {
		err := l.svcCtx.RecordHelper.Save(cid, seq, item)
		if err != nil {
			l.Error("AddConnSeqIndex failed",
				"err", err,
				"key", m.GetKey(),
				"from", m.GetFromId(),
				"cid", cid,
				"seq", seq,
			)
		}
	}
	return nil
}

func (l *PusherLogic) UniCast(m *record.PushMsg) error {
	midsMsg := &logic.MidsMsg{
		AppId: m.GetAppId(),
		ToIds: []string{m.GetTarget()},
		Msg:   m.GetMsg(),
	}

	if l.svcCtx.Config.OffPushEnabled {
		l.pushOffline(m, []string{m.GetTarget()})
	}

	//TODO 临时处理一下
	midsMsg.ToIds = []string{m.GetFromId()}
	_, err := l.svcCtx.LogicRPC.PushByMids(l.ctx, midsMsg)
	if err != nil {
		l.Error("UniCast PushByMids Failed",
			"err", err,
			"appId", midsMsg.GetAppId(),
			"toIds", midsMsg.GetToIds(),
			"len of msg", len(midsMsg.GetMsg()),
		)
	}

	midsMsg.ToIds = []string{m.GetTarget()}
	reply, err := l.svcCtx.LogicRPC.PushByMids(l.ctx, midsMsg)
	if err != nil {
		l.Error("UniCast PushByMids Failed",
			"err", err,
			"appId", midsMsg.GetAppId(),
			"toIds", midsMsg.GetToIds(),
			"len of msg", len(midsMsg.GetMsg()),
		)
		return err
	}

	index := comet.PushMsgReply{}
	err = proto.Unmarshal(reply.Msg, &index)
	if err != nil {
		return fmt.Errorf("unmarshal PushMsgReply failed: %v", err)
	}
	item := &recordhelper.ConnSeqItem{
		Type:   m.GetFrameType(),
		Sender: m.GetFromId(),
		Client: m.GetKey(),
		Logs:   []int64{m.GetMid()},
	}
	for cid, seq := range index.Index {
		err := l.svcCtx.RecordHelper.Save(cid, seq, item)
		if err != nil {
			l.Error("AddConnSeqIndex failed",
				"err", err,
				"key", m.GetKey(),
				"from", m.GetFromId(),
				"cid", cid,
				"seq", seq,
			)
		}
	}
	return nil
}

func (l *PusherLogic) GroupCast(m *record.PushMsg) error {
	gMsg := &logic.GroupMsg{
		AppId: m.GetAppId(),
		Group: m.GetTarget(),
		Msg:   m.GetMsg(),
	}

	reply, err := l.svcCtx.LogicRPC.PushGroup(l.ctx, gMsg)
	if err != nil {
		l.Error("GroupCast PushGroup Failed",
			"err", err,
			"appId", gMsg.GetAppId(),
			"to group", gMsg.GetGroup(),
			"len of msg", len(gMsg.GetMsg()),
		)
		return err
	}

	index := comet.PushMsgReply{}
	err = proto.Unmarshal(reply.Msg, &index)
	if err != nil {
		return fmt.Errorf("unmarshal PushMsgReply failed: %v", err)
	}
	item := &recordhelper.ConnSeqItem{
		Type:   m.GetFrameType(),
		Sender: m.GetFromId(),
		Client: m.GetKey(),
		Logs:   []int64{m.GetMid()},
	}
	for cid, seq := range index.Index {
		err := l.svcCtx.RecordHelper.SaveProxy(cid, seq, item)
		if err != nil {
			l.Error("AddConnSeqIndex failed",
				"err", err,
				"key", m.GetKey(),
				"from", m.GetFromId(),
				"cid", cid,
				"seq", seq,
			)
		}
	}
	return nil
}

func (l *PusherLogic) pushOffline(m *record.PushMsg, toIds []string) {
	resp, err := l.svcCtx.DeviceRPC.GetUserAllDevices(context.TODO(), &deviceclient.GetUserAllDevicesRequest{
		Uid: m.GetFromId(),
	})
	if err != nil || resp == nil || len(resp.Devices) == 0 {
		l.Error("GetAllDevices failed",
			"err", err,
			"key", m.GetKey(),
			"from", m.GetFromId(),
		)
		return
	}
	nickname := resp.Devices[0].Username

	//offline push
	for _, mid := range toIds {
		err := l.pushAllDevice(m, nickname, mid)
		if err != nil {
			continue
		}
	}
}

func (l *PusherLogic) pushAllDevice(m *record.PushMsg, nickname, mid string) error {
	resp, err := l.svcCtx.DeviceRPC.GetUserAllDevices(context.TODO(), &deviceclient.GetUserAllDevicesRequest{
		Uid: mid,
	})
	if err != nil {
		l.Error("GetAllDevices failed",
			"err", err,
			"key", m.GetKey(),
			"from", m.GetFromId(),
			"mid", mid,
		)
		return err
	}
	if resp == nil {
		return nil
	}
	for _, dev := range resp.Devices {
		if dev.IsEnabled && dev.DTUid == dev.Uid {
			//需要推送
			pushMsg := &offlinepush.OffPushMsg{
				AppId:       m.GetAppId(),
				Device:      offlinepush.Device(dev.DeviceType),
				Title:       nickname,
				Content:     "[你收到一条消息]",
				Token:       dev.DeviceToken,
				ChannelType: int32(common.Channel_ToUser),
				Target:      m.GetFromId(),
				Timeout:     time.Now().Add(time.Minute * 7).Unix(),
			}
			b, err := proto.Marshal(pushMsg)
			if err != nil {
				l.Error("Marshal pushMsg failed",
					"err", err,
					"key", m.GetKey(),
					"from", m.GetFromId(),
					"appId", m.GetAppId(),
					"toId", mid,
				)
				continue
			}
			err = l.svcCtx.OffPushPublish.PublishOfflineMsg(context.TODO(), m.GetKey(), b)
			if err != nil {
				l.Error("PublishOfflineMsg failed",
					"err", err,
					"key", m.GetKey(),
					"from", m.GetFromId(),
					"appId", m.GetAppId(),
					"toId", mid,
				)
			}
		}
	}
	return nil
}
