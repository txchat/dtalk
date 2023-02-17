package txchat

import (
	"context"

	"github.com/txchat/dtalk/api/proto/message"
	"github.com/txchat/dtalk/app/services/transfer/transferclient"

	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/txchat/dtalk/api/proto/common"
	"github.com/txchat/dtalk/api/proto/msg"
	"github.com/txchat/dtalk/api/proto/signal"
	"github.com/txchat/dtalk/app/services/answer/answerclient"
	"github.com/txchat/dtalk/internal/group"
	"github.com/txchat/dtalk/pkg/util"
)

type NoticeHub struct {
	transferClient transferclient.Transfer
}

func NewNoticeHub(transferClient transferclient.Transfer) *NoticeHub {
	return &NoticeHub{
		transferClient: transferClient,
	}
}

func (hub *NoticeHub) GroupAddNewMembers(ctx context.Context, gid int64, operator string, members []string) error {
	noticeMetadata := &msg.NoticeMsgSignInGroup{
		Group:   gid,
		Inviter: operator,
		Members: members,
	}
	body, err := proto.Marshal(noticeMetadata)
	if err != nil {
		return err
	}

	noticeMsg := &msg.NoticeMsg{
		Type: msg.NoticeMsgType_SignInGroupNoticeMsg,
		Body: body,
	}
	data, err := proto.Marshal(noticeMsg)
	if err != nil {
		return err
	}

	msg := &message.Message{
		ChannelType: 0,
		Mid:         "",
		Cid:         "",
		From:        "",
		Target:      "",
		MsgType:     0,
		Msg:         nil,
		Datetime:    0,
		Source:      nil,
		Reference:   nil,
	}

	_, err = hub.answerClient.PushNoticeMsg(ctx, &answerclient.PushNoticeMsgReq{
		Seq:         uuid.New().String(),
		ChannelType: int32(common.Channel_ToGroup),
		From:        operator,
		Target:      util.MustToString(gid),
		Data:        data,
	})
	return err
}

func (hub *NoticeHub) GroupSignOut(ctx context.Context, gid int64, target string) error {
	noticeMetadata := &msg.NoticeMsgSignOutGroup{
		Group:    gid,
		Operator: target,
	}
	body, err := proto.Marshal(noticeMetadata)
	if err != nil {
		return err
	}

	noticeMsg := &msg.NoticeMsg{
		Type: msg.NoticeMsgType_SignOutGroupNoticeMsg,
		Body: body,
	}
	data, err := proto.Marshal(noticeMsg)
	if err != nil {
		return err
	}

	_, err = hub.answerClient.PushNoticeMsg(ctx, &answerclient.PushNoticeMsgReq{
		Seq:         uuid.New().String(),
		ChannelType: int32(common.Channel_ToGroup),
		From:        target,
		Target:      util.MustToString(gid),
		Data:        data,
	})
	return err
}

func (hub *NoticeHub) GroupKickOutMembers(ctx context.Context, gid int64, operator string, members []string) error {
	noticeMetadata := &msg.NoticeMsgKickOutGroup{
		Group:    gid,
		Operator: operator,
		Members:  members,
	}
	body, err := proto.Marshal(noticeMetadata)
	if err != nil {
		return err
	}

	noticeMsg := &msg.NoticeMsg{
		Type: msg.NoticeMsgType_KickOutGroupNoticeMsg,
		Body: body,
	}
	data, err := proto.Marshal(noticeMsg)
	if err != nil {
		return err
	}

	_, err = hub.answerClient.PushNoticeMsg(ctx, &answerclient.PushNoticeMsgReq{
		Seq:         uuid.New().String(),
		ChannelType: int32(common.Channel_ToGroup),
		From:        operator,
		Target:      util.MustToString(gid),
		Data:        data,
	})
	return err
}

func (hub *NoticeHub) GroupDeleted(ctx context.Context, gid int64, operator string) error {
	noticeMetadata := &msg.NoticeMsgDeleteGroup{
		Group:    gid,
		Operator: operator,
	}
	body, err := proto.Marshal(noticeMetadata)
	if err != nil {
		return err
	}

	noticeMsg := &msg.NoticeMsg{
		Type: msg.NoticeMsgType_DeleteGroupNoticeMsg,
		Body: body,
	}
	data, err := proto.Marshal(noticeMsg)
	if err != nil {
		return err
	}

	_, err = hub.answerClient.PushNoticeMsg(ctx, &answerclient.PushNoticeMsgReq{
		Seq:         uuid.New().String(),
		ChannelType: int32(common.Channel_ToGroup),
		From:        operator,
		Target:      util.MustToString(gid),
		Data:        data,
	})
	return err
}

func (hub *NoticeHub) UpdateGroupMuteType(ctx context.Context, gid int64, operator string, tp int32) error {
	noticeMetadata := &msg.NoticeMsgUpdateGroupMuted{
		Group:    gid,
		Operator: operator,
		Type:     signal.MuteType(tp),
	}
	body, err := proto.Marshal(noticeMetadata)
	if err != nil {
		return err
	}

	noticeMsg := &msg.NoticeMsg{
		Type: msg.NoticeMsgType_UpdateGroupMutedNoticeMsg,
		Body: body,
	}
	data, err := proto.Marshal(noticeMsg)
	if err != nil {
		return err
	}

	_, err = hub.answerClient.PushNoticeMsg(ctx, &answerclient.PushNoticeMsgReq{
		Seq:         uuid.New().String(),
		ChannelType: int32(common.Channel_ToGroup),
		From:        operator,
		Target:      util.MustToString(gid),
		Data:        data,
	})
	return err
}

func (hub *NoticeHub) UpdateGroupName(ctx context.Context, gid int64, operator, name string) error {
	noticeMetadata := &msg.NoticeMsgUpdateGroupName{
		Group:    gid,
		Operator: operator,
		Name:     name,
	}
	body, err := proto.Marshal(noticeMetadata)
	if err != nil {
		return err
	}

	noticeMsg := &msg.NoticeMsg{
		Type: msg.NoticeMsgType_UpdateGroupNameNoticeMsg,
		Body: body,
	}
	data, err := proto.Marshal(noticeMsg)
	if err != nil {
		return err
	}

	_, err = hub.answerClient.PushNoticeMsg(ctx, &answerclient.PushNoticeMsgReq{
		Seq:         uuid.New().String(),
		ChannelType: int32(common.Channel_ToGroup),
		From:        operator,
		Target:      util.MustToString(gid),
		Data:        data,
	})
	return err
}

func (hub *NoticeHub) UpdateGroupMemberRole(ctx context.Context, gid int64, operator, uid string, role group.RoleType) error {
	switch role {
	case group.RoleOwner:
	default:
		return nil
	}
	noticeMetadata := &msg.NoticeMsgUpdateGroupOwner{
		Group:    gid,
		NewOwner: uid,
	}
	body, err := proto.Marshal(noticeMetadata)
	if err != nil {
		return err
	}

	noticeMsg := &msg.NoticeMsg{
		Type: msg.NoticeMsgType_UpdateGroupOwnerNoticeMsg,
		Body: body,
	}
	data, err := proto.Marshal(noticeMsg)
	if err != nil {
		return err
	}

	_, err = hub.answerClient.PushNoticeMsg(ctx, &answerclient.PushNoticeMsgReq{
		Seq:         uuid.New().String(),
		ChannelType: int32(common.Channel_ToGroup),
		From:        operator,
		Target:      util.MustToString(gid),
		Data:        data,
	})
	return err
}

func (hub *NoticeHub) UpdateMembersMuteTime(ctx context.Context, operator string, gid, muteTime int64, members []string) error {
	noticeMetadata := &msg.NoticeMsgUpdateGroupMemberMutedTime{
		Group:    gid,
		Operator: operator,
		Members:  members,
	}
	body, err := proto.Marshal(noticeMetadata)
	if err != nil {
		return err
	}

	noticeMsg := &msg.NoticeMsg{
		Type: msg.NoticeMsgType_UpdateGroupMemberMutedNoticeMsg,
		Body: body,
	}
	data, err := proto.Marshal(noticeMsg)
	if err != nil {
		return err
	}

	_, err = hub.answerClient.PushNoticeMsg(ctx, &answerclient.PushNoticeMsgReq{
		Seq:         uuid.New().String(),
		ChannelType: int32(common.Channel_ToGroup),
		From:        operator,
		Target:      util.MustToString(gid),
		Data:        data,
	})
	return err
}
