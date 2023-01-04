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

type UpdateGroupFriendTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	custom *xhttp.Custom
}

func NewUpdateGroupFriendTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupFriendTypeLogic {
	c, ok := xhttp.FromContext(ctx)
	if !ok {
		c = &xhttp.Custom{}
	}
	return &UpdateGroupFriendTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		custom: c,
	}
}

func (l *UpdateGroupFriendTypeLogic) UpdateGroupFriendType(req *types.UpdateGroupFriendTypeReq) (resp *types.UpdateGroupFriendTypeResp, err error) {
	uid := l.custom.UID
	gid, err := util.ToInt64(req.IdStr)
	if err != nil {
		gid = req.Id
	}

	memInfoResp, err := l.svcCtx.GroupRPC.MemberInfo(l.ctx, &groupclient.MemberInfoReq{
		Gid: gid,
		Uid: uid,
	})
	if err != nil {
		return nil, err
	}
	operator := memInfoResp.GetMember()
	if operator.GetRole() < group.RoleType_Manager {
		return nil, xerror.ErrGroupHigherPermission
	}

	_, err = l.svcCtx.GroupRPC.UpdateGroupFriendlyType(l.ctx, &groupclient.UpdateGroupFriendlyTypeReq{
		Gid:      gid,
		Operator: uid,
		Type:     group.GroupFriendlyType(req.FriendType),
	})
	if err != nil {
		return nil, err
	}

	resp = &types.UpdateGroupFriendTypeResp{}
	return
}
