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

type GetMuteListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	custom *xhttp.Custom
}

func NewGetMuteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMuteListLogic {
	c, ok := xhttp.FromContext(ctx)
	if !ok {
		c = &xhttp.Custom{}
	}
	return &GetMuteListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		custom: c,
	}
}

func (l *GetMuteListLogic) GetMuteList(req *types.GetMuteListReq) (resp *types.GetMuteListResp, err error) {
	uid := l.custom.UID
	gid, err := util.ToInt64(req.IdStr)
	if err != nil {
		gid = req.Id
	}
	mutedListResp, err := l.svcCtx.GroupRPC.GetMuteList(l.ctx, &groupclient.GetMuteListReq{
		Gid:      gid,
		Operator: uid,
	})

	memberReply := make([]*types.GroupMember, 0, len(mutedListResp.GetMembers()))
	for _, m := range mutedListResp.GetMembers() {
		memberReply = append(memberReply, &types.GroupMember{
			MemberId:       m.GetUid(),
			MemberName:     m.GetNickname(),
			MemberType:     int32(m.GetRole()),
			MemberMuteTime: m.GetMutedTime(),
		})
	}
	resp = &types.GetMuteListResp{Members: memberReply}
	return
}
