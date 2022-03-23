package logic

import (
	"context"
	pb "github.com/txchat/dtalk/service/group/api"
	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/dtalk/service/group/service"
)

type CreateGroupLogic struct {
	ctx context.Context
	svc *service.Service
}

func NewCreateGroupLogic(ctx context.Context, svc *service.Service) *CreateGroupLogic {
	return &CreateGroupLogic{
		ctx: ctx,
		svc: svc,
	}
}

// CreateGroup 创建群聊
func (l *CreateGroupLogic) CreateGroup(req *pb.CreateGroupReq) (*pb.CreateGroupResp, error) {
	var err error
	req.Owner.Id, err = FilteredMemberId(req.Owner.Id)
	if err != nil {
		return nil, err
	}

	group := &biz.GroupInfo{
		GroupName: req.Name,
		GroupType: int32(req.GroupType),
	}

	owner := &biz.GroupMember{
		GroupMemberId:   req.Owner.Id,
		GroupMemberName: req.Owner.Name,
	}

	members := make([]*biz.GroupMember, 0, len(req.Members))
	for _, member := range req.Members {
		if member.Id == owner.GroupMemberId {
			continue
		}

		member.Id, err = FilteredMemberId(member.Id)
		if err != nil {
			return nil, err
		}

		members = append(members, &biz.GroupMember{
			GroupMemberId:   member.Id,
			GroupMemberName: member.Name,
		})
	}

	groupId, err := l.svc.CreateGroup(l.ctx, group, owner, members)
	if err != nil {
		return nil, err
	}

	return &pb.CreateGroupResp{
		GroupId: groupId,
	}, nil
}
