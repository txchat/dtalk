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

type UpdateGroupNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	custom *xhttp.Custom
}

func NewUpdateGroupNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupNameLogic {
	c, ok := xhttp.FromContext(ctx)
	if !ok {
		c = &xhttp.Custom{}
	}
	return &UpdateGroupNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		custom: c,
	}
}

func (l *UpdateGroupNameLogic) UpdateGroupName(req *types.UpdateGroupNameReq) (resp *types.UpdateGroupNameResp, err error) {
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

	_, err = l.svcCtx.GroupRPC.UpdateGroupName(l.ctx, &groupclient.UpdateGroupNameReq{
		Gid:      gid,
		Operator: uid,
		Name:     req.Name,
		MaskName: req.PublicName,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.UpdateGroupNameResp{}
	return
}
