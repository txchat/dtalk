package service

import (
	"context"
	"math"
	"strconv"

	"github.com/txchat/imparse"
	"github.com/txchat/imparse/chat"

	"github.com/golang/protobuf/proto"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/dtalk/service/record/store/model"
	comet "github.com/txchat/im/api/comet/grpc"
	xproto "github.com/txchat/imparse/proto"
)

func (s *Service) sendPrivateSyncMsg(key, uid string, startMid int64) error {
	var p comet.Proto
	p.Op = int32(comet.Op_ReceiveMsg)
	p.Ver = 1
	p.Seq = 0
	bizP := &xproto.Proto{
		EventType: xproto.Proto_common,
	}

	records, err := s.dao.UserMsgAfter(uid, startMid, math.MaxInt64)
	if err != nil {
		return err
	}
	for _, m := range records {
		eveP := &xproto.Common{
			ChannelType: xproto.Channel_ToUser,
			Mid:         util.MustToInt64(m.Mid),
			Seq:         m.Seq,
			From:        m.SenderId,
			Target:      m.ReceiverId,
			MsgType:     xproto.MsgType(m.MsgType),
			Msg:         model.ConvertMsg(m.MsgType, []byte(m.Content)),
			Source:      model.ConvertSource([]byte(m.Source)),
			Reference:   model.ConvertReference([]byte(m.Reference)),
			Datetime:    m.CreateTime,
		}
		bizP.Body, err = proto.Marshal(eveP)
		p.Body, err = proto.Marshal(bizP)
		if err != nil {
			s.log.Error().Err(err).Msg("Push msg Marshal failed")
			continue
		}
		bytes, _ := proto.Marshal(&p)
		err = s.dao.PushClient(context.TODO(), key, m.SenderId, strconv.FormatInt(eveP.Mid, 10), key, imparse.UniCast, chat.PrivateFrameType, bytes)
		if err != nil {
			s.log.Error().Err(err).Str("toUid", uid).Int64("mid", eveP.Mid).Msg("sendPrivateUnReadMsg PushClient failed")
			return err
		}
	}
	return nil
}

func (s *Service) sendGroupSyncMsg(key, uid string, startMid int64) error {
	var p comet.Proto
	p.Op = int32(comet.Op_ReceiveMsg)
	p.Ver = 1
	p.Seq = 0
	bizP := &xproto.Proto{
		EventType: xproto.Proto_common,
	}

	records, err := s.dao.GroupMsgAfter(uid, startMid, math.MaxInt64)
	if err != nil {
		return err
	}
	for _, m := range records {
		eveP := &xproto.Common{
			ChannelType: xproto.Channel_ToGroup,
			Mid:         util.MustToInt64(m.Mid),
			Seq:         m.Seq,
			From:        m.SenderId,
			Target:      m.ReceiverId,
			MsgType:     xproto.MsgType(m.MsgType),
			Msg:         model.ConvertMsg(m.MsgType, []byte(m.Content)),
			Source:      model.ConvertSource([]byte(m.Source)),
			Reference:   model.ConvertReference([]byte(m.Reference)),
			Datetime:    m.CreateTime,
		}
		bizP.Body, err = proto.Marshal(eveP)
		p.Body, err = proto.Marshal(bizP)
		if err != nil {
			s.log.Error().Err(err).Msg("Push msg Marshal failed")
			continue
		}
		bytes, _ := proto.Marshal(&p)
		err = s.dao.PushClient(context.TODO(), key, m.SenderId, strconv.FormatInt(eveP.Mid, 10), key, imparse.GroupCast, chat.GroupFrameType, bytes)
		if err != nil {
			s.log.Error().Err(err).Str("toUid", uid).Int64("mid", eveP.Mid).Msg("sendGroupUnReadMsg PushClient failed")
			return err
		}
	}
	return nil
}

func (s *Service) sendSignalSyncMsg(key, uid string, startMid int64) error {
	var p comet.Proto
	p.Op = int32(comet.Op_ReceiveMsg)
	p.Ver = 1
	p.Seq = 0
	bizP := &xproto.Proto{
		EventType: xproto.Proto_Signal,
	}

	records, err := s.dao.SyncSignalMsg(uid, startMid, math.MaxInt64)
	if err != nil {
		return err
	}
	for _, m := range records {
		eveP := &xproto.Signal{
			Type: xproto.SignalType(m.Type),
			Body: model.ConvertSignal(uint32(m.Type), []byte(m.Content)),
		}
		bizP.Body, err = proto.Marshal(eveP)
		p.Body, err = proto.Marshal(bizP)
		if err != nil {
			s.log.Error().Err(err).Msg("Push msg Marshal failed")
			continue
		}
		bytes, _ := proto.Marshal(&p)
		err = s.dao.PushClient(context.TODO(), key, m.Uid, m.Id, key, imparse.UniCast, chat.SignalFrameType, bytes)
		if err != nil {
			s.log.Error().Err(err).Str("toUid", uid).Interface("action", eveP.Type).Msg("sendSignalUnReadMsg PushClient failed")
			return err
		}
	}
	return nil
}

func (s *Service) SendSyncMsg(key, uid string, startMid int64) error {
	//推送私聊同步消息
	err := s.sendPrivateSyncMsg(key, uid, startMid)
	if err != nil {
		s.log.Error().Err(err).Str("key", key).Str("uid", uid).Msg("sendPrivateUnReadMsg failed")
		return err
	}
	//推送群聊同步消息
	err = s.sendGroupSyncMsg(key, uid, startMid)
	if err != nil {
		s.log.Error().Err(err).Str("key", key).Str("uid", uid).Msg("sendGroupUnReadMsg failed")
		return err
	}
	//推送通知同步消息
	err = s.sendSignalSyncMsg(key, uid, startMid)
	if err != nil {
		s.log.Error().Err(err).Str("key", key).Str("uid", uid).Msg("sendSignalUnReadMsg failed")
		return err
	}
	return nil
}

func (s *Service) GetSyncMsg(key, uid string, startMid, count int64) ([][]byte, error) {
	var p comet.Proto
	p.Op = int32(comet.Op_ReceiveMsg)
	p.Ver = 1
	p.Seq = 0
	bizP := &xproto.Proto{
		EventType: xproto.Proto_common,
	}

	//私聊
	uMsg, err := s.dao.UserMsgAfter(uid, startMid, count)
	if err != nil {
		return nil, err
	}

	//群聊
	gMsg, err := s.dao.GroupMsgAfter(uid, startMid, count)
	if err != nil {
		return nil, err
	}

	var records = make([][]byte, len(uMsg)+len(gMsg))
	var num = 0
	for _, m := range uMsg {
		eveP := &xproto.Common{
			ChannelType: xproto.Channel_ToUser,
			Mid:         util.MustToInt64(m.Mid),
			Seq:         m.Seq,
			From:        m.SenderId,
			Target:      m.ReceiverId,
			MsgType:     xproto.MsgType(m.MsgType),
			Msg:         model.ConvertMsg(m.MsgType, []byte(m.Content)),
			Source:      model.ConvertSource([]byte(m.Source)),
			Reference:   model.ConvertReference([]byte(m.Reference)),
			Datetime:    m.CreateTime,
		}
		bizP.Body, err = proto.Marshal(eveP)
		p.Body, err = proto.Marshal(bizP)
		if err != nil {
			s.log.Error().Err(err).Msg("Push msg Marshal failed")
			continue
		}
		bytes, _ := proto.Marshal(&p)
		records[num] = bytes
		num++
	}

	for _, m := range gMsg {
		eveP := &xproto.Common{
			ChannelType: xproto.Channel_ToGroup,
			Mid:         util.MustToInt64(m.Mid),
			Seq:         m.Seq,
			From:        m.SenderId,
			Target:      m.ReceiverId,
			MsgType:     xproto.MsgType(m.MsgType),
			Msg:         model.ConvertMsg(m.MsgType, []byte(m.Content)),
			Source:      model.ConvertSource([]byte(m.Source)),
			Reference:   model.ConvertReference([]byte(m.Reference)),
			Datetime:    m.CreateTime,
		}
		bizP.Body, err = proto.Marshal(eveP)
		p.Body, err = proto.Marshal(bizP)
		if err != nil {
			s.log.Error().Err(err).Msg("Push msg Marshal failed")
			continue
		}
		bytes, _ := proto.Marshal(&p)
		records[num] = bytes
		num++
	}
	return records[0:num], nil
}

func (s *Service) GetSyncMsgJustBizLevel(key, uid string, startMid, count int64) ([][]byte, error) {
	bizP := &xproto.Proto{
		EventType: xproto.Proto_common,
	}

	//私聊 DESC
	uMsg, err := s.dao.UserMsgAfter(uid, startMid, count)
	if err != nil {
		return nil, err
	}

	//to ASC
	for i := 0; i < len(uMsg)/2; i++ {
		uMsg[i], uMsg[len(uMsg)-1-i] = uMsg[len(uMsg)-1-i], uMsg[i]
	}

	//群聊
	gMsg, err := s.dao.GroupMsgAfter(uid, startMid, count)
	if err != nil {
		return nil, err
	}
	//to ASC
	for i := 0; i < len(gMsg)/2; i++ {
		gMsg[i], gMsg[len(gMsg)-1-i] = gMsg[len(gMsg)-1-i], gMsg[i]
	}

	var records = make([][]byte, len(uMsg)+len(gMsg))
	var num = 0
	for _, m := range uMsg {
		eveP := &xproto.Common{
			ChannelType: xproto.Channel_ToUser,
			Mid:         util.MustToInt64(m.Mid),
			Seq:         m.Seq,
			From:        m.SenderId,
			Target:      m.ReceiverId,
			MsgType:     xproto.MsgType(m.MsgType),
			Msg:         model.ConvertMsg(m.MsgType, []byte(m.Content)),
			Source:      model.ConvertSource([]byte(m.Source)),
			Reference:   model.ConvertReference([]byte(m.Reference)),
			Datetime:    m.CreateTime,
		}
		bizP.Body, err = proto.Marshal(eveP)
		bytes, err := proto.Marshal(bizP)
		if err != nil {
			s.log.Error().Err(err).Msg("Push msg Marshal failed")
			continue
		}
		records[num] = bytes
		num++
	}

	for _, m := range gMsg {
		eveP := &xproto.Common{
			ChannelType: xproto.Channel_ToGroup,
			Mid:         util.MustToInt64(m.Mid),
			Seq:         m.Seq,
			From:        m.SenderId,
			Target:      m.ReceiverId,
			MsgType:     xproto.MsgType(m.MsgType),
			Msg:         model.ConvertMsg(m.MsgType, []byte(m.Content)),
			Source:      model.ConvertSource([]byte(m.Source)),
			Reference:   model.ConvertReference([]byte(m.Reference)),
			Datetime:    m.CreateTime,
		}
		bizP.Body, err = proto.Marshal(eveP)
		bytes, err := proto.Marshal(bizP)
		if err != nil {
			s.log.Error().Err(err).Msg("Push msg Marshal failed")
			continue
		}
		records[num] = bytes
		num++
	}
	return records[0:num], nil
}
