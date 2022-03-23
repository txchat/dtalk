package service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/txchat/dtalk/pkg/contextx"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/service/group/model"
	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/dtalk/service/group/model/db"
)

// GetMemberByMemberIdAndGroupId 查询对方在群里的信息
func (s *Service) GetMemberByMemberIdAndGroupId(ctx context.Context, memberId string, groupId int64) (*biz.GroupMember, error) {
	member, err := s.dao.GetGroupMemberByGroupIdAndMemberId(ctx, groupId, memberId)
	if err != nil {
		if errors.Is(err, model.ErrRecordNotExist) {
			return nil, xerror.NewError(xerror.GroupMemberNotExist)
		}
		return nil, err
	}

	return member, nil
}

// GetPersonByMemberIdAndGroupId 查询本人在群里的信息
func (s *Service) GetPersonByMemberIdAndGroupId(ctx context.Context, memberId string, groupId int64) (*biz.GroupMember, error) {
	member, err := s.dao.GetGroupMemberByGroupIdAndMemberId(ctx, groupId, memberId)
	if err != nil {
		if errors.Is(err, model.ErrRecordNotExist) {
			return nil, xerror.NewError(xerror.GroupPersonNotExist)
		}
		return nil, err
	}

	return member, nil
}

// GetGroupMembersByGroupIdWithLimit 查询群内前 n 个群成员信息
func (s *Service) GetGroupMembersByGroupIdWithLimit(groupId, n, m int64) ([]*biz.GroupMember, error) {
	members, err := s.dao.GetGroupMembersByGroupIdWithLimit(groupId, n, m)
	if err != nil {
		return nil, err
	}
	return members, nil
}

// GetMembersByGroupId 查询群里的所有成员信息
func (s *Service) GetMembersByGroupId(groupId int64) ([]*biz.GroupMember, error) {
	res, err := s.dao.GetMembersByGroupId(groupId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetGroupMemberMuteTime 查询群成员禁言时间
func (s *Service) GetGroupMemberMuteTime(groupId int64, memberId string) (int64, error) {
	muteTime, err := s.dao.GetGroupMemberMuteTime(groupId, memberId)
	if err != nil {
		return 0, err
	}
	return muteTime, nil
}

// AddGroupMembers 加群并发送通知
func (s *Service) AddGroupMembers(ctx context.Context, groupId int64, members []*db.GroupMember, opeId string) error {
	log := s.GetLogWithTrace(ctx)

	if err := s.ExecJoinGroupMembers(members); err != nil {
		return err
	}

	// 更新群人数
	_, err := s.UpdateGroupInfoMemberNum(ctx, groupId)
	if err != nil {
		// todo
		log.Error().Err(err).Msg("AddGroupMembers")
	}

	memberIds := make([]string, 0, len(members))
	for _, member := range members {
		memberIds = append(memberIds, member.GroupMemberId)
	}

	go s.NoticeInviteMembers(contextx.ValueOnlyFrom(ctx), groupId, opeId, memberIds)

	return nil
}
