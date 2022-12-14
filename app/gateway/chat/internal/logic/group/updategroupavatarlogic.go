package group

import (
	"context"

	"github.com/txchat/dtalk/app/gateway/chat/internal/svc"
	"github.com/txchat/dtalk/app/gateway/chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	//xhttp "github.com/txchat/dtalk/pkg/net/http"
)

type UpdateGroupAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	//custom *xhttp.Custom
}

func NewUpdateGroupAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupAvatarLogic {
	//c, ok := xhttp.FromContext(ctx)
	//if !ok {
	//    c = &xhttp.Custom{}
	//}
	return &UpdateGroupAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		//custom: c,
	}
}

func (l *UpdateGroupAvatarLogic) UpdateGroupAvatar(req *types.UpdateGroupAvatarReq) (resp *types.UpdateGroupAvatarResp, err error) {
	// todo: add your logic here and delete this line

	return
}
