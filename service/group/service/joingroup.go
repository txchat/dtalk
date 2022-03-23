package service

import (
	"context"
	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/dtalk/service/group/model/db"
)

// JoinGroupSvc 加入群聊, grpc 专属
//func (s *Service) JoinGroupSvc(ctx context.Context, req *types.JoinGroupReq) (res *types.JoinGroupResp, err error) {
//	personId := req.PersonId
//	groupId := req.Id
//
//	group, err := s.GetGroupInfoByGroupId(ctx, groupId)
//	if err != nil {
//		return nil, err
//	}
//
//	// 判断群人数上限
//	if err = group.TryJoin(1); err != nil {
//		return nil, err
//	}
//
//	// 判断是否已经在群内
//	if isExit, err := s.CheckInGroup(personId, groupId); err != nil {
//		return nil, err
//	} else if isExit == true {
//		return nil, xerror.NewError(xerror.GroupInviteMemberExist)
//	}
//
//	nowTime := s.getNowTime()
//	newMember := &db.GroupMember{
//		GroupId:               groupId,
//		GroupMemberId:         personId,
//		GroupMemberName:       "",
//		GroupMemberType:       0,
//		GroupMemberJoinTime:   nowTime,
//		GroupMemberUpdateTime: nowTime,
//	}
//
//	if err = s.AddGroupMembers(ctx, group.GroupId, []*db.GroupMember{newMember}, group.GroupOwnerId); err != nil {
//		return nil, err
//	}
//
//	res = &types.JoinGroupResp{
//		Id:    groupId,
//		IdStr: util.ToString(groupId),
//	}
//
//	return
//}

// ExecJoinGroupMembers 执行邀请群成员操作
func (s *Service) ExecJoinGroupMembers(members []*db.GroupMember) error {
	tx, err := s.dao.NewTx()
	if err != nil {
		return err
	}

	defer tx.RollBack()

	if err = s.InsertGroupMembers(tx, members); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetFilteredGroupMembers(ctx context.Context, group *biz.GroupInfo, memberIds []string) []*biz.GroupMember {
	members := make([]*biz.GroupMember, 0, len(memberIds))
	for _, memberId := range memberIds {
		member, err := s.GetMemberByMemberIdAndGroupId(ctx, memberId, group.GroupId)
		if err != nil {
			continue
		}

		members = append(members, member)
	}

	return members
}

// FilteredGroupMembers  过滤已经在群里的成员
func (s *Service) FilteredGroupMembers(members []*biz.GroupMember) []*biz.GroupMember {
	newMembers := make([]*biz.GroupMember, 0)
	for _, member := range members {
		// 判断是否已经在群内
		if isExit, err := s.CheckInGroup(member.GroupMemberId, member.GroupId); err != nil {
			continue
		} else if isExit == true {
			continue
		}

		newMembers = append(newMembers, member)
	}

	return newMembers
}

func (s *Service) AddMembers(ctx context.Context, group *biz.GroupInfo, members []*biz.GroupMember) error {
	members = s.FilteredGroupMembers(members)

	if len(members) == 0 {
		return nil
	}

	nowTime := s.getNowTime()
	newMembers := make([]*db.GroupMember, 0, len(members))
	for _, member := range members {
		newMembers = append(newMembers, &db.GroupMember{
			GroupId:               member.GroupId,
			GroupMemberId:         member.GroupMemberId,
			GroupMemberName:       member.GroupMemberName,
			GroupMemberType:       member.GroupMemberType,
			GroupMemberJoinTime:   nowTime,
			GroupMemberUpdateTime: nowTime,
		})
	}

	if err := s.AddGroupMembers(ctx, group.GroupId, newMembers, s.GetOpe(ctx)); err != nil {
		return err
	}

	return nil
}

func (s *Service) JoinGroups(ctx context.Context, members []*biz.GroupMember) {
	log := s.GetLogWithTrace(ctx)
	members = s.FilteredGroupMembers(members)

	if len(members) == 0 {
		return
	}

	nowTime := s.getNowTime()
	for _, member := range members {
		newMember := &db.GroupMember{
			GroupId:               member.GroupId,
			GroupMemberId:         member.GroupMemberId,
			GroupMemberName:       member.GroupMemberName,
			GroupMemberType:       member.GroupMemberType,
			GroupMemberJoinTime:   nowTime,
			GroupMemberUpdateTime: nowTime,
		}

		if err := s.AddGroupMembers(ctx, member.GroupId, []*db.GroupMember{newMember}, s.GetOpe(ctx)); err != nil {
			log.Error().Err(err).Msg("JoinGroups")
		}
	}
}
