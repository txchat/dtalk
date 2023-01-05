package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/group/group"
	"github.com/txchat/dtalk/app/services/group/internal/model"
	"github.com/txchat/dtalk/app/services/group/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupLimitedMembersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupLimitedMembersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupLimitedMembersLogic {
	return &GroupLimitedMembersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GroupLimitedMembersLogic) GroupLimitedMembers(in *group.GroupLimitedMembersReq) (*group.GroupLimitedMembersResp, error) {
	gid := in.GetGid()

	var membersRepo []*model.GroupMember
	var err error
	if in.Num == nil {
		membersRepo, err = l.svcCtx.Repo.GetUnLimitedMembers(gid)
	} else {
		membersRepo, err = l.svcCtx.Repo.GetLimitedMembers(gid, 0, in.GetNum())
	}
	if err != nil {
		return nil, err
	}

	members := make([]*group.MemberInfo, 0, len(membersRepo))
	for _, m := range membersRepo {
		members = append(members, &group.MemberInfo{
			Gid:        gid,
			Uid:        m.GroupMemberId,
			Nickname:   m.GroupMemberName,
			Role:       group.RoleType(m.GroupMemberType),
			MutedTime:  m.GroupMemberMuteTime,
			JoinedTime: m.GroupMemberJoinTime,
		})
	}
	return &group.GroupLimitedMembersResp{
		Members: members,
	}, nil
}
