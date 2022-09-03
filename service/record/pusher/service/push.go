package service

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog"
	"github.com/txchat/dtalk/pkg/util"
	device "github.com/txchat/dtalk/service/device/api"

	"github.com/golang/protobuf/proto"
	offlinepush "github.com/txchat/dtalk/service/offline-push/api"
	record "github.com/txchat/dtalk/service/record/proto"
	"github.com/txchat/dtalk/service/record/pusher/logH"
	comet "github.com/txchat/im/api/comet/grpc"
	logic "github.com/txchat/im/api/logic/grpc"
	xproto "github.com/txchat/imparse/proto"
)

//just socket push
func (s *Service) PushUser(ctx context.Context, key, from, mid, target string, tp int32, frameType string, body []byte) error {
	log := zerolog.Ctx(ctx).With().Str("Uid", from).Str("ConnId", key).
		Str("Mid", mid).Str("Target", target).Int32("Push Type", tp).Str("Frame Type", frameType).Logger()
	err := s.pusher.UniCast(ctx, &record.PushMsg{
		AppId:     s.cfg.AppId,
		FromId:    from,
		Mid:       util.MustToInt64(mid),
		Key:       key,
		Target:    target,
		Msg:       body,
		Type:      tp,
		FrameType: frameType,
	})
	if err != nil {
		log.Error().Stack().Err(err).Msg("PushUser failed")
	}
	return nil
}

//just socket push
func (s *Service) PushClient(ctx context.Context, key, from, mid, target string, tp int32, frameType string, body []byte) error {
	log := zerolog.Ctx(ctx).With().Str("Uid", from).Str("ConnId", key).
		Str("Mid", mid).Str("Target", target).Int32("Push Type", tp).Str("Frame Type", frameType).Logger()
	err := s.pusher.UniCastDevices(ctx, &record.PushMsg{
		AppId:     s.cfg.AppId,
		FromId:    from,
		Mid:       util.MustToInt64(mid),
		Key:       key,
		Target:    target,
		Msg:       body,
		Type:      tp,
		FrameType: frameType,
	})
	if err != nil {
		log.Error().Stack().Err(err).Msg("PushClient failed")
	}
	return nil
}

//
type Pusher struct {
	s *Service
}

func (p *Pusher) UniCastDevices(ctx context.Context, m *record.PushMsg) error {
	keysMsg := &logic.KeysMsg{
		AppId:  m.GetAppId(),
		ToKeys: []string{m.GetTarget()},
		Msg:    m.GetMsg(),
	}

	reply, err := p.s.logicClient.PushByKeys(ctx, keysMsg)
	if err != nil {
		return err
	}

	index := comet.PushMsgReply{}
	err = proto.Unmarshal(reply.Msg, &index)
	if err != nil {
		return fmt.Errorf("unmarshal PushMsgReply failed: %v", err)
	}
	item := &logH.ConnSeqItem{
		Type:   m.GetFrameType(),
		Sender: m.GetFromId(),
		Client: m.GetKey(),
		Logs:   []int64{m.GetMid()},
	}
	for cid, seq := range index.Index {
		err := p.s.lh.Save(cid, seq, item)
		if err != nil {
			p.s.log.Warn().Err(err).Str("key", m.GetKey()).Str("from", m.GetFromId()).Str("cid", cid).Int32("seq", seq).Msg("AddConnSeqIndex failed")
		}
	}
	return nil
}

func (p *Pusher) UniCast(ctx context.Context, m *record.PushMsg) error {
	midsMsg := &logic.MidsMsg{
		AppId: m.GetAppId(),
		ToIds: []string{m.GetTarget()},
		Msg:   m.GetMsg(),
	}

	if p.s.cfg.OffPush.IsEnabled {
		p.pushOffline(m, []string{m.GetTarget()})
	}

	//TODO 临时处理一下
	midsMsg.ToIds = []string{m.GetFromId()}
	_, err := p.s.logicClient.PushByMids(ctx, midsMsg)
	if err != nil {
		p.s.log.Debug().Err(err).
			Str("appId", midsMsg.GetAppId()).Strs("toIds", midsMsg.GetToIds()).Int("len of msg", len(midsMsg.GetMsg())).
			Msg("UniCast PushByMids Failed")
	}

	midsMsg.ToIds = []string{m.GetTarget()}
	reply, err := p.s.logicClient.PushByMids(ctx, midsMsg)
	if err != nil {
		p.s.log.Debug().Err(err).
			Str("appId", midsMsg.GetAppId()).Strs("toIds", midsMsg.GetToIds()).Int("len of msg", len(midsMsg.GetMsg())).
			Msg("UniCast PushByMids Failed")
		return err
	}
	p.s.log.Debug().Str("appId", midsMsg.GetAppId()).Strs("to ids", midsMsg.GetToIds()).Msg("UniCast success")

	index := comet.PushMsgReply{}
	err = proto.Unmarshal(reply.Msg, &index)
	if err != nil {
		return fmt.Errorf("unmarshal PushMsgReply failed: %v", err)
	}
	item := &logH.ConnSeqItem{
		Type:   m.GetFrameType(),
		Sender: m.GetFromId(),
		Client: m.GetKey(),
		Logs:   []int64{m.GetMid()},
	}
	for cid, seq := range index.Index {
		err := p.s.lh.Save(cid, seq, item)
		if err != nil {
			p.s.log.Warn().Err(err).Str("key", m.GetKey()).Str("from", m.GetFromId()).Str("cid", cid).Int32("seq", seq).Msg("AddConnSeqIndex failed")
		}
	}
	return nil
}

func (p *Pusher) GroupCast(ctx context.Context, m *record.PushMsg) error {
	gMsg := &logic.GroupMsg{
		AppId: m.GetAppId(),
		Group: m.GetTarget(),
		Msg:   m.GetMsg(),
	}

	reply, err := p.s.logicClient.PushGroup(ctx, gMsg)
	if err != nil {
		p.s.log.Debug().Err(err).
			Str("appId", gMsg.GetAppId()).Str("to group", gMsg.GetGroup()).Int("len of msg", len(gMsg.GetMsg())).
			Msg("GroupCast PushGroup Failed")
		return err
	}
	p.s.log.Debug().Str("appId", gMsg.AppId).Str("to group", gMsg.Group).Msg("GroupCast success")

	index := comet.PushMsgReply{}
	err = proto.Unmarshal(reply.Msg, &index)
	if err != nil {
		return fmt.Errorf("unmarshal PushMsgReply failed: %v", err)
	}
	item := &logH.ConnSeqItem{
		Type:   m.GetFrameType(),
		Sender: m.GetFromId(),
		Client: m.GetKey(),
		Logs:   []int64{m.GetMid()},
	}
	for cid, seq := range index.Index {
		err := p.s.lh.SaveProxy(cid, seq, item)
		if err != nil {
			p.s.log.Warn().Err(err).Str("key", m.GetKey()).Str("from", m.GetFromId()).Str("cid", cid).Int32("seq", seq).Msg("AddConnSeqIndex failed")
		}
	}
	return nil
}

func (p *Pusher) pushOffline(m *record.PushMsg, toIds []string) {
	resp, err := p.s.deviceClient.GetUserAllDevices(context.TODO(), &device.GetUserAllDevicesRequest{
		Uid: m.GetFromId(),
	})
	if err != nil || resp == nil || len(resp.Devices) == 0 {
		p.s.log.Error().Err(err).Str("key", m.GetKey()).Str("from", m.GetFromId()).Msg("GetAllDevices failed")
		return
	}
	nickname := resp.Devices[0].Username

	//offline push
	for _, mid := range toIds {
		err := p.pushAllDevice(m, nickname, mid)
		if err != nil {
			continue
		}
	}
}

func (p *Pusher) pushAllDevice(m *record.PushMsg, nickname, mid string) error {
	resp, err := p.s.deviceClient.GetUserAllDevices(context.TODO(), &device.GetUserAllDevicesRequest{
		Uid: mid,
	})
	if err != nil {
		p.s.log.Error().Err(err).Str("key", m.GetKey()).Str("from", m.GetFromId()).Str("mid", mid).Msg("GetAllDevices failed")
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
				ChannelType: int32(xproto.Channel_ToUser),
				Target:      m.GetFromId(),
				Timeout:     time.Now().Add(time.Minute * 7).Unix(),
			}
			b, err := proto.Marshal(pushMsg)
			if err != nil {
				p.s.log.Error().Err(err).Str("key", m.GetKey()).Str("from", m.GetFromId()).Str("appId", m.GetAppId()).Interface("toId", mid).Msg("Marshal pushMsg failed")
				continue
			}
			err = p.s.dao.PublishOfflineMsg(context.TODO(), m.GetKey(), b)
			if err != nil {
				p.s.log.Error().Err(err).Str("key", m.GetKey()).Str("from", m.GetFromId()).Str("appId", m.GetAppId()).Interface("toId", mid).Msg("PublishOfflineMsg failed")
			}
		}
	}
	return nil
}
