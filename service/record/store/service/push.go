package service

import (
	"context"
	"strconv"

	"github.com/txchat/imparse"
	"github.com/txchat/imparse/chat"

	"github.com/golang/protobuf/proto"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/dtalk/service/record/store/model"
	comet "github.com/txchat/im/api/comet/grpc"
	xproto "github.com/txchat/imparse/proto"
)

func (s *Service) sendPrivateUnReadMsg(key, uid string) error {
	var p comet.Proto
	p.Op = int32(comet.Op_ReceiveMsg)
	p.Ver = 1
	p.Seq = 0
	bizP := &xproto.Proto{
		EventType: xproto.Proto_common,
	}

	records, err := s.dao.UnReceiveMsg(uid)
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

func (s *Service) sendGroupUnReadMsg(key, uid string) error {
	var p comet.Proto
	p.Op = int32(comet.Op_ReceiveMsg)
	p.Ver = 1
	p.Seq = 0
	bizP := &xproto.Proto{
		EventType: xproto.Proto_common,
	}

	records, err := s.dao.UnReceiveGroupMsg(uid)
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

func (s *Service) sendSignalUnReadMsg(key, uid string) error {
	var p comet.Proto
	p.Op = int32(comet.Op_ReceiveMsg)
	p.Ver = 1
	p.Seq = 0
	bizP := &xproto.Proto{
		EventType: xproto.Proto_Signal,
	}

	records, err := s.dao.UnReceiveSignalMsg(uid)
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

func (s *Service) SendUnReadMsg(key, uid string) error {
	//推送私聊离线消息
	err := s.sendPrivateUnReadMsg(key, uid)
	if err != nil {
		s.log.Error().Err(err).Str("key", key).Str("uid", uid).Msg("sendPrivateUnReadMsg failed")
		return err
	}
	//推送群聊离线消息
	err = s.sendGroupUnReadMsg(key, uid)
	if err != nil {
		s.log.Error().Err(err).Str("key", key).Str("uid", uid).Msg("sendGroupUnReadMsg failed")
		return err
	}
	//推送通知离线消息
	err = s.sendSignalUnReadMsg(key, uid)
	if err != nil {
		s.log.Error().Err(err).Str("key", key).Str("uid", uid).Msg("sendSignalUnReadMsg failed")
		return err
	}
	return nil
}
