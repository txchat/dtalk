package msgfactory

import (
	"github.com/txchat/dtalk/app/services/answer/internal/model"
	"github.com/txchat/imparse"
	"github.com/txchat/imparse/chat"
	"github.com/txchat/imparse/proto/common"
)

func isChTypeOk(t common.Channel) bool {
	switch t {
	case common.Channel_ToUser:
		return true
	case common.Channel_ToGroup:
		return true
	}
	return false
}

func isMsgTypeOk(t common.MsgType) bool {
	switch t {
	case common.MsgType_System:
		return false
	case common.MsgType_Text:
		return true
	case common.MsgType_Audio:
		return true
	case common.MsgType_Image:
		return true
	case common.MsgType_Video:
		return true
	case common.MsgType_File:
		return true
	case common.MsgType_Card:
		return true
	case common.MsgType_Notice:
		return false
	case common.MsgType_Forward:
		return true
	case common.MsgType_Transfer:
		return true
	case common.MsgType_Collect:
		return false
	case common.MsgType_RedPacket:
		return true
	case common.MsgType_ContactCard:
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
