package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/group/internal/model"
	"github.com/txchat/dtalk/pkg/util"

	"github.com/txchat/dtalk/app/services/group/group"
	"github.com/txchat/dtalk/app/services/group/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGroupFriendlyTypeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateGroupFriendlyTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupFriendlyTypeLogic {
	return &UpdateGroupFriendlyTypeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateGroupFriendlyTypeLogic) UpdateGroupFriendlyType(in *group.UpdateGroupFriendlyTypeReq) (*group.UpdateGroupFriendlyTypeResp, error) {
	_, _, err := l.svcCtx.Repo.UpdateGroupFriendlyType(&model.GroupInfo{
		GroupId:         in.GetGid(),
		GroupUpdateTime: util.TimeNowUnixMilli(),
		GroupFriendType: int32(in.GetType()),
	})
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.SignalHub.UpdateGroupFriendlyType(l.ctx, in.GetGid(), int32(in.GetType()))
	return &group.UpdateGroupFriendlyTypeResp{}, nil
}
