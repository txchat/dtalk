package msgfactory

import (
	"github.com/txchat/dtalk/app/services/answer/internal/model"
	"github.com/txchat/imparse"
	"github.com/txchat/imparse/chat"
	xproto "github.com/txchat/imparse/proto"
)

func isChTypeOk(t xproto.Channel) bool {
	switch t {
	case xproto.Channel_ToUser:
		return true
	case xproto.Channel_ToGroup:
		return true
	}
	return false
}

func isMsgTypeOk(t xproto.MsgType) bool {
	switch t {
	case xproto.MsgType_System:
		return false
	case xproto.MsgType_Text:
		return true
	case xproto.MsgType_Audio:
		return true
	case xproto.MsgType_Image:
		return true
	case xproto.MsgType_Video:
		return true
	case xproto.MsgType_File:
		return true
	case xproto.MsgType_Card:
		return true
	case xproto.MsgType_Notice:
		return false
	case xproto.MsgType_Forward:
		return true
	case xproto.MsgType_Transfer:
		return true
	case xproto.MsgType_Collect:
		return false
	case xproto.MsgType_RedPacket:
		return true
	case xproto.MsgType_ContactCard:
		return true
	}
	return false
}

type Checker struct {
}

func NewChecker() *Checker {
	return &Checker{}
}

func (c *Checker) CheckFrame(frame imparse.Frame) error {
	switch frame.Type() {
	case chat.PrivateFrameType:
		f := frame.(*chat.PrivateFrame)
		if !isChTypeOk(f.GetChannelType()) {
			return model.ErrorChType
		}
		if !isMsgTypeOk(f.GetMsgType()) {
			return model.ErrorMsgType
		}
	case chat.GroupFrameType:
		f := frame.(*chat.GroupFrame)
		if !isChTypeOk(f.GetChannelType()) {
			return model.ErrorChType
		}
		if !isMsgTypeOk(f.GetMsgType()) {
			return model.ErrorMsgType
		}
	case chat.SignalFrameType:
		return model.ErrorEnvType
	default:
		return model.ErrorEnvType
	}
	return nil
}
