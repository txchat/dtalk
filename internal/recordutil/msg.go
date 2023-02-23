package recordutil

import (
	"encoding/base64"
	"encoding/json"

	"github.com/golang/protobuf/proto"
	"github.com/txchat/dtalk/api/proto/message"
	"github.com/txchat/dtalk/api/proto/msg"
	"github.com/txchat/dtalk/api/proto/signal"
)

var msgFactory = map[message.MsgType]func() proto.Message{
	message.MsgType_Text: func() proto.Message {
		return &msg.TextMsg{}
	},
	message.MsgType_Audio: func() proto.Message {
		return &msg.AudioMsg{}
	},
	message.MsgType_Image: func() proto.Message {
		return &msg.ImageMsg{}
	},
	message.MsgType_Video: func() proto.Message {
		return &msg.VideoMsg{}
	},
	message.MsgType_File: func() proto.Message {
		return &msg.FileMsg{}
	},
	message.MsgType_Card: func() proto.Message {
		return &msg.CardMsg{}
	},
	message.MsgType_Notice: func() proto.Message {
		return &msg.NoticeMsg{}
	},
	message.MsgType_Forward: func() proto.Message {
		return &msg.ForwardMsg{}
	},
	message.MsgType_Transfer: func() proto.Message {
		return &msg.TransferMsg{}
	},
	message.MsgType_RedPacket: func() proto.Message {
		return &msg.RedPacketMsg{}
	},
	message.MsgType_ContactCard: func() proto.Message {
		return &msg.ContactCardMsg{}
	},
}

var signalFactory = map[signal.SignalType]func() proto.Message{
	signal.SignalType_Received: func() proto.Message {
		return &signal.SignalReceived{}
	},
	signal.SignalType_Revoke: func() proto.Message {
		return &signal.SignalRevoke{}
	},
	signal.SignalType_SignInGroup: func() proto.Message {
		return &signal.SignalSignInGroup{}
	},
	signal.SignalType_SignOutGroup: func() proto.Message {
		return &signal.SignalSignOutGroup{}
	},
	signal.SignalType_DeleteGroup: func() proto.Message {
		return &signal.SignalDeleteGroup{}
	},
	signal.SignalType_FocusMessage: func() proto.Message {
		return &signal.SignalFocusMessage{}
	},
	signal.SignalType_UpdateGroupJoinType: func() proto.Message {
		return &signal.SignalUpdateGroupJoinType{}
	},
	signal.SignalType_UpdateGroupFriendType: func() proto.Message {
		return &signal.SignalUpdateGroupFriendType{}
	},
	signal.SignalType_UpdateGroupMuteType: func() proto.Message {
		return &signal.SignalUpdateGroupMuteType{}
	},
	signal.SignalType_UpdateGroupMemberType: func() proto.Message {
		return &signal.SignalUpdateGroupMemberType{}
	},
	signal.SignalType_UpdateGroupMemberMuteTime: func() proto.Message {
		return &signal.SignalUpdateGroupMemberMuteTime{}
	},
	signal.SignalType_UpdateGroupName: func() proto.Message {
		return &signal.SignalUpdateGroupName{}
	},
	signal.SignalType_UpdateGroupAvatar: func() proto.Message {
		return &signal.SignalUpdateGroupAvatar{}
	},
	signal.SignalType_StartCall: func() proto.Message {
		return &signal.SignalStartCall{}
	},
	signal.SignalType_AcceptCall: func() proto.Message {
		return &signal.SignalAcceptCall{}
	},
	signal.SignalType_StopCall: func() proto.Message {
		return &signal.SignalStopCall{}
	},
}

func IsMsgSupport(msgType message.MsgType) bool {
	creator, ok := msgFactory[msgType]
	return ok && creator != nil
}

func protobufDataToJSONData(m proto.Message, data []byte) []byte {
	err := proto.Unmarshal(data, m)
	if err != nil {
		return []byte(base64.StdEncoding.EncodeToString(data))
	}
	b, err := json.Marshal(m)
	if err != nil {
		return data
	}
	return b
}

func jsonDataToProtobufData(m proto.Message, data []byte) []byte {
	err := json.Unmarshal(data, m)
	if err != nil {
		v, err := base64.StdEncoding.DecodeString(string(data))
		if err != nil {
			return data
		}
		return v
	}
	b, err := proto.Marshal(m)
	if err != nil {
		return data
	}
	return b
}

func jsonDataToProtobufMessage(m proto.Message, data []byte) {
	err := json.Unmarshal(data, m)
	if err != nil {
		m = &msg.EncryptMsg{
			Content: string(data),
		}
		return
	}
}

func CommonMsgProtobufDataToJSONData(m *message.Message) []byte {
	creator, ok := msgFactory[m.MsgType]
	if !ok || creator == nil {
		return m.Msg
	}
	protoMsg := creator()
	return protobufDataToJSONData(protoMsg, m.Msg)
}

func SourceJSONMarshal(m *message.Message) []byte {
	if m.Source == nil {
		return nil
	}
	b, err := json.Marshal(m.Source)
	if err != nil {
		return b
	}
	return b
}

func ReferenceJSONMarshal(m *message.Message) []byte {
	if m.Reference == nil {
		return nil
	}
	b, err := json.Marshal(m.Reference)
	if err != nil {
		return b
	}
	return b
}

func SignalContentToJSONData(m *signal.Signal) []byte {
	creator, ok := signalFactory[m.Type]
	if !ok || creator == nil {
		return m.Body
	}
	protoMsg := creator()
	return protobufDataToJSONData(protoMsg, m.Body)
}

func CommonMsgJSONDataToProtobuf(msgType int32, data []byte) proto.Message {
	creator, ok := msgFactory[message.MsgType(msgType)]
	if !ok || creator == nil {
		return nil
	}
	protoMsg := creator()
	jsonDataToProtobufMessage(protoMsg, data)
	return protoMsg
}

func CommonMsgJSONDataToProtobufData(msgType int32, data []byte) []byte {
	creator, ok := msgFactory[message.MsgType(msgType)]
	if !ok || creator == nil {
		return data
	}
	protoMsg := creator()
	return jsonDataToProtobufData(protoMsg, data)
}

func SourceJSONUnmarshal(data []byte) *message.Source {
	if len(data) == 0 {
		return nil
	}
	var src message.Source
	err := json.Unmarshal(data, &src)
	if err != nil {
		return &src
	}
	return &src
}

func ReferenceJSONUnmarshal(data []byte) *message.Reference {
	if len(data) == 0 {
		return nil
	}
	var src message.Reference
	err := json.Unmarshal(data, &src)
	if err != nil {
		return &src
	}
	return &src
}

func SignalContentJSONDataToProtobufData(actionType int32, data []byte) []byte {
	creator, ok := signalFactory[signal.SignalType(actionType)]
	if !ok || creator == nil {
		return data
	}
	protoMsg := creator()
	return jsonDataToProtobufData(protoMsg, data)
}
