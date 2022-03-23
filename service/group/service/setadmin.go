package service

import (
	"context"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/dtalk/service/group/model/db"
	"github.com/txchat/dtalk/service/group/model/types"
)

// SetAdminSvc 设置管理员
func (s *Service) SetAdminSvc(ctx context.Context, req *types.SetAdminRequest) (res *types.SetAdminResponse, err error) {
	groupId := req.Id
	personId := req.PersonId
	memberId := req.MemberId
	memberType := req.MemberType
	log := s.GetLogWithTrace(ctx)

	group, err := s.GetGroupInfoByGroupId(ctx, groupId)
	if err != nil {
		return nil, err
	}

	person, err := s.GetPersonByMemberIdAndGroupId(ctx, personId, groupId)
	if err != nil {
		return nil, err
	}
	// 只有群主可以设置管理员
	if personId != group.GroupOwnerId || person.GroupMemberType != biz.GroupMemberTypeOwner || memberId == personId {
		err = xerror.NewError(xerror.GroupOwnerSetAdmin)
		return nil, err
	}

	_, err = s.GetMemberByMemberIdAndGroupId(ctx, memberId, groupId)
	if err != nil {
		return nil, err
	}

	memberType, err = s.CheckSetAdminType(memberType)
	if err != nil {
		return nil, err
	}

	if err = group.TrySetAdmin(); memberType == biz.GroupMemberTypeAdmin && err != nil {
		return nil, err
	}

	if err = s.UpdateGroupMemberType(ctx, groupId, memberId, memberType); err != nil {
		return nil, err
	}

	// 发送给 pusher
	if err = s.PusherSignalMemberType(ctx, groupId, memberId, memberType); err != nil {
		log.Error().Err(err).Msg("SetAdminSvc pusher")
	}

	if memberType == biz.GroupMuteTypeAdmin {
		// 解除新管理员禁言
		nowTime := s.getNowTime()
		groupMemberMute := make([]*db.GroupMemberMute, 1, 1)
		groupMemberMute[0] = &db.GroupMemberMute{
			GroupId:                   groupId,
			GroupMemberId:             memberId,
			GroupMemberMuteTime:       0,
			GroupMemberMuteUpdateTime: nowTime,
		}
		if err = s.execSetMemberMuteTimes(ctx, groupMemberMute); err != nil {
			return nil, err
		}

		// 发送给 pusher
		if err = s.PusherSignalMemberMuteTime(ctx, groupId, []string{memberId}, 0); err != nil {
			log.Error().Err(err).Msg("UpdateMembersMuteTimeSvc pusher")
		}
	}

	return res, nil
}

// CheckSetAdminType 检查SetAdmin是否合法
func (s *Service) CheckSetAdminType(memberType int32) (int32, error) {
	switch memberType {
	case biz.GroupMemberTypeNormal:
		return biz.GroupMemberTypeNormal, nil
	case biz.GroupMuteTypeAdmin:
		return biz.GroupMemberTypeAdmin, nil
	default:
		return 0, xerror.NewError(xerror.ParamsError)
	}
}

// UpdateGroupMemberType 更新群成员类型
func (s *Service) UpdateGroupMemberType(ctx context.Context, groupId int64, memberId string, memberType int32) error {
	err := s.dao.UpdateGroupMemberType(ctx, groupId, memberId, memberType)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) SetAdmin(ctx context.Context, group *biz.GroupInfo, member *biz.GroupMember, memberType int32) error {
	log := s.GetLogWithTrace(ctx)
	groupId := group.GroupId
	memberId := member.GroupMemberId

	if err := s.UpdateGroupMemberType(ctx, groupId, memberId, memberType); err != nil {
		return err
	}

	// 发送给 pusher
	if err := s.PusherSignalMemberType(ctx, groupId, memberId, memberType); err != nil {
		log.Error().Err(err).Msg("SetAdmin pusher")
	}

	if memberType == biz.GroupMuteTypeAdmin {
		// 解除新管理员禁言
		groupMemberMute := make([]*db.GroupMemberMute, 1, 1)
		groupMemberMute[0] = &db.GroupMemberMute{
			GroupId:             groupId,
			GroupMemberId:       memberId,
			GroupMemberMuteTime: 0,
		}
		if err := s.execSetMemberMuteTimes(ctx, groupMemberMute); err != nil {
			log.Error().Err(err).Msg("execSetMemberMuteTimes pusher")
		}

		// 发送给 pusher
		if err := s.PusherSignalMemberMuteTime(ctx, groupId, []string{memberId}, 0); err != nil {
			log.Error().Err(err).Msg("PusherSignalMemberMuteTime pusher")
		}
	}

	return nil
}
