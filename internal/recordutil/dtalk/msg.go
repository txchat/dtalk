package checker

import (
	"github.com/txchat/dtalk/api/proto/chat"
	"github.com/txchat/dtalk/api/proto/message"
	"github.com/txchat/dtalk/internal/recordutil"
)

func checkChannelType(channel message.Channel) bool {
	switch channel {
	case message.Channel_Private, message.Channel_Group:
		return true
	}
	return false
}

func checkCid(cid string) bool {
	return cid != ""
}

func checkTarget(target string) bool {
	return target != ""
}

func checkMsgType(msgType message.MsgType) bool {
	return recordutil.IsMsgSupport(msgType)
}

func CheckMessage(msg *message.Message) chat.SendMessageReply_FailedType {
	if !checkCid(msg.GetCid()) || !checkTarget(msg.GetTarget()) {
		return chat.SendMessageReply_IllegalFormat
	}
	if !checkChannelType(msg.GetChannelType()) || !checkMsgType(msg.GetMsgType()) {
		return chat.SendMessageReply_UnSupportedMessageType
	}
	return chat.SendMessageReply_IsOK
}
