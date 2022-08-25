package backup

import (
	"context"

	"github.com/txchat/dtalk/app/gateway/center/internal/svc"
	"github.com/txchat/dtalk/app/gateway/center/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmailExportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailExportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailExportLogic {
	return &EmailExportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmailExportLogic) EmailExport(req *types.EmailExportReq) (resp *types.EmailExportResp, err error) {
	// todo: add your logic here and delete this line

	return
}
