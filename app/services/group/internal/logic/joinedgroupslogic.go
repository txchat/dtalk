package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/group/group"
	"github.com/txchat/dtalk/app/services/group/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type JoinedGroupsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJoinedGroupsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinedGroupsLogic {
	return &JoinedGroupsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JoinedGroupsLogic) JoinedGroups(in *group.JoinedGroupsReq) (*group.JoinedGroupsResp, error) {
	gid, err := l.svcCtx.Repo.JoinedGroups(in.GetUid())
	if err != nil {
		return nil, err
	}
	return &group.JoinedGroupsResp{
		Gid: gid,
	}, nil
}
