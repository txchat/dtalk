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

type MemberExitGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	custom *xhttp.Custom
}

func NewMemberExitGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberExitGroupLogic {
	c, ok := xhttp.FromContext(ctx)
	if !ok {
		c = &xhttp.Custom{}
	}
	return &MemberExitGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		custom: c,
	}
}

func (l *MemberExitGroupLogic) MemberExitGroup(req *types.MemberExitGroupReq) (resp *types.MemberExitGroupResp, err error) {
	uid := l.custom.UID
	gid, err := util.ToInt64(req.IdStr)
	if err != nil {
		gid = req.Id
	}

	operatorResp, err := l.svcCtx.GroupRPC.MemberInfo(l.ctx, &groupclient.MemberInfoReq{
		Gid: gid,
		Uid: uid,
	})
	if err != nil {
		return nil, err
	}
	operator := operatorResp.GetMember()
	if operator.GetRole() == group.RoleType_Owner || operator.GetRole() == group.RoleType_Out {
		return nil, xerror.ErrPermissionDenied
	}

	_, err = l.svcCtx.GroupRPC.MemberExit(l.ctx, &groupclient.MemberExitReq{
		Gid:      gid,
		Operator: uid,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.MemberExitGroupResp{}
	return
}
