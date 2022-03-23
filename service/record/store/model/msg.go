package model

import (
	"encoding/base64"
	"encoding/json"

	"github.com/golang/protobuf/proto"
	zlog "github.com/rs/zerolog/log"
	xproto "github.com/txchat/imparse/proto"
)

var msgFactory = map[xproto.MsgType]func() proto.Message{
	xproto.MsgType_Text: func() proto.Message {
		return &xproto.TextMsg{}
	},
	xproto.MsgType_Audio: func() proto.Message {
		return &xproto.AudioMsg{}
	},
	xproto.MsgType_Image: func() proto.Message {
		return &xproto.ImageMsg{}
	},
	xproto.MsgType_Video: func() proto.Message {
		return &xproto.VideoMsg{}
	},
	xproto.MsgType_File: func() proto.Message {
		return &xproto.FileMsg{}
	},
	xproto.MsgType_Card: func() proto.Message {
		return &xproto.CardMsg{}
	},
	xproto.MsgType_Notice: func() proto.Message {
		return &xproto.NoticeMsg{}
	},
	xproto.MsgType_Forward: func() proto.Message {
		return &xproto.ForwardMsg{}
	},
	xproto.MsgType_Transfer: func() proto.Message {
		return &xproto.TransferMsg{}
	},
	xproto.MsgType_RedPacket: func() proto.Message {
		return &xproto.RedPacketMsg{}
	},
	xproto.MsgType_ContactCard: func() proto.Message {
		return &xproto.ContactCardMsg{}
	},
}

var signalFactory = map[xproto.SignalType]func() proto.Message{
	xproto.SignalType_Received: func() proto.Message {
		return &xproto.SignalReceived{}
	},
	xproto.SignalType_Revoke: func() proto.Message {
		return &xproto.SignalRevoke{}
	},
	xproto.SignalType_SignInGroup: func() proto.Message {
		return &xproto.SignalSignInGroup{}
	},
	xproto.SignalType_SignOutGroup: func() proto.Message {
		return &xproto.SignalSignOutGroup{}
	},
	xproto.SignalType_DeleteGroup: func() proto.Message {
		return &xproto.SignalDeleteGroup{}
	},
	xproto.SignalType_FocusMessage: func() proto.Message {
		return &xproto.SignalFocusMessage{}
	},
	xproto.SignalType_UpdateGroupJoinType: func() proto.Message {
		return &xproto.SignalUpdateGroupJoinType{}
	},
	xproto.SignalType_UpdateGroupFriendType: func() proto.Message {
		return &xproto.SignalUpdateGroupFriendType{}
	},
	xproto.SignalType_UpdateGroupMuteType: func() proto.Message {
		return &xproto.SignalUpdateGroupMuteType{}
	},
	xproto.SignalType_UpdateGroupMemberType: func() proto.Message {
		return &xproto.SignalUpdateGroupMemberType{}
	},
	xproto.SignalType_UpdateGroupMemberMuteTime: func() proto.Message {
		return &xproto.SignalUpdateGroupMemberMuteTime{}
	},
	xproto.SignalType_UpdateGroupName: func() proto.Message {
		return &xproto.SignalUpdateGroupName{}
	},
	xproto.SignalType_UpdateGroupAvatar: func() proto.Message {
		return &xproto.SignalUpdateGroupAvatar{}
	},
	xproto.SignalType_StartCall: func() proto.Message {
		return &xproto.SignalStartCall{}
	},
	xproto.SignalType_AcceptCall: func() proto.Message {
		return &xproto.SignalAcceptCall{}
	},
	xproto.SignalType_StopCall: func() proto.Message {
		return &xproto.SignalStopCall{}
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
		m = &xproto.EncryptMsg{
			Content: string(data),
		}
		return
	}
	return
}

func ParseCommonMsg(m *xproto.Common) []byte {
	creator, ok := msgFactory[m.MsgType]
	if !ok || creator == nil {
		zlog.Debug().Int32("type", int32(m.MsgType)).Interface("creator", creator).Msg("ParseCommon unknown type")
		return m.Msg
	}
	msg := creator()
	return parse(msg, m.Msg)
}

func ParseSource(m *xproto.Common) []byte {
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

func ParseReference(m *xproto.Common) []byte {
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

func ParseSignal(m *xproto.Signal) []byte {
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
	creator, ok := msgFactory[xproto.MsgType(msgType)]
	if !ok || creator == nil {
		zlog.Error().Uint32("type", msgType).Interface("creator", creator).Msg("ConvertMsg unknown type")
		return nil
	}
	msg := creator()
	jsonConvert(msg, data)
	return msg
}

func ConvertMsg(msgType uint32, data []byte) []byte {
	creator, ok := msgFactory[xproto.MsgType(msgType)]
	if !ok || creator == nil {
		zlog.Error().Uint32("type", msgType).Interface("creator", creator).Msg("ConvertMsg unknown type")
		return data
	}
	msg := creator()
	return convert(msg, data)
}

func ConvertSource(data []byte) *xproto.Source {
	if len(data) == 0 {
		return nil
	}
	var src xproto.Source
	err := json.Unmarshal(data, &src)
	if err != nil {
		zlog.Error().Err(err).Msg("ParseMsg json.Unmarshal err")
		return &src
	}
	return &src
}

func ConvertReference(data []byte) *xproto.Reference {
	if len(data) == 0 {
		return nil
	}
	var src xproto.Reference
	err := json.Unmarshal(data, &src)
	if err != nil {
		zlog.Error().Err(err).Msg("ParseMsg json.Unmarshal err")
		return &src
	}
	return &src
}

func ConvertSignal(actionType uint32, data []byte) []byte {
	creator, ok := signalFactory[xproto.SignalType(actionType)]
	if !ok || creator == nil {
		zlog.Error().Uint32("action", actionType).Interface("creator", creator).Msg("ConvertMsg unknown type")
		return data
	}
	msg := creator()
	return convert(msg, data)
}
