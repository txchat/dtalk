package logic

import (
	"context"
	pb "github.com/txchat/dtalk/service/group/api"
	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/dtalk/service/group/service"
)

type ForceAddMemberLogic struct {
	ctx context.Context
	svc *service.Service
}

func NewForceAddMemberLogic(ctx context.Context, svc *service.Service) *ForceAddMemberLogic {
	return &ForceAddMemberLogic{
		ctx: ctx,
		svc: svc,
	}
}

// ForceAddMember 多个人加入一个群
// 无视操作者是否在群里, 是否有管理权限, 群人数是否已满
// 强行拉 member 进群
func (l *ForceAddMemberLogic) ForceAddMember(req *pb.ForceAddMemberReq) (*pb.ForceAddMemberResp, error) {
	group, err := l.svc.GetGroupInfoByGroupId(l.ctx, req.GroupId)
	if err != nil {
		return nil, err
	}

	members := make([]*biz.GroupMember, 0)
	_, err = FilteredMemberId(req.MemberId)

	members = append(members, &biz.GroupMember{
		GroupId:         group.GroupId,
		GroupMemberId:   req.MemberId,
		GroupMemberName: "",
		GroupMemberType: biz.GroupTypeNormal,
	})

	err = l.svc.AddMembers(l.ctx, group, members)
	if err != nil {
		return nil, err
	}

	return &pb.ForceAddMemberResp{}, nil
}
