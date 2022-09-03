package service

import (
	"context"

	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/dtalk/service/group/model/db"
	"github.com/txchat/dtalk/service/group/model/types"
)

// InviteGroupMembersSvc 邀请新群员
// 加群设置为 Any
//    - 邀请人为空, 直接加入
//    - 邀请人不为空, 直接加入
// 加群设置为 Apply
//    - 邀请人为空, 走审批流程
//    - 邀请人为普通人, 走审批流程
//    - 邀请人为管理员, 直接加入
// 加群设置为 Admin
//    - 邀请人为空, 拒绝加入
//    - 邀请人为普通人, 拒绝加入
//    - 邀请人为管理员, 直接加入
func (s *Service) InviteGroupMembersSvc(ctx context.Context, req *types.InviteGroupMembersRequest) (res *types.InviteGroupMembersResponse, err error) {
	//personId := req.Inviter.MemberId
	groupId := req.Id

	group, err := s.GetGroupInfoByGroupId(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	if req.Inviter.MemberId == "" {
		if group.GroupJoinType == biz.GroupJoinTypeAny {
			return s.inviteMember(ctx, group, req)
		} else if group.GroupJoinType == biz.GroupJoinTypeApply {
			return s.invite2apply(req.NewMemberIds, groupId, "")
		}
		return nil, xerror.NewError(xerror.GroupInvitePermissionDenied)
	}

	// 得到邀请人信息
	inviter, err := s.GetPersonByMemberIdAndGroupId(ctx, req.Inviter.MemberId, groupId)
	if err != nil {
		return nil, err
	}

	switch inviter.TryInvite(group) {
	case biz.InviteOk:
		return s.inviteMember(ctx, group, req)
	case biz.InviteApply:
		return s.invite2apply(req.NewMemberIds, groupId, inviter.GroupMemberId)
	case biz.InviteFail:
		return nil, xerror.NewError(xerror.GroupInvitePermissionDenied)
	}
	return nil, xerror.NewError(xerror.CodeInnerError)
}

// ExecInviteGroupMembers 执行邀请群成员操作
//func (s *Service) ExecInviteGroupMembers(members []*db.GroupMember) error {
//	tx, err := s.dao.NewTx()
//	if err != nil {
//		return err
//	}
//
//	defer tx.RollBack()
//
//	if err = s.InsertGroupMembers(tx, members); err != nil {
//		return err
//	}
//
//	if err = tx.Commit(); err != nil {
//		return err
//	}
//
//	return nil
//}

func (s *Service) invite2apply(MemberIds []string, groupId int64, inviterId string) (res *types.InviteGroupMembersResponse, err error) {
	newMemberIds := []string{}
	for _, memberId := range MemberIds {
		// 判断是否已经在群内
		if isExit, err := s.CheckInGroup(memberId, groupId); err != nil {
			continue
		} else if isExit == true {
			continue
		}

		newMemberIds = append(newMemberIds, memberId)
	}

	if len(newMemberIds) == 0 {
		return nil, xerror.NewError(xerror.GroupInviteNoMembers)
	}

	err = s.ExecCreateGroupApply(groupId, inviterId, newMemberIds, "")
	if err != nil {
		return nil, err
	}

	return &types.InviteGroupMembersResponse{}, nil
}

func (s *Service) inviteMember(ctx context.Context, group *biz.GroupInfo, req *types.InviteGroupMembersRequest) (res *types.InviteGroupMembersResponse, err error) {
	groupId := group.GroupId
	personId := req.Inviter.MemberId

	// 判断群人数上限
	if err = group.TryJoin(int32(len(req.NewMembers))); err != nil {
		return nil, err
	}

	// 插入新群员
	nowTime := s.getNowTime()
	newMembers := make([]*db.GroupMember, 0)
	for _, member := range req.NewMembers {
		// 判断是否已经在群内
		if isExit, err := s.CheckInGroup(member.MemberId, groupId); err != nil {
			continue
		} else if isExit == true {
			continue
		}

		newMembers = append(newMembers, &db.GroupMember{
			GroupId:               groupId,
			GroupMemberId:         member.MemberId,
			GroupMemberName:       member.MemberName,
			GroupMemberType:       member.MemberType,
			GroupMemberJoinTime:   nowTime,
			GroupMemberUpdateTime: nowTime,
		})
	}

	// 被邀请的人都已经在本群中
	if len(newMembers) == 0 {
		err = xerror.NewError(xerror.GroupInviteNoMembers)
		return nil, err
	}

	if err = s.AddGroupMembers(ctx, group.GroupId, newMembers, personId); err != nil {
		return nil, err
	}

	// ? 不知道有没有用
	res = &types.InviteGroupMembersResponse{
		Id:        req.Id,
		IdStr:     util.MustToString(groupId),
		MemberNum: group.GroupMemberNum + int32(len(newMembers)),
		//Inviter:    req.Inviter,
		//NewMembers: model.GroupMemberConvertGroupMemberInfo(newMembers),
	}
	return res, nil
}

func (s *Service) InviteMembers(ctx context.Context, group *biz.GroupInfo, newMemberIds []string) error {
	groupId := group.GroupId

	// 判断群人数上限
	if err := group.TryJoin(int32(len(newMemberIds))); err != nil {
		return err
	}

	// 插入新群员
	members := make([]*biz.GroupMember, 0, len(newMemberIds))
	for _, memberId := range newMemberIds {
		members = append(members, &biz.GroupMember{
			GroupId:         groupId,
			GroupMemberId:   memberId,
			GroupMemberName: "",
			GroupMemberType: biz.GroupTypeNormal,
		})
	}

	err := s.AddMembers(ctx, group, members)
	if err != nil {
		return err
	}

	return nil
}
