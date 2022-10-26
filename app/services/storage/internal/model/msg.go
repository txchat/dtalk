package model

import (
	"encoding/base64"
	"encoding/json"

	"github.com/golang/protobuf/proto"
	zlog "github.com/rs/zerolog/log"
	"github.com/txchat/imparse/proto/common"
	"github.com/txchat/imparse/proto/msg"
	"github.com/txchat/imparse/proto/signal"
)

var msgFactory = map[common.MsgType]func() proto.Message{
	common.MsgType_Text: func() proto.Message {
		return &msg.TextMsg{}
	},
	common.MsgType_Audio: func() proto.Message {
		return &msg.AudioMsg{}
	},
	common.MsgType_Image: func() proto.Message {
		return &msg.ImageMsg{}
	},
	common.MsgType_Video: func() proto.Message {
		return &msg.VideoMsg{}
	},
	common.MsgType_File: func() proto.Message {
		return &msg.FileMsg{}
	},
	common.MsgType_Card: func() proto.Message {
		return &msg.CardMsg{}
	},
	common.MsgType_Notice: func() proto.Message {
		return &msg.NoticeMsg{}
	},
	common.MsgType_Forward: func() proto.Message {
		return &msg.ForwardMsg{}
	},
	common.MsgType_Transfer: func() proto.Message {
		return &msg.TransferMsg{}
	},
	common.MsgType_RedPacket: func() proto.Message {
		return &msg.RedPacketMsg{}
	},
	common.MsgType_ContactCard: func() proto.Message {
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

func parse(m proto.Message, data []byte) []byte {
	err := proto.Unmarshal(data, m)
	if err != nil {
		zlog.Debug().Err(err).Msg("ParseCommonMsg proto.Unmarshal err")
		return []byte(base64.StdEncoding.EncodeToString(data))
	}
	b, err := json.Marshal(m)
	if err != nil {
		zlog.Error().Err(err).Msg("arseMsg json.Marshal err")
		return data
	}
	return b
}

func convert(m proto.Message, data []byte) []byte {
	err := json.Unmarshal(data, m)
	if err != nil {
		zlog.Debug().Err(err).Msg("ConvertMsg json.Unmarshal")
		v, err := base64.StdEncoding.DecodeString(string(data))
		if err != nil {
			zlog.Error().Err(err).Msg("ConvertMsg base64.StdEncoding.DecodeString")
			return data
		}
		return v
	}
	b, err := proto.Marshal(m)
	if err != nil {
		zlog.Error().Err(err).Msg("ConvertMsg proto.Marshal")
		return data
	}
	return b
}

func jsonConvert(m proto.Message, data []byte) {
	err := json.Unmarshal(data, m)
	if err != nil {
		zlog.Error().Err(err).Msg("ConvertMsg json.Unmarshal")
		//v, err := base64.StdEncoding.DecodeString(string(data))
		//if err != nil {
		//	log15.Error("ConvertMsg base64.StdEncoding.DecodeString", "err", err)
		//	return data
		//}
		m = &msg.EncryptMsg{
			Content: string(data),
		}
		return
	}
	return
}

func ParseCommonMsg(m *common.Common) []byte {
	creator, ok := msgFactory[m.MsgType]
	if !ok || creator == nil {
		zlog.Debug().Int32("type", int32(m.MsgType)).Interface("creator", creator).Msg("ParseCommon unknown type")
		return m.Msg
	}
	msg := creator()
	return parse(msg, m.Msg)
}

func ParseSource(m *common.Common) []byte {
	if m.Source == nil {
		return nil
	}
	b, err := json.Marshal(m.Source)
	if err != nil {
		zlog.Error().Err(err).Msg("ParseCommon json.Marshal err")
		return b
	}
	return b
}

func ParseReference(m *common.Common) []byte {
	if m.Reference == nil {
		return nil
	}
	b, err := json.Marshal(m.Reference)
	if err != nil {
		zlog.Error().Err(err).Msg("ParseCommon json.Marshal err")
		return b
	}
	return b
}

func ParseSignal(m *signal.Signal) []byte {
	creator, ok := signalFactory[m.Type]
	if !ok || creator == nil {
		zlog.Error().Interface("action", m.Type).Interface("creator", creator).Msg("ParseSignal unknown type")
		return m.Body
	}
	msg := creator()
	return parse(msg, m.Body)
}

//
func JsonUnmarshal(msgType uint32, data []byte) proto.Message {
	creator, ok := msgFactory[common.MsgType(msgType)]
	if !ok || creator == nil {
		zlog.Error().Uint32("type", msgType).Interface("creator", creator).Msg("ConvertMsg unknown type")
		return nil
	}
	msg := creator()
	jsonConvert(msg, data)
	return msg
}

func ConvertMsg(msgType uint32, data []byte) []byte {
	creator, ok := msgFactory[common.MsgType(msgType)]
	if !ok || creator == nil {
		zlog.Error().Uint32("type", msgType).Interface("creator", creator).Msg("ConvertMsg unknown type")
		return data
	}
	msg := creator()
	return convert(msg, data)
}

func ConvertSource(data []byte) *common.Source {
	if len(data) == 0 {
		return nil
	}
	var src common.Source
	err := json.Unmarshal(data, &src)
	if err != nil {
		zlog.Error().Err(err).Msg("ParseMsg json.Unmarshal err")
		return &src
	}
	return &src
}

func ConvertReference(data []byte) *common.Reference {
	if len(data) == 0 {
		return nil
	}
	var src common.Reference
	err := json.Unmarshal(data, &src)
	if err != nil {
		zlog.Error().Err(err).Msg("ParseMsg json.Unmarshal err")
		return &src
	}
	return &src
}

func ConvertSignal(actionType uint32, data []byte) []byte {
	creator, ok := signalFactory[signal.SignalType(actionType)]
	if !ok || creator == nil {
		zlog.Error().Uint32("action", actionType).Interface("creator", creator).Msg("ConvertMsg unknown type")
		return data
	}
	msg := creator()
	return convert(msg, data)
}
