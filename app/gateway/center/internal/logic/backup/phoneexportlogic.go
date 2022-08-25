package backup

import (
	"context"

	"github.com/txchat/dtalk/app/gateway/center/internal/svc"
	"github.com/txchat/dtalk/app/gateway/center/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PhoneExportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPhoneExportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PhoneExportLogic {
	return &PhoneExportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PhoneExportLogic) PhoneExport(req *types.PhoneExportReq) (resp *types.PhoneExportResp, err error) {
	// todo: add your logic here and delete this line

	return
}
