package group

import (
	"context"

	"github.com/txchat/dtalk/app/gateway/chat/internal/svc"
	"github.com/txchat/dtalk/app/gateway/chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	//xhttp "github.com/txchat/dtalk/pkg/net/http"
)

type UpdateGroupNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	//custom *xhttp.Custom
}

func NewUpdateGroupNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupNameLogic {
	//c, ok := xhttp.FromContext(ctx)
	//if !ok {
	//    c = &xhttp.Custom{}
	//}
	return &UpdateGroupNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		//custom: c,
	}
}

func (l *UpdateGroupNameLogic) UpdateGroupName(req *types.UpdateGroupNameReq) (resp *types.UpdateGroupNameResp, err error) {
	// todo: add your logic here and delete this line

	return
}
