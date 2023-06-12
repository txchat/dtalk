package recordutil

import (
	"encoding/json"

	"github.com/golang/protobuf/proto"
	"github.com/txchat/dtalk/api/proto/content"
	"github.com/txchat/dtalk/api/proto/message"
	"github.com/txchat/dtalk/api/proto/signal"
)

var msgFactory = map[message.MsgType]func() proto.Message{
	message.MsgType_Text: func() proto.Message {
		return &content.TextMsg{}
	},
	message.MsgType_Audio: func() proto.Message {
		return &content.AudioMsg{}
	},
	message.MsgType_Image: func() proto.Message {
		return &content.ImageMsg{}
	},
	message.MsgType_Video: func() proto.Message {
		return &content.VideoMsg{}
	},
	message.MsgType_File: func() proto.Message {
		return &content.FileMsg{}
	},
	message.MsgType_Card: func() proto.Message {
		return &content.CardMsg{}
	},
	message.MsgType_Notice: func() proto.Message {
		return &content.NoticeMsg{}
	},
	message.MsgType_Forward: func() proto.Message {
		return &content.ForwardMsg{}
	},
	message.MsgType_Transfer: func() proto.Message {
		return &content.TransferMsg{}
	},
	message.MsgType_RedPacket: func() proto.Message {
		return &content.RedPacketMsg{}
	},
	message.MsgType_ContactCard: func() proto.Message {
		return &content.ContactCardMsg{}
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

func protobufDataToJSONData(m proto.Message, data []byte) ([]byte, error) {
	err := proto.Unmarshal(data, m)
	if err != nil {
		return nil, err
	}
	return json.Marshal(m)
}

//
//func jsonDataToProtobufData(m proto.Message, data []byte) ([]byte, error) {
//	err := json.Unmarshal(data, m)
//	if err != nil {
//		return nil, err
//	}
//	return proto.Marshal(m)
//}

func MessageContentProtobufDataToJSONData(m *message.Message) []byte {
	creator, ok := msgFactory[m.MsgType]
	if !ok || creator == nil {
		return m.Content
	}
	protoMsg := creator()
	data, err := protobufDataToJSONData(protoMsg, m.Content)
	if err != nil {
		return m.Content
	}
	return data
}

func MessageContentJSONDataToProtobuf(msgType message.MsgType, data []byte) proto.Message {
	creator, ok := msgFactory[msgType]
	if !ok || creator == nil {
		return nil
	}
	protoMsg := creator()
	err := json.Unmarshal(data, protoMsg)
	if err != nil {
		return nil
	}
	return protoMsg
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

func SignalContentProtobufDataToJSONData(m *signal.Signal) []byte {
	creator, ok := signalFactory[m.Type]
	if !ok || creator == nil {
		return m.Body
	}
	protoMsg := creator()
	data, err := protobufDataToJSONData(protoMsg, m.Body)
	if err != nil {
		return m.Body
	}
	return data
}
