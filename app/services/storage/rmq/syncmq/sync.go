package syncmq

import (
	"context"
	"strconv"

	"github.com/golang/protobuf/proto"
	"github.com/txchat/dtalk/app/services/pusher/pusherclient"
	"github.com/txchat/dtalk/internal/bizproto"
	"github.com/txchat/dtalk/internal/recordutil"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/im/api/protocol"
	"github.com/txchat/imparse"
	"github.com/txchat/imparse/proto/common"
	"github.com/txchat/imparse/proto/signal"
)

func (s *Service) pushClient(ctx context.Context, key, from, mid, target string, tp imparse.Channel, frameType imparse.FrameType, data []byte) error {
	_, err := s.svcCtx.PusherRPC.PushClient(ctx, &pusherclient.PushReq{
		Key:       key,
		From:      from,
		Mid:       mid,
		Target:    target,
		Data:      data,
		Type:      int32(tp),
		FrameType: string(frameType),
	})
	return err
}

func (s *Service) sendPrivateUnReadMsg(key, uid string) error {
	var p protocol.Proto
	p.Op = int32(protocol.Op_ReceiveMsg)
	p.Ver = 1
	p.Seq = 0
	bizP := &common.Proto{
		EventType: common.Proto_common,
	}

	records, err := s.svcCtx.Repo.UnReceiveMsg(uid)
	if err != nil {
		return err
	}
	for _, m := range records {
		eveP := &common.Common{
			ChannelType: common.Channel_ToUser,
			Mid:         util.MustToInt64(m.Mid),
			Seq:         m.Seq,
			From:        m.SenderId,
			Target:      m.ReceiverId,
			MsgType:     common.MsgType(m.MsgType),
			Msg:         recordutil.CommonMsgJSONDataToProtobufData(m.MsgType, []byte(m.Content)),
			Source:      recordutil.SourceJSONUnmarshal([]byte(m.Source)),
			Reference:   recordutil.ReferenceJSONUnmarshal([]byte(m.Reference)),
			Datetime:    m.CreateTime,
		}
		bizP.Body, err = proto.Marshal(eveP)
		p.Body, err = proto.Marshal(bizP)
		if err != nil {
			s.Error("Push msg Marshal failed", "err", err)
			continue
		}
		bytes, _ := proto.Marshal(&p)
		err = s.pushClient(context.TODO(), key, m.SenderId, strconv.FormatInt(eveP.Mid, 10), key, imparse.UniCast, bizproto.PrivateFrameType, bytes)
		if err != nil {
			s.Error("sendPrivateUnReadMsg PushClient failed", "toUID", uid, "mid", eveP.Mid, "err", err)
			return err
		}
	}
	return nil
}

func (s *Service) sendGroupUnReadMsg(key, uid string) error {
	var p protocol.Proto
	p.Op = int32(protocol.Op_ReceiveMsg)
	p.Ver = 1
	p.Seq = 0
	bizP := &common.Proto{
		EventType: common.Proto_common,
	}

	records, err := s.svcCtx.Repo.UnReceiveGroupMsg(uid)
	if err != nil {
		return err
	}
	for _, m := range records {
		eveP := &common.Common{
			ChannelType: common.Channel_ToGroup,
			Mid:         util.MustToInt64(m.Mid),
			Seq:         m.Seq,
			From:        m.SenderId,
			Target:      m.ReceiverId,
			MsgType:     common.MsgType(m.MsgType),
			Msg:         recordutil.CommonMsgJSONDataToProtobufData(m.MsgType, []byte(m.Content)),
			Source:      recordutil.SourceJSONUnmarshal([]byte(m.Source)),
			Reference:   recordutil.ReferenceJSONUnmarshal([]byte(m.Reference)),
			Datetime:    m.CreateTime,
		}
		bizP.Body, err = proto.Marshal(eveP)
		p.Body, err = proto.Marshal(bizP)
		if err != nil {
			s.Error("Push msg Marshal failed", "err", err)
			continue
		}
		bytes, _ := proto.Marshal(&p)
		err = s.pushClient(context.TODO(), key, m.SenderId, strconv.FormatInt(eveP.Mid, 10), key, imparse.GroupCast, bizproto.GroupFrameType, bytes)
		if err != nil {
			s.Error("sendGroupUnReadMsg PushClient failed", "toUID", uid, "mid", eveP.Mid, "err", err)
			return err
		}
	}
	return nil
}

func (s *Service) sendSignalUnReadMsg(key, uid string) error {
	var p protocol.Proto
	p.Op = int32(protocol.Op_ReceiveMsg)
	p.Ver = 1
	p.Seq = 0
	bizP := &common.Proto{
		EventType: common.Proto_Signal,
	}

	records, err := s.svcCtx.Repo.UnReceiveSignalMsg(uid)
	if err != nil {
		return err
	}
	for _, m := range records {
		eveP := &signal.Signal{
			Type: signal.SignalType(m.Type),
			Body: recordutil.SignalContentJSONDataToProtobufData(uint32(m.Type), []byte(m.Content)),
		}
		bizP.Body, err = proto.Marshal(eveP)
		p.Body, err = proto.Marshal(bizP)
		if err != nil {
			s.Error("Push msg Marshal failed", "err", err)
			continue
		}
		bytes, _ := proto.Marshal(&p)
		err = s.pushClient(context.TODO(), key, m.Uid, m.Id, key, imparse.UniCast, bizproto.SignalFrameType, bytes)
		if err != nil {
			s.Error("sendSignalUnReadMsg PushClient failed", "toUID", uid, "action", eveP.Type, "err", err)
			return err
		}
	}
	return nil
}

func (s *Service) SendUnReadMsg(key, uid string) error {
	//推送私聊离线消息
	err := s.sendPrivateUnReadMsg(key, uid)
	if err != nil {
		s.Error("sendPrivateUnReadMsg failed", "key", key, "uid", uid, "err", err)
		return err
	}
	//推送群聊离线消息
	err = s.sendGroupUnReadMsg(key, uid)
	if err != nil {
		s.Error("sendGroupUnReadMsg failed", "key", key, "uid", uid, "err", err)
		return err
	}
	//推送通知离线消息
	err = s.sendSignalUnReadMsg(key, uid)
	if err != nil {
		s.Error("sendSignalUnReadMsg failed", "key", key, "uid", uid, "err", err)
		return err
	}
	return nil
}

func (s *Service) CheckOnline(ctx context.Context, key string) (bool, error) {
	//TODO check client online
	return true, nil
}

func (s *Service) MarkReceived(tp, uid string, mid int64) error {
	switch tp {
	case string(bizproto.PrivateFrameType):
		_, _, err := s.svcCtx.Repo.MarkMsgReceived(uid, mid)
		return err
	case string(bizproto.GroupFrameType):
		_, _, err := s.svcCtx.Repo.MarkGroupMsgReceived(uid, mid)
		return err
	case string(bizproto.SignalFrameType):
		_, _, err := s.svcCtx.Repo.MarkSignalReceived(uid, mid)
		return err
	default:
		return nil
	}
}
