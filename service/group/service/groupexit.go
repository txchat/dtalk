package service

import (
	"context"

	"github.com/txchat/dtalk/pkg/contextx"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/dtalk/service/group/model/types"
)

// GroupExitHttp 退群
func (s *Service) GroupExitHttp(ctx context.Context, req *types.GroupExitRequest) (res *types.GroupExitResponse, err error) {
	group, err := s.GetGroupInfoByGroupId(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	//person, err := s.GetPersonByMemberIdAndGroupId(personId, groupId)
	person, err := s.GetPersonByMemberIdAndGroupId(ctx, req.PersonId, req.Id)
	if err != nil {
		return nil, err
	}

	if err = person.IsOwner(); err == nil {
		return nil, xerror.NewError(xerror.GroupOwnerExit)
	}

	err = s.ExitGroup(ctx, group, person)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Service) ExitGroup(ctx context.Context, group *biz.GroupInfo, member *biz.GroupMember) error {
	groupId := group.GroupId
	memberId := member.GroupMemberId
	log := s.GetLogWithTrace(ctx)

	// 执行退群
	if err := s.ExecGroupExit(ctx, groupId, memberId); err != nil {
		return err
	}

	// 更新群人数
	_, err := s.UpdateGroupInfoMemberNum(ctx, groupId)
	if err != nil {
		log.Error().Err(err).Msg("GroupExitHttp exec")
	}

	go func() {
		// 发送给 pusher
		if err = s.PusherSignalLeave(contextx.ValueOnlyFrom(ctx), groupId, []string{memberId}); err != nil {
			log.Error().Err(err).Msg("GroupExitHttp pusher")
		}
		// 发送给 logic
		if err = s.LogicNoticeLeave(contextx.ValueOnlyFrom(ctx), groupId, []string{memberId}); err != nil {
			log.Error().Err(err).Msg("GroupExitHttp logic")
		}
		// 发送给 alert
		if err = s.NoticeMsgSignOutGroup(contextx.ValueOnlyFrom(ctx), groupId, memberId); err != nil {
			log.Error().Err(err).Msg("GroupExitHttp alert")
		}
	}()

	return nil
}

// ExecGroupExit 执行退群操作
func (s *Service) ExecGroupExit(ctx context.Context, groupId int64, memberId string) error {
	err := s.dao.UpdateGroupMemberType(ctx, groupId, memberId, biz.GroupMemberTypeOther)
	if err != nil {
		return err
	}
	return nil
}
