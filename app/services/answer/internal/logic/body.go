package logic

import (
	"github.com/golang/protobuf/proto"
	comet "github.com/txchat/im/api/comet/grpc"
	xproto "github.com/txchat/imparse/proto"
)

var signalReliableMap = map[xproto.SignalType]bool{
	xproto.SignalType_Received:      true,
	xproto.SignalType_Revoke:        true,
	xproto.SignalType_SignInGroup:   true,
	xproto.SignalType_SignOutGroup:  true,
	xproto.SignalType_DeleteGroup:   true,
	xproto.SignalType_FocusMessage:  true,
	xproto.SignalType_EndpointLogin: false,
	//
	xproto.SignalType_UpdateGroupJoinType:       true,
	xproto.SignalType_UpdateGroupFriendType:     true,
	xproto.SignalType_UpdateGroupMuteType:       true,
	xproto.SignalType_UpdateGroupMemberType:     true,
	xproto.SignalType_UpdateGroupMemberMuteTime: true,
	xproto.SignalType_UpdateGroupName:           true,
	xproto.SignalType_UpdateGroupAvatar:         true,
	//
	xproto.SignalType_StartCall:  true,
	xproto.SignalType_AcceptCall: true,
	xproto.SignalType_StopCall:   true,
}

func noticeMsgData(channelType int32, from, target string, seq string, data []byte) ([]byte, error) {
	var p comet.Proto
	var err error
	p.Op = int32(comet.Op_SendMsg)
	p.Ver = 1
	p.Seq = 0

	eventProto := &xproto.Proto{
		EventType: xproto.Proto_common,
	}
	comm := &xproto.Common{
		ChannelType: xproto.Channel(channelType),
		Seq:         seq,
		From:        from,
		Target:      target,
		MsgType:     xproto.MsgType_Notice,
		Msg:         data,
	}
	eventProto.Body, err = proto.Marshal(comm)
	if err != nil {
		return nil, err
	}
	p.Body, err = proto.Marshal(eventProto)
	if err != nil {
		return nil, err
	}
	bytes, err := proto.Marshal(&p)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func signalBody(target string, tp xproto.SignalType, actionData []byte) ([]byte, error) {
	var p comet.Proto
	var err error
	p.Op = int32(comet.Op_SendMsg)
	p.Ver = 1
	p.Seq = 0

	eventProto := &xproto.Proto{
		EventType: xproto.Proto_Signal,
	}
	noticeProto := &xproto.Signal{
		Type:     tp,
		Reliable: signalReliableMap[tp],
	}
	noticeProto.Body = actionData
	eventProto.Body, err = proto.Marshal(noticeProto)
	if err != nil {
		return nil, err
	}
	p.Body, err = proto.Marshal(eventProto)
	if err != nil {
		return nil, err
	}
	bytes, err := proto.Marshal(&p)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
