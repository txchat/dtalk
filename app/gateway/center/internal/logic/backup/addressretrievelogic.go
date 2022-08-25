package backup

import (
	"context"

	"github.com/txchat/dtalk/app/gateway/center/internal/svc"
	"github.com/txchat/dtalk/app/gateway/center/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddressRetrieveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddressRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddressRetrieveLogic {
	return &AddressRetrieveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddressRetrieveLogic) AddressRetrieve(req *types.AddressRetrieveReq) (resp *types.AddressRetrieveResp, err error) {
	// todo: add your logic here and delete this line

	return
}
