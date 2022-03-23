package service

import (
	"github.com/txchat/dtalk/service/record/store/model"
	xproto "github.com/txchat/imparse/proto"
)

// GetPriRecord 所有聊天记录
func (s *Service) DelRecord(tp xproto.Channel, mid int64) error {
	switch tp {
	case xproto.Channel_ToUser:
		_, _, err := s.dao.DelMsgContent(mid)
		return err
	case xproto.Channel_ToGroup:
		_, _, err := s.dao.DelGroupMsgContent(mid)
		return err
	default:
		return model.ErrRecordNotFind
	}
}

// GetPriRecord 所有聊天记录
func (s *Service) GetSpecifyRecord(tp xproto.Channel, mid int64) (*model.MsgContent, error) {
	switch tp {
	case xproto.Channel_ToUser:
		return s.dao.GetSpecifyRecord(mid)
	case xproto.Channel_ToGroup:
		return s.dao.GetSpecifyGroupRecord(mid)
	default:
		return nil, model.ErrRecordNotFind
	}
}

// GetPriRecord 所有聊天记录
func (s *Service) GetRecordsAfterMid(tp xproto.Channel, fromId, targetId string, mid int64, recordCount int64) ([]*model.MsgContent, error) {
	switch tp {
	case xproto.Channel_ToUser:
		return s.dao.GetPriRecord(fromId, targetId, mid, recordCount)
	case xproto.Channel_ToGroup:
		return nil, model.ErrCustomNotSupport
	default:
		return nil, model.ErrRecordNotFind
	}
}

// AddRecordFocus 添加消息关注记录
func (s *Service) AddRecordFocus(uid string, mid int64, time uint64) error {
	return s.dao.AddRecordFocus(uid, mid, time)
}

// StatisticRecordFocusNumber 统计消息关注数
func (s *Service) StatisticRecordFocusNumber(mid int64) (int32, error) {
	return s.dao.GetRecordFocusNumber(mid)
}
