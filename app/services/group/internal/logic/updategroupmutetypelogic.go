package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/group/internal/model"
	"github.com/txchat/dtalk/pkg/util"

	"github.com/txchat/dtalk/app/services/group/group"
	"github.com/txchat/dtalk/app/services/group/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGroupMuteTypeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateGroupMuteTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupMuteTypeLogic {
	return &UpdateGroupMuteTypeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateGroupMuteTypeLogic) UpdateGroupMuteType(in *group.UpdateGroupMuteTypeReq) (*group.UpdateGroupMuteTypeResp, error) {
	_, _, err := l.svcCtx.Repo.UpdateGroupMuteType(&model.GroupInfo{
		GroupId:         in.GetGid(),
		GroupUpdateTime: util.TimeNowUnixMilli(),
		GroupMuteType:   int32(in.GetType()),
	})
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.SignalHub.UpdateGroupMuteType(l.ctx, in.GetGid(), int32(in.GetType()))
	err = l.svcCtx.NoticeHub.UpdateGroupMuteType(l.ctx, in.GetGid(), in.GetOperator(), int32(in.GetType()))
	return &group.UpdateGroupMuteTypeResp{}, nil
}
