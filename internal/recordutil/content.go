package recordutil

import (
	"github.com/golang/protobuf/proto"
	"github.com/txchat/dtalk/api/proto/content"
	"github.com/txchat/dtalk/api/proto/message"
)

func MessageAlterContent(m *message.Message) string {
	creator, ok := msgFactory[m.MsgType]
	if !ok || creator == nil {
		return ""
	}
	protoMsg := creator()
	switch m.MsgType {
	case message.MsgType_Text:
		text := protoMsg.(*content.TextMsg)
		err := proto.Unmarshal(m.GetContent(), text)
		if err != nil {
			return ""
		}
		return text.GetContent()
	case message.MsgType_Audio:
		return "[收到一条语言]"
	case message.MsgType_Image:
		return "[收到一张图片]"
	case message.MsgType_Video:
		return "[收到一条视频]"
	case message.MsgType_File:
		return "[收到一个文件]"
	case message.MsgType_RedPacket:
		return "[收到一个红包]"
	}
	return ""
}
