package logic

import (
	"context"
	xerror "github.com/txchat/dtalk/pkg/error"
	pb "github.com/txchat/dtalk/service/group/api"
	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/dtalk/service/group/service"
)

type SetAdminLogic struct {
	ctx context.Context
	svc *service.Service
}

func NewSetAdminLogic(ctx context.Context, svc *service.Service) *SetAdminLogic {
	return &SetAdminLogic{
		ctx: ctx,
		svc: svc,
	}
}

// SetAdmin 设置管理员
func (l *SetAdminLogic) SetAdmin(req *pb.SetAdminReq) (*pb.SetAdminResp, error) {
	groupId := req.GroupId
	personId := req.PersonId
	memberId := req.MemberId
	memberType := int32(req.GroupMemberType)

	group, err := l.svc.GetGroupInfoByGroupId(l.ctx, groupId)
	if err != nil {
		return nil, err
	}

	person, err := l.svc.GetPersonByMemberIdAndGroupId(l.ctx, personId, groupId)
	if err != nil {
		return nil, err
	}
	// 只有群主可以设置管理员
	if personId != group.GroupOwnerId || person.GroupMemberType != biz.GroupMemberTypeOwner || memberId == personId {
		err = xerror.NewError(xerror.GroupOwnerSetAdmin)
		return nil, err
	}

	member, err := l.svc.GetMemberByMemberIdAndGroupId(l.ctx, memberId, groupId)
	if err != nil {
		return nil, err
	}

	if err = group.TrySetAdmin(); memberType == biz.GroupMemberTypeAdmin && err != nil {
		return nil, err
	}

	err = l.svc.SetAdmin(l.ctx, group, member, memberType)
	if err != nil {
		return nil, err
	}

	return &pb.SetAdminResp{}, nil
}
