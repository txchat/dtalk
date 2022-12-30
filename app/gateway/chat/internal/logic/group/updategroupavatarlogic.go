package group

import (
	"context"

	"github.com/txchat/dtalk/app/services/group/group"
	"github.com/txchat/dtalk/app/services/group/groupclient"
	xerror "github.com/txchat/dtalk/pkg/error"
	xhttp "github.com/txchat/dtalk/pkg/net/http"
	"github.com/txchat/dtalk/pkg/util"

	"github.com/txchat/dtalk/app/gateway/chat/internal/svc"
	"github.com/txchat/dtalk/app/gateway/chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	//xhttp "github.com/txchat/dtalk/pkg/net/http"
)

type UpdateGroupAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	custom *xhttp.Custom
}

func NewUpdateGroupAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupAvatarLogic {
	c, ok := xhttp.FromContext(ctx)
	if !ok {
		c = &xhttp.Custom{}
	}
	return &UpdateGroupAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		custom: c,
	}
}

func (l *UpdateGroupAvatarLogic) UpdateGroupAvatar(req *types.UpdateGroupAvatarReq) (resp *types.UpdateGroupAvatarResp, err error) {
	uid := l.custom.UID
	gid, err := util.ToInt64(req.IdStr)
	if err != nil {
		gid = req.Id
	}

	memInfoResp, err := l.svcCtx.GroupRPC.MemberInfo(l.ctx, &groupclient.MemberInfoReq{
		Gid: gid,
		Uid: []string{uid},
	})
	if err != nil {
		return nil, err
	}
	if len(memInfoResp.GetMembers()) < 1 {
		return nil, xerror.ErrGroupPersonNotExist
	}
	operator := memInfoResp.GetMembers()[0]
	if operator.GetRole() < group.RoleType_Manager {
		return nil, xerror.ErrGroupHigherPermission
	}

	_, err = l.svcCtx.GroupRPC.UpdateGroupAvatar(l.ctx, &groupclient.UpdateGroupAvatarReq{
		Gid:      gid,
		Operator: uid,
		Avatar:   req.Avatar,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.UpdateGroupAvatarResp{}
	return
}
