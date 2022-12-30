package group

import (
	"context"

	"github.com/txchat/dtalk/app/services/group/groupclient"
	xerror "github.com/txchat/dtalk/pkg/error"
	xhttp "github.com/txchat/dtalk/pkg/net/http"
	"github.com/txchat/dtalk/pkg/util"

	"github.com/txchat/dtalk/app/gateway/chat/internal/svc"
	"github.com/txchat/dtalk/app/gateway/chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	//xhttp "github.com/txchat/dtalk/pkg/net/http"
)

type GetGroupMemberInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	custom *xhttp.Custom
}

func NewGetGroupMemberInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupMemberInfoLogic {
	c, ok := xhttp.FromContext(ctx)
	if !ok {
		c = &xhttp.Custom{}
	}
	return &GetGroupMemberInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		custom: c,
	}
}

func (l *GetGroupMemberInfoLogic) GetGroupMemberInfo(req *types.GetGroupMemberInfoReq) (resp *types.GetGroupMemberInfoResp, err error) {
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
		return nil, xerror.ErrGroupMemberNotExist
	}
	memInfo := memInfoResp.GetMembers()[0]

	resp = &types.GetGroupMemberInfoResp{
		GroupMember: types.GroupMember{
			MemberId:       memInfo.GetUid(),
			MemberName:     memInfo.GetNickname(),
			MemberType:     int32(memInfo.GetRole()),
			MemberMuteTime: memInfo.GetMutedTime(),
		},
	}
	return
}
