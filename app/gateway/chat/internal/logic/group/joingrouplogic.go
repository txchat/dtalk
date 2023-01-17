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

type JoinGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	custom *xhttp.Custom
}

func NewJoinGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinGroupLogic {
	c, ok := xhttp.FromContext(ctx)
	if !ok {
		c = &xhttp.Custom{}
	}
	return &JoinGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		custom: c,
	}
}

func (l *JoinGroupLogic) JoinGroup(req *types.JoinGroupReq) (resp *types.JoinGroupResp, err error) {
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

	mInfo, err := l.svcCtx.GroupRPC.MemberInfo(l.ctx, &groupclient.MemberInfoReq{
		Gid: gid,
		Uid: req.InviterId,
	})
	if err != nil {
		if err == xerror.ErrGroupMemberNotExist {
			err = xerror.ErrPersonOutOfGroup
		}
		return nil, err
	}

	inviter := mInfo.GetMember()
	if inviter.GetRole() == group.RoleType_Out {
		return nil, xerror.ErrPersonOutOfGroup
	}

	switch gInfo.GetGroup().GetJoinType() {
	case group.GroupJoinType_AnybodyCanJoinGroup:
		_, err = l.svcCtx.GroupRPC.InviteMembers(l.ctx, &groupclient.InviteMembersReq{
			Gid:      gid,
			Operator: inviter.GetUid(),
			Mid:      []string{uid},
		})
		if err != nil {
			return nil, err
		}
	case group.GroupJoinType_JustManagerCanInvite:
		if inviter.GetRole() < group.RoleType_Manager || inviter.GetRole() == group.RoleType_Out {
			return nil, xerror.ErrGroupInvitePermissionDenied
		}
		_, err = l.svcCtx.GroupRPC.InviteMembers(l.ctx, &groupclient.InviteMembersReq{
			Gid:      gid,
			Operator: inviter.GetUid(),
			Mid:      []string{uid},
		})
		if err != nil {
			return nil, err
		}
	case group.GroupJoinType_NormalMemberCanInvite:
		return nil, xerror.ErrFeaturesUnSupported
	}
	resp = &types.JoinGroupResp{
		Id:    gid,
		IdStr: util.MustToString(gid),
	}
	return
}
