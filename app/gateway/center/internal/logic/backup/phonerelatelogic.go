package backup

import (
	"context"

	"github.com/txchat/dtalk/app/gateway/center/internal/svc"
	"github.com/txchat/dtalk/app/gateway/center/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PhoneRelateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPhoneRelateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PhoneRelateLogic {
	return &PhoneRelateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PhoneRelateLogic) PhoneRelate(req *types.PhoneRelateReq) (resp *types.PhoneRelateResp, err error) {
	// todo: add your logic here and delete this line

	return
}
