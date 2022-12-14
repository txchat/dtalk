package group

import (
	"context"

	"github.com/txchat/dtalk/app/gateway/chat/internal/svc"
	"github.com/txchat/dtalk/app/gateway/chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	//xhttp "github.com/txchat/dtalk/pkg/net/http"
)

type UpdateGroupFriendTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	//custom *xhttp.Custom
}

func NewUpdateGroupFriendTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupFriendTypeLogic {
	//c, ok := xhttp.FromContext(ctx)
	//if !ok {
	//    c = &xhttp.Custom{}
	//}
	return &UpdateGroupFriendTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		//custom: c,
	}
}

func (l *UpdateGroupFriendTypeLogic) UpdateGroupFriendType(req *types.UpdateGroupFriendTypeReq) (resp *types.UpdateGroupFriendTypeResp, err error) {
	// todo: add your logic here and delete this line

	return
}
