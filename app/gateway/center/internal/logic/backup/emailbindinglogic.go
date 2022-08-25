package backup

import (
	"context"

	"github.com/txchat/dtalk/app/gateway/center/internal/svc"
	"github.com/txchat/dtalk/app/gateway/center/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmailBindingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailBindingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailBindingLogic {
	return &EmailBindingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmailBindingLogic) EmailBinding(req *types.EmailBindingReq) (resp *types.EmailBindingResp, err error) {
	// todo: add your logic here and delete this line

	return
}
