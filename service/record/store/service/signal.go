package service

import (
	"context"

	"github.com/txchat/dtalk/service/record/store/model"
	xproto "github.com/txchat/imparse/proto"
)

func (s *Service) AppendGroupCastSignal(mid, target string, signalType xproto.SignalType, content []byte, createTime uint64) error {
	//获取所有群成员
	members, err := s.dao.AllGroupMembers(context.TODO(), target)
	if err != nil {
		s.log.Error().Err(err).Msg("AllGroupMembers failed")
		return err
	}
	var signalItems = make([]*model.SignalContent, len(members))
	for i, member := range members {
		signalItems[i] = &model.SignalContent{
			Id:         mid,
			Uid:        member,
			Type:       uint8(signalType),
			State:      uint8(model.UnReceive),
			Content:    string(content),
			CreateTime: createTime,
			UpdateTime: createTime,
		}
	}
	_, _, err = s.dao.BatchAppendSignalContent(signalItems)
	if err != nil {
		s.log.Error().Err(err).Msg("BatchAppendSignalContent failed")
		return err
	}
	return nil
}

func (s *Service) AppendUniCastSignal(mid, target string, signalType xproto.SignalType, content []byte, createTime uint64) error {
	signal := &model.SignalContent{
		Id:         mid,
		Uid:        target,
		Type:       uint8(signalType),
		State:      uint8(model.UnReceive),
		Content:    string(content),
		CreateTime: createTime,
		UpdateTime: createTime,
	}
	_, _, err := s.dao.AppendSignalContent(signal)
	if err != nil {
		s.log.Error().Err(err).Interface("signal", signal).Msg("AppendSignalContent failed")
		return err
	}
	return nil
}
