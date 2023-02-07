package logic

import (
	"github.com/golang/protobuf/proto"
	"github.com/txchat/im/api/protocol"
	"github.com/txchat/imparse/proto/common"
	"github.com/txchat/imparse/proto/signal"
)

var signalReliableMap = map[signal.SignalType]bool{
	signal.SignalType_Received:      true,
	signal.SignalType_Revoke:        true,
	signal.SignalType_SignInGroup:   true,
	signal.SignalType_SignOutGroup:  true,
	signal.SignalType_DeleteGroup:   true,
	signal.SignalType_FocusMessage:  true,
	signal.SignalType_EndpointLogin: false,
	//
	signal.SignalType_UpdateGroupJoinType:       true,
	signal.SignalType_UpdateGroupFriendType:     true,
	signal.SignalType_UpdateGroupMuteType:       true,
	signal.SignalType_UpdateGroupMemberType:     true,
	signal.SignalType_UpdateGroupMemberMuteTime: true,
	signal.SignalType_UpdateGroupName:           true,
	signal.SignalType_UpdateGroupAvatar:         true,
	//
	signal.SignalType_StartCall:  true,
	signal.SignalType_AcceptCall: true,
	signal.SignalType_StopCall:   true,
}

func noticeMsgData(channelType int32, from, target string, seq string, data []byte) ([]byte, error) {
	var p protocol.Proto
	var err error
	p.Op = int32(protocol.Op_SendMsg)
	p.Ver = 1
	p.Seq = 0

	eventProto := &common.Proto{
		EventType: common.Proto_common,
	}
	comm := &common.Common{
		ChannelType: common.Channel(channelType),
		Seq:         seq,
		From:        from,
		Target:      target,
		MsgType:     common.MsgType_Notice,
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

func signalBody(target string, tp signal.SignalType, actionData []byte) ([]byte, error) {
	var p protocol.Proto
	var err error
	p.Op = int32(protocol.Op_SendMsg)
	p.Ver = 1
	p.Seq = 0

	eventProto := &common.Proto{
		EventType: common.Proto_Signal,
	}
	noticeProto := &signal.Signal{
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
