package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/group/internal/model"
	"github.com/txchat/dtalk/pkg/util"

	"github.com/txchat/dtalk/app/services/group/group"
	"github.com/txchat/dtalk/app/services/group/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGroupAvatarLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateGroupAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupAvatarLogic {
	return &UpdateGroupAvatarLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateGroupAvatarLogic) UpdateGroupAvatar(in *group.UpdateGroupAvatarReq) (*group.UpdateGroupAvatarResp, error) {
	_, _, err := l.svcCtx.Repo.UpdateGroupAvatar(&model.GroupInfo{
		GroupId:         in.GetGid(),
		GroupAvatar:     in.GetAvatar(),
		GroupUpdateTime: util.TimeNowUnixMilli(),
	})
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.SignalHub.UpdateGroupAvatar(l.ctx, in.GetGid(), in.GetOperator())
	return &group.UpdateGroupAvatarResp{}, nil
}
