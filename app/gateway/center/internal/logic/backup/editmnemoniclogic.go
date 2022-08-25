package backup

import (
	"context"

	"github.com/txchat/dtalk/app/gateway/center/internal/svc"
	"github.com/txchat/dtalk/app/gateway/center/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditMnemonicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditMnemonicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditMnemonicLogic {
	return &EditMnemonicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditMnemonicLogic) EditMnemonic(req *types.EditMnemonicReq) (resp *types.EditMnemonicResp, err error) {
	// todo: add your logic here and delete this line

	return
}
