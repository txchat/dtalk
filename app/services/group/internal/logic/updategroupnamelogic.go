package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/group/internal/model"
	"github.com/txchat/dtalk/pkg/util"

	"github.com/txchat/dtalk/app/services/group/group"
	"github.com/txchat/dtalk/app/services/group/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGroupNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateGroupNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupNameLogic {
	return &UpdateGroupNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateGroupNameLogic) UpdateGroupName(in *group.UpdateGroupNameReq) (*group.UpdateGroupNameResp, error) {
	g, err := l.svcCtx.Repo.GetGroupById(in.GetGid())
	if err != nil {
		return nil, err
	}

	if in.GetName() != g.GroupName && in.GetName() != "" {
		g.GroupName = in.GetName()
	}

	if in.GetMaskName() != g.GroupPubName && in.GetMaskName() != "" {
		g.GroupPubName = in.GetMaskName()
	}

	_, _, err = l.svcCtx.Repo.UpdateGroupName(&model.GroupInfo{
		GroupId:         in.GetGid(),
		GroupUpdateTime: util.TimeNowUnixMilli(),
		GroupName:       g.GroupName,
		GroupPubName:    g.GroupPubName,
	})
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.SignalHub.UpdateGroupName(l.ctx, in.GetGid(), g.GroupName)
	err = l.svcCtx.NoticeHub.UpdateGroupName(l.ctx, in.GetGid(), in.GetOperator(), g.GroupName)

	return &group.UpdateGroupNameResp{}, nil
}
