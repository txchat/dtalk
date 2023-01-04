package group

import (
	"context"

	"github.com/txchat/dtalk/app/services/group/groupclient"
	xhttp "github.com/txchat/dtalk/pkg/net/http"
	"github.com/txchat/dtalk/pkg/util"

	"github.com/txchat/dtalk/app/gateway/chat/internal/svc"
	"github.com/txchat/dtalk/app/gateway/chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	//xhttp "github.com/txchat/dtalk/pkg/net/http"
)

type UpdateGroupMemberNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	custom *xhttp.Custom
}

func NewUpdateGroupMemberNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupMemberNameLogic {
	c, ok := xhttp.FromContext(ctx)
	if !ok {
		c = &xhttp.Custom{}
	}
	return &UpdateGroupMemberNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		custom: c,
	}
}

func (l *UpdateGroupMemberNameLogic) UpdateGroupMemberName(req *types.UpdateGroupMemberNameReq) (resp *types.UpdateGroupMemberNameResp, err error) {
	uid := l.custom.UID
	gid, err := util.ToInt64(req.IdStr)
	if err != nil {
		gid = req.Id
	}

	_, err = l.svcCtx.GroupRPC.UpdateGroupMemberName(l.ctx, &groupclient.UpdateGroupMemberNameReq{
		Gid:      gid,
		Operator: uid,
		Name:     req.MemberName,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.UpdateGroupMemberNameResp{}
	return
}
