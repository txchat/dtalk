package logic

import (
	"context"

	xerror "github.com/txchat/dtalk/pkg/error"
	pb "github.com/txchat/dtalk/service/group/api"
	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/dtalk/service/group/service"
)

type InviteGroupMembersLogic struct {
	ctx context.Context
	svc *service.Service
}

func NewInviteGroupMembersLogic(ctx context.Context, svc *service.Service) *InviteGroupMembersLogic {
	return &InviteGroupMembersLogic{
		ctx: ctx,
		svc: svc,
	}
}

// InviteGroupMembers 邀请新成员
func (l *InviteGroupMembersLogic) InviteGroupMembers(req *pb.InviteGroupMembersReq) (*pb.InviteGroupMembersResp, error) {
	//personId := req.Inviter.MemberId
	groupId := req.GroupId
	inviterId := req.InviterId
	newMemberIds := FilteredMemberIds(req.MemberIds)

	group, err := l.svc.GetGroupInfoByGroupId(l.ctx, groupId)
	if err != nil {
		return nil, err
	}

	if inviterId == "" {
		switch group.GroupJoinType {
		case biz.GroupJoinTypeAny:
			err = l.svc.InviteMembers(l.ctx, group, newMemberIds)
			if err != nil {
				return nil, err
			}
			return &pb.InviteGroupMembersResp{}, nil
		default:
			return nil, xerror.NewError(xerror.GroupInvitePermissionDenied)
		}
	}

	// 得到邀请人信息
	inviter, err := l.svc.GetPersonByMemberIdAndGroupId(l.ctx, inviterId, groupId)
	if err != nil {
		return nil, err
	}

	switch inviter.TryInvite(group) {
	case biz.InviteOk:
		err = l.svc.InviteMembers(l.ctx, group, newMemberIds)
		if err != nil {
			return nil, err
		}
		return &pb.InviteGroupMembersResp{}, nil
	case biz.InviteApply:
		return nil, xerror.NewError(xerror.GroupInvitePermissionDenied)
	case biz.InviteFail:
		return nil, xerror.NewError(xerror.GroupInvitePermissionDenied)
	default:
		return nil, xerror.NewError(xerror.GroupInvitePermissionDenied)
	}
}
