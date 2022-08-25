package backup

import (
	"context"

	"github.com/txchat/dtalk/app/gateway/center/internal/svc"
	"github.com/txchat/dtalk/app/gateway/center/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PhoneBindingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPhoneBindingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PhoneBindingLogic {
	return &PhoneBindingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PhoneBindingLogic) PhoneBinding(req *types.PhoneBindingReq) (resp *types.PhoneBindingResp, err error) {
	// todo: add your logic here and delete this line

	return
}
