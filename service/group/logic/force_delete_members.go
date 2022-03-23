package logic

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/txchat/dtalk/service/group/model/biz"

	pb "github.com/txchat/dtalk/service/group/api"

	"github.com/txchat/dtalk/service/group/service"
)

type ForceDeleteMembersLogic struct {
	ctx context.Context
	svc *service.Service
	log zerolog.Logger
}

func NewForceDeleteMembersLogic(ctx context.Context, svc *service.Service) *ForceDeleteMembersLogic {
	return &ForceDeleteMembersLogic{
		ctx: ctx,
		svc: svc,
		log: svc.GetLog(),
	}
}

// ForceDeleteMembers 多个人退出同一个群
func (l *ForceDeleteMembersLogic) ForceDeleteMembers(req *pb.ForceDeleteMembersReq) (*pb.ForceDeleteMembersResp, error) {
	groupId := req.GroupId

	// 判断群是否存在
	group, err := l.svc.GetGroupInfoByGroupId(l.ctx, groupId)
	if err != nil {
		return nil, err
	}

	needDeleteMembers := l.svc.GetFilteredGroupMembers(l.ctx, group, FilteredMemberIds(req.MemberIds))
	if len(needDeleteMembers) == 0 {
		return &pb.ForceDeleteMembersResp{}, nil
	}

	members, err := l.svc.GetGroupMembersByGroupIdWithLimit(group.GroupId, 0, int64(len(needDeleteMembers))+1)
	if err != nil {
		return nil, err
	}

	if len(needDeleteMembers) == len(members) {
		// 解散群
		err = l.svc.GroupDisband(l.ctx, groupId, l.svc.GetOpe(l.ctx))
		if err != nil {
			return nil, err
		}

		return &pb.ForceDeleteMembersResp{}, nil
	}

	owner, ok := hasOwner(needDeleteMembers)
	if ok {
		CandidateOwner := getCandidateOwner(members, needDeleteMembers)
		err = l.svc.ChangeOwner(l.ctx, group, owner, CandidateOwner)
		if err != nil {
			return nil, err
		}
	}

	// 退群
	err = l.svc.RemoveGroupMembers(l.ctx, group, needDeleteMembers)
	if err != nil {
		return nil, err
	}

	return &pb.ForceDeleteMembersResp{}, nil
}

func hasOwner(members []*biz.GroupMember) (*biz.GroupMember, bool) {
	for _, member := range members {
		if err := member.IsOwner(); err == nil {
			return member, true
		}
	}

	return nil, false
}

func getCandidateOwner(members []*biz.GroupMember, deleteMembers []*biz.GroupMember) *biz.GroupMember {
	for i := 0; i < len(members); i++ {
		isDelete := false
		for j := 0; j < len(deleteMembers); j++ {
			if members[i].GroupMemberId == deleteMembers[j].GroupMemberId {
				isDelete = true
				break
			}
		}

		if !isDelete {
			return members[i]
		}
	}

	return &biz.GroupMember{}
}
