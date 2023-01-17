package logic

import (
	"context"

	xgroup "github.com/txchat/dtalk/internal/group"
	xerror "github.com/txchat/dtalk/pkg/error"

	"github.com/txchat/dtalk/app/services/group/group"
	"github.com/txchat/dtalk/app/services/group/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckMemberInGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckMemberInGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckMemberInGroupLogic {
	return &CheckMemberInGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckMemberInGroupLogic) CheckMemberInGroup(in *group.CheckMemberInGroupReq) (*group.CheckMemberInGroupResp, error) {
	member, err := l.svcCtx.Repo.GetMemberById(in.GetGid(), in.GetMid())
	if err != nil {
		if err == xerror.ErrGroupMemberNotExist {
			return &group.CheckMemberInGroupResp{Ok: false}, nil
		}
		return nil, err
	}

	if xgroup.RoleType(member.GroupMemberType) == xgroup.Out {
		return &group.CheckMemberInGroupResp{Ok: false}, nil
	}
	return &group.CheckMemberInGroupResp{
		Ok: true,
	}, nil
}
