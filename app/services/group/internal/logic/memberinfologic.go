package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/group/group"
	"github.com/txchat/dtalk/app/services/group/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberInfoLogic {
	return &MemberInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberInfoLogic) MemberInfo(in *group.MemberInfoReq) (*group.MemberInfoResp, error) {
	gid := in.GetGid()
	mid := in.GetUid()

	m, err := l.svcCtx.Repo.GetMemberById(gid, mid)
	if err != nil {
		return nil, err
	}
	return &group.MemberInfoResp{
		Member: &group.MemberInfo{
			Gid:        gid,
			Uid:        mid,
			Nickname:   m.GroupMemberName,
			Role:       group.RoleType(m.GroupMemberType),
			MutedTime:  m.GroupMemberMuteTime,
			JoinedTime: m.GroupMemberJoinTime,
		},
	}, nil
}
