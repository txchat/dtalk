package group

import (
	"context"

	"github.com/txchat/dtalk/app/gateway/chat/internal/svc"
	"github.com/txchat/dtalk/app/gateway/chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	//xhttp "github.com/txchat/dtalk/pkg/net/http"
)

type UpdateGroupMuteTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	//custom *xhttp.Custom
}

func NewUpdateGroupMuteTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupMuteTypeLogic {
	//c, ok := xhttp.FromContext(ctx)
	//if !ok {
	//    c = &xhttp.Custom{}
	//}
	return &UpdateGroupMuteTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		//custom: c,
	}
}

func (l *UpdateGroupMuteTypeLogic) UpdateGroupMuteType(req *types.UpdateGroupMuteTypeReq) (resp *types.UpdateGroupMuteTypeResp, err error) {
	// todo: add your logic here and delete this line

	return
}