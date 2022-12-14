package group

import (
	"context"

	"github.com/txchat/dtalk/app/gateway/chat/internal/svc"
	"github.com/txchat/dtalk/app/gateway/chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	//xhttp "github.com/txchat/dtalk/pkg/net/http"
)

type GroupDisbandLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	//custom *xhttp.Custom
}

func NewGroupDisbandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupDisbandLogic {
	//c, ok := xhttp.FromContext(ctx)
	//if !ok {
	//    c = &xhttp.Custom{}
	//}
	return &GroupDisbandLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		//custom: c,
	}
}

func (l *GroupDisbandLogic) GroupDisband(req *types.GroupDisbandReq) (resp *types.GroupDisbandResp, err error) {
	// todo: add your logic here and delete this line

	return
}
