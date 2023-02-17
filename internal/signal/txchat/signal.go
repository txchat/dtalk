package txchat

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/txchat/dtalk/api/proto/chat"
	"github.com/txchat/dtalk/api/proto/message"
	"github.com/txchat/dtalk/api/proto/signal"
	"github.com/txchat/dtalk/app/services/pusher/pusherclient"
	"github.com/txchat/dtalk/app/services/transfer/transferclient"
	"github.com/txchat/dtalk/internal/group"
	"github.com/txchat/dtalk/internal/recordhelper"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/im/api/protocol"
)

type SignalHub struct {
	transferClient transferclient.Transfer
	pusherClient   pusherclient.Pusher
}

func NewSignalHub(transferClient transferclient.Transfer, pusherClient pusherclient.Pusher) *SignalHub {
	return &SignalHub{
		transferClient: transferClient,
		pusherClient:   pusherClient,
	}
}

func (hub *SignalHub) convertSignalProtoData(signalType signal.SignalType, actionBody []byte) (*chat.Chat, error) {
	sig := &signal.Signal{
		Type: signalType,
		Body: actionBody,
	}
	sigData, err := proto.Marshal(sig)
	if err != nil {
		return nil, err
	}

	signalEvent := &chat.Chat{
		Type: chat.Chat_signal,
		Seq:  0,
		Body: sigData,
	}
	return signalEvent, nil
}

func (hub *SignalHub) GroupAddNewMembers(ctx context.Context, gid int64, members []string) error {
	action := &signal.SignalSignInGroup{
		Uid:   members,
		Group: gid,
		Time:  util.MustToUint64(util.TimeNowUnixMilli()),
	}
	body, err := proto.Marshal(action)
	if err != nil {
		return err
	}
	chatProto, err := hub.convertSignalProtoData(signal.SignalType_SignInGroup, body)
	if err != nil {
		return err
	}

	_, err = hub.transferClient.TransferMessage(ctx, &transferclient.TransferMessageReq{
		From:        "",
		Target:      util.MustToString(gid),
		ChannelType: message.Channel_Group,
		Body:        chatProto,
	})
	return err
}

func (hub *SignalHub) GroupRemoveMembers(ctx context.Context, gid int64, members []string) error {
	action := &signal.SignalSignOutGroup{
		Uid:   members,
		Group: gid,
		Time:  util.MustToUint64(util.TimeNowUnixMilli()),
	}
	body, err := proto.Marshal(action)
	if err != nil {
		return err
	}

	chatProto, err := hub.convertSignalProtoData(signal.SignalType_SignOutGroup, body)
	if err != nil {
		return err
	}

	_, err = hub.transferClient.TransferMessage(ctx, &transferclient.TransferMessageReq{
		From:        "",
		Target:      util.MustToString(gid),
		ChannelType: message.Channel_Group,
		Body:        chatProto,
	})
	return err
}

func (hub *SignalHub) GroupDeleted(ctx context.Context, gid int64) error {
	action := &signal.SignalDeleteGroup{
		Group: gid,
		Time:  util.MustToUint64(util.TimeNowUnixMilli()),
	}
	body, err := proto.Marshal(action)
	if err != nil {
		return err
	}

	chatProto, err := hub.convertSignalProtoData(signal.SignalType_DeleteGroup, body)
	if err != nil {
		return err
	}

	_, err = hub.transferClient.TransferMessage(ctx, &transferclient.TransferMessageReq{
		From:        "",
		Target:      util.MustToString(gid),
		ChannelType: message.Channel_Group,
		Body:        chatProto,
	})
	return err
}

func (hub *SignalHub) UpdateGroupJoinType(ctx context.Context, gid int64, tp int32) error {
	action := &signal.SignalUpdateGroupJoinType{
		Group: gid,
		Type:  signal.JoinType(tp),
		Time:  util.MustToUint64(util.TimeNowUnixMilli()),
	}
	body, err := proto.Marshal(action)
	if err != nil {
		return err
	}

	chatProto, err := hub.convertSignalProtoData(signal.SignalType_UpdateGroupJoinType, body)
	if err != nil {
		return err
	}

	_, err = hub.transferClient.TransferMessage(ctx, &transferclient.TransferMessageReq{
		From:        "",
		Target:      util.MustToString(gid),
		ChannelType: message.Channel_Group,
		Body:        chatProto,
	})
	return err
}

func (hub *SignalHub) UpdateGroupFriendlyType(ctx context.Context, gid int64, tp int32) error {
	action := &signal.SignalUpdateGroupFriendType{
		Group: gid,
		Type:  signal.FriendType(tp),
		Time:  util.MustToUint64(util.TimeNowUnixMilli()),
	}
	body, err := proto.Marshal(action)
	if err != nil {
		return err
	}

	chatProto, err := hub.convertSignalProtoData(signal.SignalType_UpdateGroupFriendType, body)
	if err != nil {
		return err
	}

	_, err = hub.transferClient.TransferMessage(ctx, &transferclient.TransferMessageReq{
		From:        "",
		Target:      util.MustToString(gid),
		ChannelType: message.Channel_Group,
		Body:        chatProto,
	})
	return err
}

func (hub *SignalHub) UpdateGroupMuteType(ctx context.Context, gid int64, tp int32) error {
	action := &signal.SignalUpdateGroupMuteType{
		Group: gid,
		Type:  signal.MuteType(tp),
		Time:  util.MustToUint64(util.TimeNowUnixMilli()),
	}
	body, err := proto.Marshal(action)
	if err != nil {
		return err
	}

	chatProto, err := hub.convertSignalProtoData(signal.SignalType_UpdateGroupMuteType, body)
	if err != nil {
		return err
	}

	_, err = hub.transferClient.TransferMessage(ctx, &transferclient.TransferMessageReq{
		From:        "",
		Target:      util.MustToString(gid),
		ChannelType: message.Channel_Group,
		Body:        chatProto,
	})
	return err
}

func (hub *SignalHub) UpdateGroupName(ctx context.Context, gid int64, name string) error {
	action := &signal.SignalUpdateGroupName{
		Group: gid,
		Name:  name,
		Time:  util.MustToUint64(util.TimeNowUnixMilli()),
	}
	body, err := proto.Marshal(action)
	if err != nil {
		return err
	}

	chatProto, err := hub.convertSignalProtoData(signal.SignalType_UpdateGroupName, body)
	if err != nil {
		return err
	}

	_, err = hub.transferClient.TransferMessage(ctx, &transferclient.TransferMessageReq{
		From:        "",
		Target:      util.MustToString(gid),
		ChannelType: message.Channel_Group,
		Body:        chatProto,
	})
	return err
}

func (hub *SignalHub) UpdateGroupAvatar(ctx context.Context, gid int64, avatar string) error {
	action := &signal.SignalUpdateGroupAvatar{
		Group:  gid,
		Avatar: avatar,
		Time:   util.MustToUint64(util.TimeNowUnixMilli()),
	}
	body, err := proto.Marshal(action)
	if err != nil {
		return err
	}

	chatProto, err := hub.convertSignalProtoData(signal.SignalType_UpdateGroupAvatar, body)
	if err != nil {
		return err
	}

	_, err = hub.transferClient.TransferMessage(ctx, &transferclient.TransferMessageReq{
		From:        "",
		Target:      util.MustToString(gid),
		ChannelType: message.Channel_Group,
		Body:        chatProto,
	})
	return err
}

func (hub *SignalHub) UpdateGroupMemberRole(ctx context.Context, gid int64, uid string, role group.RoleType) error {
	action := &signal.SignalUpdateGroupMemberType{
		Group: gid,
		Uid:   uid,
		Type:  signal.MemberType(role),
		Time:  util.MustToUint64(util.TimeNowUnixMilli()),
	}
	body, err := proto.Marshal(action)
	if err != nil {
		return err
	}

	chatProto, err := hub.convertSignalProtoData(signal.SignalType_UpdateGroupMemberType, body)
	if err != nil {
		return err
	}

	_, err = hub.transferClient.TransferMessage(ctx, &transferclient.TransferMessageReq{
		From:        "",
		Target:      util.MustToString(gid),
		ChannelType: message.Channel_Group,
		Body:        chatProto,
	})
	return err
}

func (hub *SignalHub) UpdateMembersMuteTime(ctx context.Context, gid, muteTime int64, members []string) error {
	action := &signal.SignalUpdateGroupMemberMuteTime{
		Group:    gid,
		Uid:      members,
		MuteTime: muteTime,
		Time:     util.MustToUint64(util.TimeNowUnixMilli()),
	}
	body, err := proto.Marshal(action)
	if err != nil {
		return err
	}

	chatProto, err := hub.convertSignalProtoData(signal.SignalType_UpdateGroupMemberMuteTime, body)
	if err != nil {
		return err
	}

	_, err = hub.transferClient.TransferMessage(ctx, &transferclient.TransferMessageReq{
		From:        "",
		Target:      util.MustToString(gid),
		ChannelType: message.Channel_Group,
		Body:        chatProto,
	})
	return err
}

func (hub *SignalHub) StartCall(ctx context.Context, target string, traceId int64) error {
	action := &signal.SignalStartCall{
		TraceId: traceId,
	}
	body, err := proto.Marshal(action)
	if err != nil {
		return errors.WithMessagef(err, "proto.Marshal, action=%+v", action)
	}

	chatProto, err := hub.convertSignalProtoData(signal.SignalType_StartCall, body)
	if err != nil {
		return err
	}

	_, err = hub.transferClient.TransferMessage(ctx, &transferclient.TransferMessageReq{
		From:        "",
		Target:      target,
		ChannelType: message.Channel_Private,
		Body:        chatProto,
	})
	return err
}

func (hub *SignalHub) AcceptCall(ctx context.Context, target string, action *signal.SignalAcceptCall) error {
	body, err := proto.Marshal(action)
	if err != nil {
		return errors.WithMessagef(err, "proto.Marshal, action=%+v", action)
	}

	chatProto, err := hub.convertSignalProtoData(signal.SignalType_AcceptCall, body)
	if err != nil {
		return err
	}

	_, err = hub.transferClient.TransferMessage(ctx, &transferclient.TransferMessageReq{
		From:        "",
		Target:      target,
		ChannelType: message.Channel_Private,
		Body:        chatProto,
	})
	return err
}

func (hub *SignalHub) StopCall(ctx context.Context, target string, action *signal.SignalStopCall) error {
	body, err := proto.Marshal(action)
	if err != nil {
		return errors.WithMessagef(err, "proto.Marshal, action=%+v", action)
	}

	chatProto, err := hub.convertSignalProtoData(signal.SignalType_StopCall, body)
	if err != nil {
		return err
	}

	_, err = hub.transferClient.TransferMessage(ctx, &transferclient.TransferMessageReq{
		From:        "",
		Target:      target,
		ChannelType: message.Channel_Private,
		Body:        chatProto,
	})
	return err
}

func (hub *SignalHub) MessageReceived(ctx context.Context, item *recordhelper.ConnSeqItem) error {
	action := &signal.SignalReceived{
		Logs: item.Logs,
	}
	body, err := proto.Marshal(action)
	if err != nil {
		return err
	}

	chatProto, err := hub.convertSignalProtoData(signal.SignalType_Received, body)
	if err != nil {
		return err
	}

	_, err = hub.transferClient.TransferMessage(ctx, &transferclient.TransferMessageReq{
		From:        "",
		Target:      item.Sender,
		ChannelType: message.Channel_Private,
		Body:        chatProto,
	})
	return err
}

func (hub *SignalHub) EndpointLogin(ctx context.Context, uid string, action *signal.SignalEndpointLogin) error {
	body, err := proto.Marshal(action)
	if err != nil {
		return err
	}

	chatProto, err := hub.convertSignalProtoData(signal.SignalType_EndpointLogin, body)
	if err != nil {
		return err
	}

	data, err := proto.Marshal(chatProto)
	if err != nil {
		return err
	}
	p := &protocol.Proto{
		Ver:  1,
		Op:   int32(protocol.Op_ReceiveMsg),
		Seq:  0,
		Ack:  0,
		Body: data,
	}
	data, err = proto.Marshal(p)
	if err != nil {
		return err
	}

	_, err = hub.pusherClient.PushList(ctx, &pusherclient.PushListReq{
		App:  "",
		From: "",
		Uid:  []string{uid},
		Body: data,
	})
	return err
}

func (hub *SignalHub) FocusPrivateMessage(ctx context.Context, users []string, action *signal.SignalFocusMessage) error {
	body, err := proto.Marshal(action)
	if err != nil {
		return err
	}

	for _, user := range users {

		chatProto, err := hub.convertSignalProtoData(signal.SignalType_FocusMessage, body)
		if err != nil {
			return err
		}

		if _, err = hub.transferClient.TransferMessage(ctx, &transferclient.TransferMessageReq{
			From:        "",
			Target:      user,
			ChannelType: message.Channel_Private,
			Body:        chatProto,
		}); err != nil {
			return err
		}
	}
	return nil
}

func (hub *SignalHub) FocusGroupMessage(ctx context.Context, gid int64, action *signal.SignalFocusMessage) error {
	body, err := proto.Marshal(action)
	if err != nil {
		return err
	}

	chatProto, err := hub.convertSignalProtoData(signal.SignalType_FocusMessage, body)
	if err != nil {
		return err
	}

	_, err = hub.transferClient.TransferMessage(ctx, &transferclient.TransferMessageReq{
		From:        "",
		Target:      util.MustToString(gid),
		ChannelType: message.Channel_Group,
		Body:        chatProto,
	})
	return err
}

func (hub *SignalHub) RevokePrivateMessage(ctx context.Context, users []string, action *signal.SignalRevoke) error {
	body, err := proto.Marshal(action)
	if err != nil {
		return errors.WithMessagef(err, "proto.Marshal, action=%+v", action)
	}

	for _, user := range users {

		chatProto, err := hub.convertSignalProtoData(signal.SignalType_Revoke, body)
		if err != nil {
			return err
		}

		if _, err = hub.transferClient.TransferMessage(ctx, &transferclient.TransferMessageReq{
			From:        "",
			Target:      user,
			ChannelType: message.Channel_Private,
			Body:        chatProto,
		}); err != nil {
			return err
		}
	}
	return nil
}

func (hub *SignalHub) RevokeGroupMessage(ctx context.Context, gid int64, action *signal.SignalRevoke) error {
	body, err := proto.Marshal(action)
	if err != nil {
		return errors.WithMessagef(err, "proto.Marshal, action=%+v", action)
	}

	chatProto, err := hub.convertSignalProtoData(signal.SignalType_Revoke, body)
	if err != nil {
		return err
	}

	_, err = hub.transferClient.TransferMessage(ctx, &transferclient.TransferMessageReq{
		From:        "",
		Target:      util.MustToString(gid),
		ChannelType: message.Channel_Group,
		Body:        chatProto,
	})
	return err
}
