package group

import (
	"context"

	"github.com/txchat/dtalk/app/gateway/chat/internal/model"
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

type SetAdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	custom *xhttp.Custom
}

func NewSetAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetAdminLogic {
	c, ok := xhttp.FromContext(ctx)
	if !ok {
		c = &xhttp.Custom{}
	}
	return &SetAdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		custom: c,
	}
}

func (l *SetAdminLogic) SetAdmin(req *types.SetAdminReq) (resp *types.SetAdminResp, err error) {
	uid := l.custom.UID
	gid, err := util.ToInt64(req.IdStr)
	if err != nil {
		gid = req.Id
	}

	gInfo, err := l.svcCtx.GroupRPC.GroupInfo(l.ctx, &groupclient.GroupInfoReq{
		Gid: gid,
	})
	if err != nil {
		return nil, err
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
		return nil, xerror.ErrGroupOwnerSetAdmin
	}

	if gInfo.GetGroup().GetManagerNumbers() >= model.GroupManagerLimit {
		return nil, xerror.ErrGroupAdminNumLimit
	}

	_, err = l.svcCtx.GroupRPC.ChangeMemberRole(l.ctx, &groupclient.ChangeMemberRoleReq{
		Gid:      gid,
		Operator: uid,
		Mid:      req.MemberId,
		Role:     group.RoleType(req.MemberType),
	})
	if err != nil {
		return nil, err
	}

	// 解除新管理员禁言
	if group.RoleType(req.MemberType) == group.RoleType_Manager {
		_, err = l.svcCtx.GroupRPC.UnMuteMembers(l.ctx, &groupclient.UnMuteMembersReq{
			Gid:      gid,
			Operator: uid,
			Mid:      []string{req.MemberId},
		})
		if err != nil {
			return nil, err
		}
	}

	return
}
