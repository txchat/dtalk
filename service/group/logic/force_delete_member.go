package logic

import (
	"context"

	"github.com/txchat/dtalk/service/group/model/biz"

	pb "github.com/txchat/dtalk/service/group/api"

	"github.com/txchat/dtalk/service/group/service"
)

type ForceDeleteMemberLogic struct {
	ctx context.Context
	svc *service.Service
}

func NewForceDeleteMemberLogic(ctx context.Context, svc *service.Service) *ForceDeleteMemberLogic {
	return &ForceDeleteMemberLogic{
		ctx: ctx,
		svc: svc,
	}
}

// ForceDeleteMember 一个人退出一个群
func (l *ForceDeleteMemberLogic) ForceDeleteMember(req *pb.ForceDeleteMemberReq) (*pb.ForceDeleteMemberResp, error) {
	groupId := req.GroupId
	memberId, err := FilteredMemberId(req.MemberId)
	if err != nil {
		return nil, err
	}

	// 判断群是否存在
	group, err := l.svc.GetGroupInfoByGroupId(l.ctx, groupId)
	if err != nil {
		return nil, err
	}

	// 判断这个人是否在群里
	deleteMember, err := l.svc.GetMemberByMemberIdAndGroupId(l.ctx, memberId, groupId)
	if err != nil {
		return nil, err
	}

	members, err := l.svc.GetGroupMembersByGroupIdWithLimit(group.GroupId, 0, 2)
	if err != nil {
		return nil, err
	}

	// 如果退群的人是群主
	if group.GroupOwnerId == deleteMember.GroupMemberId {
		// 如果群内有群主之外的人
		if len(members) > 1 {
			// 转让群主, 再到下面一起退群
			// todo
			err = l.svc.ChangeOwner(l.ctx, group, deleteMember, members[1])
			if err != nil {
				return nil, err
			}
		} else {
			// 解散群
			err = l.svc.GroupDisband(l.ctx, groupId, memberId)
			if err != nil {
				return nil, err
			}

			return &pb.ForceDeleteMemberResp{}, nil
		}
	}

	// 退群
	err = l.svc.RemoveGroupMembers(l.ctx, group, []*biz.GroupMember{deleteMember})
	if err != nil {
		return nil, err
	}

	return &pb.ForceDeleteMemberResp{}, nil
}
