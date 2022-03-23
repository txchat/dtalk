package service

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/dtalk/service/group/model"
	xproto "github.com/txchat/imparse/proto"
)

func (s *Service) PusherSignalJoin(ctx context.Context, groupId int64, groupMemberIds []string) error {
	nowTime := s.getNowTime()
	actionSignInGroup := &xproto.SignalSignInGroup{
		Uid:   groupMemberIds,
		Group: groupId,
		Time:  util.ToUInt64(nowTime),
	}
	body, err := proto.Marshal(actionSignInGroup)
	if err != nil {
		return err
	}

	Signal := xproto.SignalType_SignInGroup
	Target := util.ToString(groupId)
	Body := body

	if err = s.answerClient.GroupCastSignal(ctx, Signal, Target, Body); err != nil {
		return err
	}
	return nil
}

func (s *Service) PusherSignalLeave(ctx context.Context, groupId int64, groupMemberIds []string) error {
	nowTime := s.getNowTime()
	actionSignOutGroup := &xproto.SignalSignOutGroup{
		Uid:   groupMemberIds,
		Group: groupId,
		Time:  util.ToUInt64(nowTime),
	}
	body, err := proto.Marshal(actionSignOutGroup)
	if err != nil {
		return err
	}

	Signal := xproto.SignalType_SignOutGroup
	Target := util.ToString(groupId)
	Body := body

	if err = s.answerClient.GroupCastSignal(ctx, Signal, Target, Body); err != nil {
		return err
	}
	return nil
}

func (s *Service) PusherSignalDel(ctx context.Context, groupId int64) error {
	nowTime := s.getNowTime()
	actionDealeteGroup := &xproto.SignalDeleteGroup{
		Group: groupId,
		Time:  util.ToUInt64(nowTime),
	}
	body, err := proto.Marshal(actionDealeteGroup)
	if err != nil {
		return err
	}

	Signal := xproto.SignalType_DeleteGroup
	Target := util.ToString(groupId)
	Body := body

	if err = s.answerClient.GroupCastSignal(ctx, Signal, Target, Body); err != nil {
		return err
	}
	return nil
}

func (s *Service) PusherSignalJoinType(ctx context.Context, groupId int64, joinType int32) error {
	nowTime := s.getNowTime()

	var proJoinType xproto.JoinType

	switch joinType {
	case int32(xproto.JoinType_JoinAllow):
		proJoinType = xproto.JoinType_JoinAllow
	case int32(xproto.JoinType_JoinDeny):
		proJoinType = xproto.JoinType_JoinDeny
	case int32(xproto.JoinType_JoinApply):
		proJoinType = xproto.JoinType_JoinApply

	default:
		return model.ErrType
	}

	action := &xproto.SignalUpdateGroupJoinType{
		Group: groupId,
		Type:  proJoinType,
		Time:  util.ToUInt64(nowTime),
	}
	body, err := proto.Marshal(action)
	if err != nil {
		return err
	}

	Signal := xproto.SignalType_UpdateGroupJoinType
	Target := util.ToString(groupId)
	Body := body

	if err = s.answerClient.GroupCastSignal(ctx, Signal, Target, Body); err != nil {
		return err
	}
	return nil
}

func (s *Service) PusherSignalFriendType(ctx context.Context, groupId int64, friendType int32) error {
	nowTime := s.getNowTime()

	var proFriendType xproto.FriendType

	switch friendType {
	case int32(xproto.FriendType_FriendAllow):
		proFriendType = xproto.FriendType_FriendAllow
	case int32(xproto.FriendType_FriendDeny):
		proFriendType = xproto.FriendType_FriendDeny
	default:
		return model.ErrType
	}

	action := &xproto.SignalUpdateGroupFriendType{
		Group: groupId,
		Type:  proFriendType,
		Time:  util.ToUInt64(nowTime),
	}
	body, err := proto.Marshal(action)
	if err != nil {
		return err
	}

	Signal := xproto.SignalType_UpdateGroupFriendType
	Target := util.ToString(groupId)
	Body := body

	if err = s.answerClient.GroupCastSignal(ctx, Signal, Target, Body); err != nil {
		return err
	}
	return nil
}

func (s *Service) PusherSignalMuteType(ctx context.Context, groupId int64, muteType int32) error {
	nowTime := s.getNowTime()

	var proMuteType xproto.MuteType

	switch muteType {
	case int32(xproto.MuteType_MuteAllow):
		proMuteType = xproto.MuteType_MuteAllow
	case int32(xproto.MuteType_MuteDeny):
		proMuteType = xproto.MuteType_MuteDeny
	default:
		return model.ErrType
	}

	action := &xproto.SignalUpdateGroupMuteType{
		Group: groupId,
		Type:  proMuteType,
		Time:  util.ToUInt64(nowTime),
	}
	body, err := proto.Marshal(action)
	if err != nil {
		return err
	}

	Signal := xproto.SignalType_UpdateGroupMuteType
	Target := util.ToString(groupId)
	Body := body

	if err = s.answerClient.GroupCastSignal(ctx, Signal, Target, Body); err != nil {
		return err
	}
	return nil
}

func (s *Service) PusherSignalMemberType(ctx context.Context, groupId int64, memberId string, memberType int32) error {
	nowTime := s.getNowTime()

	var proMemberType xproto.MemberType

	switch memberType {
	case int32(xproto.MemberType_Normal):
		proMemberType = xproto.MemberType_Normal
	case int32(xproto.MemberType_Admin):
		proMemberType = xproto.MemberType_Admin
	case int32(xproto.MemberType_Owner):
		proMemberType = xproto.MemberType_Owner
	default:
		return model.ErrType
	}

	action := &xproto.SignalUpdateGroupMemberType{
		Group: groupId,
		Type:  proMemberType,
		Uid:   memberId,
		Time:  util.ToUInt64(nowTime),
	}
	body, err := proto.Marshal(action)
	if err != nil {
		return err
	}

	Signal := xproto.SignalType_UpdateGroupMemberType
	Target := util.ToString(groupId)
	Body := body

	if err = s.answerClient.GroupCastSignal(ctx, Signal, Target, Body); err != nil {
		return err
	}
	return nil
}

func (s *Service) PusherSignalMemberMuteTime(ctx context.Context, groupId int64, memberIds []string, muteTime int64) error {
	nowTime := s.getNowTime()
	action := &xproto.SignalUpdateGroupMemberMuteTime{
		Group:    groupId,
		Uid:      memberIds,
		MuteTime: muteTime,
		Time:     util.ToUInt64(nowTime),
	}
	body, err := proto.Marshal(action)
	if err != nil {
		return err
	}

	Signal := xproto.SignalType_UpdateGroupMemberMuteTime
	Target := util.ToString(groupId)
	Body := body

	if err = s.answerClient.GroupCastSignal(ctx, Signal, Target, Body); err != nil {
		return err
	}
	return nil
}

func (s *Service) PusherSignalGroupName(ctx context.Context, groupId int64, name string) error {
	if name == "" {
		return nil
	}

	nowTime := s.getNowTime()
	action := &xproto.SignalUpdateGroupName{
		Group: groupId,
		Name:  name,
		Time:  util.ToUInt64(nowTime),
	}
	body, err := proto.Marshal(action)
	if err != nil {
		return err
	}

	Signal := xproto.SignalType_UpdateGroupName
	Target := util.ToString(groupId)
	Body := body

	if err = s.answerClient.GroupCastSignal(ctx, Signal, Target, Body); err != nil {
		return err
	}
	return nil
}

func (s *Service) PusherSignalGroupAvatar(ctx context.Context, groupId int64, avatar string) error {
	nowTime := s.getNowTime()
	action := &xproto.SignalUpdateGroupAvatar{
		Group:  groupId,
		Avatar: avatar,
		Time:   util.ToUInt64(nowTime),
	}
	body, err := proto.Marshal(action)
	if err != nil {
		return err
	}

	Signal := xproto.SignalType_UpdateGroupAvatar
	Target := util.ToString(groupId)
	Body := body

	if err = s.answerClient.GroupCastSignal(ctx, Signal, Target, Body); err != nil {
		return err
	}
	return nil
}
