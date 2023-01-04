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

type ChangeOwnerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	custom *xhttp.Custom
}

func NewChangeOwnerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeOwnerLogic {
	c, ok := xhttp.FromContext(ctx)
	if !ok {
		c = &xhttp.Custom{}
	}
	return &ChangeOwnerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		custom: c,
	}
}

func (l *ChangeOwnerLogic) ChangeOwner(req *types.ChangeOwnerReq) (resp *types.ChangeOwnerResp, err error) {
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
	if operator.GetRole() != group.RoleType_Owner {
		return nil, xerror.ErrPermissionDenied
	}

	_, err = l.svcCtx.GroupRPC.ChangeOwner(l.ctx, &groupclient.ChangeOwnerReq{
		Gid:      gid,
		Operator: uid,
		New:      req.MemberId,
	})
	if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.GroupRPC.UnMuteMembers(l.ctx, &groupclient.UnMuteMembersReq{
		Gid:      gid,
		Operator: req.MemberId,
		Mid:      []string{req.MemberId},
	})
	if err != nil {
		return nil, err
	}

	resp = &types.ChangeOwnerResp{}
	return
}
