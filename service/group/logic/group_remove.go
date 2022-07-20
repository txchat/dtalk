package logic

import (
	"context"

	pb "github.com/txchat/dtalk/service/group/api"
	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/dtalk/service/group/service"
)

type GroupRemoveLogic struct {
	ctx context.Context
	svc *service.Service
}

func NewGroupRemoveLogic(ctx context.Context, svc *service.Service) *GroupRemoveLogic {
	return &GroupRemoveLogic{
		ctx: ctx,
		svc: svc,
	}
}

// GroupRemove 踢人
func (l *GroupRemoveLogic) GroupRemove(req *pb.GroupRemoveReq) (*pb.GroupRemoveResp, error) {
	groupId := req.GroupId
	memberIds := req.MemberIds
	personId := req.PersonId

	// 判断一下该群是否存在
	group, err := l.svc.GetGroupInfoByGroupId(l.ctx, groupId)
	if err != nil {
		return nil, err
	}

	// 判断踢人者权限
	person, err := l.svc.GetPersonByMemberIdAndGroupId(l.ctx, personId, groupId)
	if err != nil {
		return nil, err
	}
	if err = person.IsAdmin(); err != nil {
		return nil, err
	}

	// 过滤可以踢人的列表
	needDeleteMembers := l.svc.GetFilteredGroupMembers(l.ctx, group, FilteredMemberIds(memberIds))
	if len(needDeleteMembers) == 0 {
		return &pb.GroupRemoveResp{MemberNum: group.GroupMemberNum}, nil
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
	if err = l.svc.RemoveGroupMembers(l.ctx, group, canRemoveMembers); err != nil {
		return nil, err
	}

	return &pb.GroupRemoveResp{
		MemberNum: group.GroupMemberNum - int32(len(canRemoveMemberIds)),
		MemberIds: canRemoveMemberIds,
	}, nil
}
