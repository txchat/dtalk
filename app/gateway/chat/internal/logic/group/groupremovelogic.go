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

type GroupRemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	custom *xhttp.Custom
}

func NewGroupRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupRemoveLogic {
	c, ok := xhttp.FromContext(ctx)
	if !ok {
		c = &xhttp.Custom{}
	}
	return &GroupRemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		custom: c,
	}
}

func (l *GroupRemoveLogic) GroupRemove(req *types.GroupRemoveReq) (resp *types.GroupRemoveResp, err error) {
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
	if operator.GetRole() < group.RoleType_Manager {
		return nil, xerror.ErrPermissionDenied
	}

	membersResp, err := l.svcCtx.GroupRPC.MembersInfo(l.ctx, &groupclient.MembersInfoReq{
		Gid: gid,
		Uid: req.MemberIds,
	})
	if err != nil {
		return nil, err
	}

	members := make([]string, 0)
	for _, member := range membersResp.GetMembers() {
		if member.GetRole() == group.RoleType_Owner || member.GetRole() == group.RoleType_Out {
			continue
		}
		members = append(members, member.GetUid())
	}

	kickOutResp, err := l.svcCtx.GroupRPC.KickOutMembers(l.ctx, &groupclient.KickOutMembersReq{
		Gid:      gid,
		Operator: uid,
		Mid:      members,
	})

	resp = &types.GroupRemoveResp{
		MemberNum: kickOutResp.GetNumber(),
		MemberIds: members,
	}
	return
}
