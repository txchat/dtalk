package logic

import (
	"context"
	pb "github.com/txchat/dtalk/service/group/api"
	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/dtalk/service/group/service"
)

type ForceAddMembersLogic struct {
	ctx context.Context
	svc *service.Service
}

func NewForceAddMembersLogic(ctx context.Context, svc *service.Service) *ForceAddMembersLogic {
	return &ForceAddMembersLogic{
		ctx: ctx,
		svc: svc,
	}
}

// ForceAddMembers 多个人加入一个群
// 无视操作者是否在群里, 是否有管理权限, 群人数是否已满
// 强行拉 member 进群
func (l *ForceAddMembersLogic) ForceAddMembers(req *pb.ForceAddMembersReq) (*pb.ForceAddMembersResp, error) {
	group, err := l.svc.GetGroupInfoByGroupId(l.ctx, req.GroupId)
	if err != nil {
		return nil, err
	}

	members := make([]*biz.GroupMember, 0, len(req.Members))
	for _, member := range req.Members {
		_, err := FilteredMemberId(member.Id)
		if err != nil {
			continue
		}

		members = append(members, &biz.GroupMember{
			GroupId:         group.GroupId,
			GroupMemberId:   member.Id,
			GroupMemberName: member.Name,
			GroupMemberType: biz.GroupTypeNormal,
		})
	}

	err = l.svc.AddMembers(l.ctx, group, members)
	if err != nil {
		return nil, err
	}

	return &pb.ForceAddMembersResp{}, nil
}
