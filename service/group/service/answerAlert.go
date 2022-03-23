package service

import (
	"context"

	"github.com/txchat/dtalk/pkg/util"

	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/txchat/dtalk/service/group/model"
	answer "github.com/txchat/dtalk/service/record/answer/api"
	xproto "github.com/txchat/imparse/proto"
)

func (s *Service) PushNoticeMsg(ctx context.Context, channelType int32, from, target string, data []byte) error {
	noticeMsgReq := &answer.PushNoticeMsgReq{
		Seq:         uuid.New().String(),
		ChannelType: channelType,
		From:        from,
		Target:      target,
		Data:        data,
	}

	if _, err := s.answerClient.PushNoticeMsg(
		ctx, noticeMsgReq.Seq, noticeMsgReq.ChannelType,
		noticeMsgReq.From, noticeMsgReq.Target, noticeMsgReq.Data); err != nil {
		return err
	}
	return nil
}

func (s *Service) PushGroupNoticeMsg(ctx context.Context, groupId int64, personId string, noticeMsgMsg *xproto.NoticeMsg) error {
	data, err := proto.Marshal(noticeMsgMsg)
	if err != nil {
		return err
	}
	from := personId
	target := util.ToString(groupId)

	if err = s.PushNoticeMsg(ctx, 1, from, target, data); err != nil {
		return err
	}
	return nil
}

func convertNoticeMsgMsg(noticeMsgType xproto.NoticeMsgType, noticeMsg proto.Message) (*xproto.NoticeMsg, error) {
	body, err := proto.Marshal(noticeMsg)
	if err != nil {
		return nil, err
	}
	noticeMsgMsg := &xproto.NoticeMsg{
		Type: noticeMsgType,
		Body: body,
	}
	return noticeMsgMsg, nil
}

func (s *Service) NoticeMsgUpdateGroupName(ctx context.Context, groupId int64, personId string, name string) error {
	if name == "" {
		return nil
	}

	noticeMsgType := xproto.NoticeMsgType_UpdateGroupNameNoticeMsg
	noticeMsg := &xproto.NoticeMsgUpdateGroupName{
		Group:    groupId,
		Operator: personId,
		Name:     name,
	}
	noticeMsgMsg, err := convertNoticeMsgMsg(noticeMsgType, noticeMsg)
	if err != nil {
		return err
	}

	if err = s.PushGroupNoticeMsg(ctx, groupId, personId, noticeMsgMsg); err != nil {
		return err
	}
	return nil
}

func (s *Service) NoticeMsgSignInGroup(ctx context.Context, groupId int64, personId string, memberIds []string) error {
	noticeMsgType := xproto.NoticeMsgType_SignInGroupNoticeMsg
	noticeMsg := &xproto.NoticeMsgSignInGroup{
		Group:   groupId,
		Inviter: personId,
		Members: memberIds,
	}
	noticeMsgMsg, err := convertNoticeMsgMsg(noticeMsgType, noticeMsg)
	if err != nil {
		return err
	}

	if err = s.PushGroupNoticeMsg(ctx, groupId, personId, noticeMsgMsg); err != nil {
		return err
	}
	return nil
}

func (s *Service) NoticeMsgSignOutGroup(ctx context.Context, groupId int64, personId string) error {
	noticeMsgType := xproto.NoticeMsgType_SignOutGroupNoticeMsg
	noticeMsg := &xproto.NoticeMsgSignOutGroup{
		Group:    groupId,
		Operator: personId,
	}
	body, err := proto.Marshal(noticeMsg)
	if err != nil {
		return err
	}

	noticeMsgMsg := &xproto.NoticeMsg{
		Type: noticeMsgType,
		Body: body,
	}

	if err = s.PushGroupNoticeMsg(ctx, groupId, personId, noticeMsgMsg); err != nil {
		return err
	}
	return nil
}

func (s *Service) NoticeMsgKickOutGroup(ctx context.Context, groupId int64, personId string, memberIds []string) error {
	noticeMsgType := xproto.NoticeMsgType_KickOutGroupNoticeMsg
	noticeMsg := &xproto.NoticeMsgKickOutGroup{
		Group:    groupId,
		Operator: personId,
		Members:  memberIds,
	}
	body, err := proto.Marshal(noticeMsg)
	if err != nil {
		return err
	}

	noticeMsgMsg := &xproto.NoticeMsg{
		Type: noticeMsgType,
		Body: body,
	}

	if err = s.PushGroupNoticeMsg(ctx, groupId, personId, noticeMsgMsg); err != nil {
		return err
	}
	return nil
}

func (s *Service) NoticeMsgDeleteGroup(ctx context.Context, groupId int64, personId string) error {
	noticeMsgType := xproto.NoticeMsgType_DeleteGroupNoticeMsg
	noticeMsg := &xproto.NoticeMsgDeleteGroup{
		Group:    groupId,
		Operator: personId,
	}
	body, err := proto.Marshal(noticeMsg)
	if err != nil {
		return err
	}

	noticeMsgMsg := &xproto.NoticeMsg{
		Type: noticeMsgType,
		Body: body,
	}

	if err = s.PushGroupNoticeMsg(ctx, groupId, personId, noticeMsgMsg); err != nil {
		return err
	}
	return nil
}

func (s *Service) NoticeMsgUpdateGroupMuted(ctx context.Context, groupId int64, personId string, muteType int32) error {
	var proMuteType xproto.MuteType
	switch muteType {
	case int32(xproto.MuteType_MuteAllow):
		proMuteType = xproto.MuteType_MuteAllow
	case int32(xproto.MuteType_MuteDeny):
		proMuteType = xproto.MuteType_MuteDeny
	default:
		return model.ErrType
	}
	noticeMsgType := xproto.NoticeMsgType_UpdateGroupMutedNoticeMsg
	noticeMsg := &xproto.NoticeMsgUpdateGroupMuted{
		Group:    groupId,
		Operator: personId,
		Type:     proMuteType,
	}
	body, err := proto.Marshal(noticeMsg)
	if err != nil {
		return err
	}

	noticeMsgMsg := &xproto.NoticeMsg{
		Type: noticeMsgType,
		Body: body,
	}

	if err = s.PushGroupNoticeMsg(ctx, groupId, personId, noticeMsgMsg); err != nil {
		return err
	}
	return nil
}

func (s *Service) NoticeMsgUpdateGroupMemberMutedTime(ctx context.Context, groupId int64, personId string, memberIds []string) error {
	noticeMsgType := xproto.NoticeMsgType_UpdateGroupMemberMutedNoticeMsg
	noticeMsg := &xproto.NoticeMsgUpdateGroupMemberMutedTime{
		Group:    groupId,
		Operator: personId,
		Members:  memberIds,
	}
	body, err := proto.Marshal(noticeMsg)
	if err != nil {
		return err
	}

	noticeMsgMsg := &xproto.NoticeMsg{
		Type: noticeMsgType,
		Body: body,
	}

	if err = s.PushGroupNoticeMsg(ctx, groupId, personId, noticeMsgMsg); err != nil {
		return err
	}
	return nil
}

func (s *Service) NoticeMsgUpdateGroupOwner(ctx context.Context, groupId int64, personId, newOwner string) error {
	noticeMsgType := xproto.NoticeMsgType_UpdateGroupOwnerNoticeMsg
	noticeMsg := &xproto.NoticeMsgUpdateGroupOwner{
		Group:    groupId,
		NewOwner: newOwner,
	}
	body, err := proto.Marshal(noticeMsg)
	if err != nil {
		return err
	}

	noticeMsgMsg := &xproto.NoticeMsg{
		Type: noticeMsgType,
		Body: body,
	}

	if err = s.PushGroupNoticeMsg(ctx, groupId, personId, noticeMsgMsg); err != nil {
		return err
	}
	return nil
}
