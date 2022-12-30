package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/group/internal/model"
	"github.com/txchat/dtalk/pkg/util"

	"github.com/txchat/dtalk/app/services/group/group"
	"github.com/txchat/dtalk/app/services/group/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGroupJoinTypeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateGroupJoinTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupJoinTypeLogic {
	return &UpdateGroupJoinTypeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateGroupJoinTypeLogic) UpdateGroupJoinType(in *group.UpdateGroupJoinTypeReq) (*group.UpdateGroupJoinTypeResp, error) {
	_, _, err := l.svcCtx.Repo.UpdateGroupJoinType(&model.GroupInfo{
		GroupId:         in.GetGid(),
		GroupUpdateTime: util.TimeNowUnixMilli(),
		GroupJoinType:   int32(in.GetType()),
	})
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.SignalHub.UpdateGroupJoinType(l.ctx, in.GetGid(), int32(in.GetType()))
	return &group.UpdateGroupJoinTypeResp{}, nil
}
