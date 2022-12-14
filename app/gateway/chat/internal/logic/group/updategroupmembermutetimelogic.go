package group

import (
	"context"

	"github.com/txchat/dtalk/app/gateway/chat/internal/svc"
	"github.com/txchat/dtalk/app/gateway/chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	//xhttp "github.com/txchat/dtalk/pkg/net/http"
)

type UpdateGroupMemberMuteTimeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	//custom *xhttp.Custom
}

func NewUpdateGroupMemberMuteTimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupMemberMuteTimeLogic {
	//c, ok := xhttp.FromContext(ctx)
	//if !ok {
	//    c = &xhttp.Custom{}
	//}
	return &UpdateGroupMemberMuteTimeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		//custom: c,
	}
}

func (l *UpdateGroupMemberMuteTimeLogic) UpdateGroupMemberMuteTime(req *types.UpdateGroupMemberMuteTimeReq) (resp *types.UpdateGroupMemberMuteTimeResp, err error) {
	// todo: add your logic here and delete this line

	return
}
