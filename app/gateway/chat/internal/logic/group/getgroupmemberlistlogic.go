package group

import (
	"context"

	"github.com/txchat/dtalk/app/services/group/groupclient"
	"github.com/txchat/dtalk/pkg/util"

	"github.com/txchat/dtalk/app/gateway/chat/internal/svc"
	"github.com/txchat/dtalk/app/gateway/chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	//xhttp "github.com/txchat/dtalk/pkg/net/http"
)

type GetGroupMemberListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	//custom *xhttp.Custom
}

func NewGetGroupMemberListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupMemberListLogic {
	//c, ok := xhttp.FromContext(ctx)
	//if !ok {
	//    c = &xhttp.Custom{}
	//}
	return &GetGroupMemberListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		//custom: c,
	}
}

func (l *GetGroupMemberListLogic) GetGroupMemberList(req *types.GetGroupMemberListReq) (resp *types.GetGroupMemberListResp, err error) {
	gid, err := util.ToInt64(req.IdStr)
	if err != nil {
		gid = req.Id
	}
	membersResp, err := l.svcCtx.GroupRPC.GroupLimitedMembers(l.ctx, &groupclient.GroupLimitedMembersReq{
		Gid: gid,
	})
	if err != nil {
		return nil, err
	}

	members := make([]*types.GroupMember, 0, len(membersResp.GetMembers()))
	for _, m := range membersResp.GetMembers() {
		members = append(members, &types.GroupMember{
			MemberId:       m.GetUid(),
			MemberName:     m.GetNickname(),
			MemberType:     int32(m.GetRole()),
			MemberMuteTime: m.GetMutedTime(),
		})
	}
	resp = &types.GetGroupMemberListResp{
		Id:      gid,
		IdStr:   util.MustToString(gid),
		Members: members,
	}
	return
}
