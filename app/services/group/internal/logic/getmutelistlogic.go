package logic

import (
	"context"

	"github.com/txchat/dtalk/pkg/util"

	"github.com/txchat/dtalk/app/services/group/group"
	"github.com/txchat/dtalk/app/services/group/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMuteListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMuteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMuteListLogic {
	return &GetMuteListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMuteListLogic) GetMuteList(in *group.GetMuteListReq) (*group.GetMuteListResp, error) {
	members, err := l.svcCtx.Repo.GetMutedMembers(in.GetGid(), util.TimeNowUnixMilli())
	if err != nil {
		return nil, err
	}
	memberReply := make([]*group.MemberInfo, 0, len(members))
	for _, member := range members {
		memberReply = append(memberReply, &group.MemberInfo{
			Gid:        member.GroupId,
			Uid:        member.GroupMemberId,
			Nickname:   member.GroupMemberName,
			Role:       group.RoleType(member.GroupMemberType),
			MutedTime:  member.GroupMemberMuteTime,
			JoinedTime: member.GroupMemberJoinTime,
		})
	}
	return &group.GetMuteListResp{
		Members: memberReply,
	}, nil
}
