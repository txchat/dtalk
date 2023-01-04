package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/group/internal/model"
	"github.com/txchat/dtalk/pkg/util"

	"github.com/txchat/dtalk/app/services/group/group"
	"github.com/txchat/dtalk/app/services/group/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGroupMemberNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateGroupMemberNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupMemberNameLogic {
	return &UpdateGroupMemberNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateGroupMemberNameLogic) UpdateGroupMemberName(in *group.UpdateGroupMemberNameReq) (*group.UpdateGroupMemberNameResp, error) {
	_, _, err := l.svcCtx.Repo.UpdateGroupMemberName(&model.GroupMember{
		GroupId:               in.GetGid(),
		GroupMemberId:         in.GetOperator(),
		GroupMemberName:       in.GetName(),
		GroupMemberUpdateTime: util.TimeNowUnixMilli(),
	})
	if err != nil {
		return nil, err
	}
	return &group.UpdateGroupMemberNameResp{}, nil
}
