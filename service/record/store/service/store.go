package service

import (
	"fmt"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/dtalk/service/record/store/model"
	"github.com/txchat/imparse"
	"github.com/txchat/imparse/chat"
	xproto "github.com/txchat/imparse/proto"
	"strconv"
	"time"
)

func (s *Service) MarkReceived(tp, uid string, mid int64) error {
	switch tp {
	case string(chat.PrivateFrameType):
		_, _, err := s.dao.MarkMsgReceived(uid, mid)
		return err
	case string(chat.GroupFrameType):
		_, _, err := s.dao.MarkGroupMsgReceived(uid, mid)
		return err
	case string(chat.SignalFrameType):
		_, _, err := s.dao.MarkSignalReceived(uid, mid)
		return err
	default:
		return nil
	}
}

type DB struct {
	s *Service
}

func (d *DB) SaveMsg(deviceType xproto.Device, frame imparse.Frame) error {
	switch frame.Type() {
	case chat.PrivateFrameType:
		pv := frame.(*chat.PrivateFrame)
		if pv.GetTransmissionMethod() != imparse.UniCast {
			return model.ErrFrameType
		}
		return d.s.StoreMsg(deviceType, pv.GetBase())
	case chat.GroupFrameType:
		pv := frame.(*chat.GroupFrame)
		if pv.GetTransmissionMethod() != imparse.GroupCast {
			return model.ErrFrameType
		}
		return d.s.StoreGroupMsg(deviceType, pv.GetBase())
	case chat.SignalFrameType:
		pv := frame.(*chat.SignalFrame)
		base := pv.GetBase()
		if !base.GetReliable() {
			return nil
		}
		data := model.ParseSignal(base)
		createTime := uint64(util.TimeNowUnixNano() / int64(time.Millisecond))
		switch pv.GetTransmissionMethod() {
		case imparse.UniCast:
			return d.s.AppendUniCastSignal(strconv.FormatInt(pv.GetMid(), 10), pv.GetTarget(), base.GetType(), data, createTime)
		case imparse.GroupCast:
			return d.s.AppendGroupCastSignal(strconv.FormatInt(pv.GetMid(), 10), pv.GetTarget(), base.GetType(), data, createTime)
		default:
			return fmt.Errorf("%v : frame type: %v, tarnsmission method : %v", model.ErrFrameType, frame.Type(), pv.GetTransmissionMethod())
		}
	}
	return fmt.Errorf("%v : frame type: %v, tarnsmission method : %v", model.ErrFrameType, frame.Type())
}
