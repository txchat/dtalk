package logic

import (
	"context"
	xerror "github.com/txchat/dtalk/pkg/error"
	pb "github.com/txchat/dtalk/service/group/api"
	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/dtalk/service/group/service"
)

type UpdateGroupMemberMuteTimeLogic struct {
	ctx context.Context
	svc *service.Service
}

func NewUpdateGroupMemberMuteTimeLogic(ctx context.Context, svc *service.Service) *UpdateGroupMemberMuteTimeLogic {
	return &UpdateGroupMemberMuteTimeLogic{
		ctx: ctx,
		svc: svc,
	}
}

// UpdateGroupMemberMuteTime 更新群成员禁言时间
func (l *UpdateGroupMemberMuteTimeLogic) UpdateGroupMemberMuteTime(req *pb.UpdateGroupMemberMuteTimeReq) (*pb.UpdateGroupMemberMuteTimeResp, error) {
	groupId := req.GroupId
	memberIds := req.MemberIds
	personId := req.PersonId
	muteTime := req.MuteTime
	var members []*biz.GroupMember

	group, err := l.svc.GetGroupInfoByGroupId(l.ctx, groupId)
	if err != nil {
		return nil, err
	}

	person, err := l.svc.GetPersonByMemberIdAndGroupId(l.ctx, personId, groupId)
	if err != nil {
		return nil, err
	}
	if err = person.IsAdmin(); err != nil {
		return nil, err
	}

	// 过滤
	for _, memberId := range memberIds {
		member, err := l.svc.GetMemberByMemberIdAndGroupId(l.ctx, memberId, groupId)
		if err != nil {
			return nil, err
		}
		if err := member.IsAdmin(); err == nil {
			return nil, xerror.NewError(xerror.GroupMutePermission)
		}

		members = append(members, member)
	}

	members, err = l.svc.UpdateMembersMuteTime(l.ctx, group, members, muteTime)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateGroupMemberMuteTimeResp{
		Members: NewRPCGroupMemberInfos(members),
	}, nil
}
