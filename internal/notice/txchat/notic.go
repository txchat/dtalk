package txchat

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/txchat/dtalk/api/proto/chat"
	"github.com/txchat/dtalk/api/proto/content"
	"github.com/txchat/dtalk/api/proto/message"
	"github.com/txchat/dtalk/api/proto/signal"
	"github.com/txchat/dtalk/app/services/transfer/transferclient"
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

func (hub *NoticeHub) convertNoticeProtoData(channelType message.Channel, from, target string, noticeType content.NoticeMsgType, noticeMetadata proto.Message) (*chat.Chat, error) {
	noticeData, err := proto.Marshal(noticeMetadata)
	if err != nil {
		return nil, err
	}
	msgData, err := proto.Marshal(&content.NoticeMsg{
		Type: noticeType,
		Body: noticeData,
	})
	if err != nil {
		return nil, err
	}
	body, err := proto.Marshal(&message.Message{
		ChannelType: channelType,
		Cid:         uuid.New().String(),
		From:        from,
		Target:      target,
		MsgType:     message.MsgType_Notice,
		Content:     msgData,
	})
	if err != nil {
		return nil, err
	}
	chatProto := &chat.Chat{
		Type: chat.Chat_message,
		Seq:  0,
		Body: body,
	}
	return chatProto, nil
}

func (hub *NoticeHub) GroupAddNewMembers(ctx context.Context, gid int64, operator string, members []string) error {
	noticeMetadata := &content.NoticeMsgSignInGroup{
		Group:   gid,
		Inviter: operator,
		Members: members,
	}
	chatProto, err := hub.convertNoticeProtoData(message.Channel_Group, operator, util.MustToString(gid), content.NoticeMsgType_SignInGroupNoticeMsg, noticeMetadata)
	if err != nil {
		return err
	}
	_, err = hub.transferClient.TransferMessage(ctx, &transferclient.TransferMessageReq{
		From:        operator,
		Target:      util.MustToString(gid),
		ChannelType: message.Channel_Group,
		Body:        chatProto,
	})
	return err
}

func (hub *NoticeHub) GroupSignOut(ctx context.Context, gid int64, target string) error {
	noticeMetadata := &content.NoticeMsgSignOutGroup{
		Group:    gid,
		Operator: target,
	}
	chatProto, err := hub.convertNoticeProtoData(message.Channel_Group, target, util.MustToString(gid), content.NoticeMsgType_SignOutGroupNoticeMsg, noticeMetadata)
	if err != nil {
		return err
	}
	_, err = hub.transferClient.TransferMessage(ctx, &transferclient.TransferMessageReq{
		From:        target,
		Target:      util.MustToString(gid),
		ChannelType: message.Channel_Group,
		Body:        chatProto,
	})
	return err
}

func (hub *NoticeHub) GroupKickOutMembers(ctx context.Context, gid int64, operator string, members []string) error {
	noticeMetadata := &content.NoticeMsgKickOutGroup{
		Group:    gid,
		Operator: operator,
		Members:  members,
	}
	chatProto, err := hub.convertNoticeProtoData(message.Channel_Group, operator, util.MustToString(gid), content.NoticeMsgType_KickOutGroupNoticeMsg, noticeMetadata)
	if err != nil {
		return err
	}
	_, err = hub.transferClient.TransferMessage(ctx, &transferclient.TransferMessageReq{
		From:        operator,
		Target:      util.MustToString(gid),
		ChannelType: message.Channel_Group,
		Body:        chatProto,
	})
	return err
}

func (hub *NoticeHub) GroupDeleted(ctx context.Context, gid int64, operator string) error {
	noticeMetadata := &content.NoticeMsgDeleteGroup{
		Group:    gid,
		Operator: operator,
	}
	chatProto, err := hub.convertNoticeProtoData(message.Channel_Group, operator, util.MustToString(gid), content.NoticeMsgType_DeleteGroupNoticeMsg, noticeMetadata)
	if err != nil {
		return err
	}
	_, err = hub.transferClient.TransferMessage(ctx, &transferclient.TransferMessageReq{
		From:        operator,
		Target:      util.MustToString(gid),
		ChannelType: message.Channel_Group,
		Body:        chatProto,
	})
	return err
}

func (hub *NoticeHub) UpdateGroupMuteType(ctx context.Context, gid int64, operator string, tp int32) error {
	noticeMetadata := &content.NoticeMsgUpdateGroupMuted{
		Group:    gid,
		Operator: operator,
		Type:     signal.MuteType(tp),
	}
	chatProto, err := hub.convertNoticeProtoData(message.Channel_Group, operator, util.MustToString(gid), content.NoticeMsgType_UpdateGroupMutedNoticeMsg, noticeMetadata)
	if err != nil {
		return err
	}
	_, err = hub.transferClient.TransferMessage(ctx, &transferclient.TransferMessageReq{
		From:        operator,
		Target:      util.MustToString(gid),
		ChannelType: message.Channel_Group,
		Body:        chatProto,
	})
	return err
}

func (hub *NoticeHub) UpdateGroupName(ctx context.Context, gid int64, operator, name string) error {
	noticeMetadata := &content.NoticeMsgUpdateGroupName{
		Group:    gid,
		Operator: operator,
		Name:     name,
	}
	chatProto, err := hub.convertNoticeProtoData(message.Channel_Group, operator, util.MustToString(gid), content.NoticeMsgType_UpdateGroupNameNoticeMsg, noticeMetadata)
	if err != nil {
		return err
	}
	_, err = hub.transferClient.TransferMessage(ctx, &transferclient.TransferMessageReq{
		From:        operator,
		Target:      util.MustToString(gid),
		ChannelType: message.Channel_Group,
		Body:        chatProto,
	})
	return err
}

func (hub *NoticeHub) UpdateGroupMemberRole(ctx context.Context, gid int64, operator, uid string, role group.RoleType) error {
	switch role {
	case group.RoleOwner:
	default:
		return nil
	}
	noticeMetadata := &content.NoticeMsgUpdateGroupOwner{
		Group:    gid,
		NewOwner: uid,
	}
	chatProto, err := hub.convertNoticeProtoData(message.Channel_Group, operator, util.MustToString(gid), content.NoticeMsgType_UpdateGroupOwnerNoticeMsg, noticeMetadata)
	if err != nil {
		return err
	}
	_, err = hub.transferClient.TransferMessage(ctx, &transferclient.TransferMessageReq{
		From:        operator,
		Target:      util.MustToString(gid),
		ChannelType: message.Channel_Group,
		Body:        chatProto,
	})
	return err
}

func (hub *NoticeHub) UpdateMembersMuteTime(ctx context.Context, operator string, gid, muteTime int64, members []string) error {
	noticeMetadata := &content.NoticeMsgUpdateGroupMemberMutedTime{
		Group:    gid,
		Operator: operator,
		Members:  members,
	}
	chatProto, err := hub.convertNoticeProtoData(message.Channel_Group, operator, util.MustToString(gid), content.NoticeMsgType_UpdateGroupMemberMutedNoticeMsg, noticeMetadata)
	if err != nil {
		return err
	}
	_, err = hub.transferClient.TransferMessage(ctx, &transferclient.TransferMessageReq{
		From:        operator,
		Target:      util.MustToString(gid),
		ChannelType: message.Channel_Group,
		Body:        chatProto,
	})
	return err
}
