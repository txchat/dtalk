package logic

import (
	"context"
	pb "github.com/txchat/dtalk/service/group/api"
	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/dtalk/service/group/service"
)

type ForceJoinGroupsLogic struct {
	ctx context.Context
	svc *service.Service
}

func NewForceJoinGroupsLogic(ctx context.Context, svc *service.Service) *ForceJoinGroupsLogic {
	return &ForceJoinGroupsLogic{
		ctx: ctx,
		svc: svc,
	}
}

// ForceJoinGroups 一个人加入多个群
// 相当于多次 AddMember
func (l *ForceJoinGroupsLogic) ForceJoinGroups(req *pb.ForceJoinGroupsReq) (*pb.ForceJoinGroupsResp, error) {
	_, err := FilteredMemberId(req.Member.Id)
	if err != nil {
		return nil, err
	}

	groups := make([]*biz.GroupInfo, 0, len(req.GroupIds))
	for _, groupId := range req.GroupIds {
		group, err := l.svc.GetGroupInfoByGroupId(l.ctx, groupId)
		if err != nil {
			continue
		}

		groups = append(groups, group)
	}

	members := make([]*biz.GroupMember, 0, len(groups))
	for _, group := range groups {
		members = append(members, &biz.GroupMember{
			GroupId:         group.GroupId,
			GroupMemberId:   req.Member.Id,
			GroupMemberName: req.Member.Name,
			GroupMemberType: biz.GroupMemberTypeNormal,
		})
	}

	l.svc.JoinGroups(l.ctx, members)

	return &pb.ForceJoinGroupsResp{}, nil
}
