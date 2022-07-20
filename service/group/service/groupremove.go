package service

import (
	"context"

	"github.com/txchat/dtalk/pkg/contextx"
	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/dtalk/service/group/model/db"
	"github.com/txchat/dtalk/service/group/model/types"
)

// GroupRemoveSvc 踢人
func (s *Service) GroupRemoveSvc(ctx context.Context, req *types.GroupRemoveRequest) (res *types.GroupRemoveResponse, err error) {
	groupId := req.Id
	memberIds := req.MemberIds
	personId := req.PersonId

	// 判断一下该群是否存在
	group, err := s.GetGroupInfoByGroupId(ctx, groupId)
	if err != nil {
		return nil, err
	}

	// 判断踢人者权限
	person, err := s.GetPersonByMemberIdAndGroupId(ctx, personId, groupId)
	if err != nil {
		return nil, err
	}
	if err = person.IsAdmin(); err != nil {
		return nil, err
	}

	// 过滤可以踢人的列表
	needDeleteMembers := s.GetFilteredGroupMembers(ctx, group, memberIds)
	if len(needDeleteMembers) == 0 {
		return &types.GroupRemoveResponse{MemberNum: group.GroupMemberNum}, nil
	}

	canRemoveMemberIds := make([]string, 0)
	canRemoveMembers := make([]*biz.GroupMember, 0)
	for _, member := range needDeleteMembers {
		if err := person.RemoveOneMember(member); err != nil {
			return nil, err
		}

		canRemoveMemberIds = append(canRemoveMemberIds, member.GroupMemberId)
		canRemoveMembers = append(canRemoveMembers, member)
	}

	// 执行踢人
	if err = s.RemoveGroupMembers(ctx, group, canRemoveMembers); err != nil {
		return nil, err
	}

	res = &types.GroupRemoveResponse{
		MemberIds: canRemoveMemberIds,
	}
	return res, nil
}

// RemoveGroupMembers 执行踢人操作
func (s *Service) RemoveGroupMembers(ctx context.Context, group *biz.GroupInfo, members []*biz.GroupMember) error {
	if len(members) == 0 {
		return nil
	}
	groupId := group.GroupId
	memberIds := make([]string, 0, len(members))
	for _, member := range members {
		memberIds = append(memberIds, member.GroupMemberId)
	}

	log := s.GetLogWithTrace(ctx)
	nowTime := s.getNowTime()

	tx, err := s.dao.NewTx()
	if err != nil {
		return err
	}
	defer tx.RollBack()

	for _, memberId := range memberIds {
		groupMemberInfo := &db.GroupMember{
			GroupId:               groupId,
			GroupMemberId:         memberId,
			GroupMemberUpdateTime: nowTime,
			GroupMemberType:       biz.GroupMemberTypeOther,
		}
		if _, _, err = s.dao.UpdateGroupMemberTypeWithTx(tx, groupMemberInfo); err != nil {
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	// 更新群人数
	_, err = s.UpdateGroupInfoMemberNum(ctx, groupId)
	if err != nil {
		log.Error().Err(err).Msg("GroupRemoveSvc")
	}

	go func() {
		// 发送给 alert
		if err = s.NoticeMsgKickOutGroup(contextx.ValueOnlyFrom(ctx), groupId, s.GetOpe(ctx), memberIds); err != nil {
			log.Error().Err(err).Msg("GroupRemoveSvc alert")
		}
		// 发送个 pusher
		if err = s.PusherSignalLeave(contextx.ValueOnlyFrom(ctx), groupId, memberIds); err != nil {
			log.Error().Msg("GroupRemoveSvc pusher")
		}
		// 发送给 logic
		if err = s.LogicNoticeLeave(contextx.ValueOnlyFrom(ctx), groupId, memberIds); err != nil {
			log.Error().Msg("GroupRemoveSvc logic")
		}
	}()

	return nil
}
