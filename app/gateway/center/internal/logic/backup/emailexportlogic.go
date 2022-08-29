package backup

import (
	"context"

	xerror "github.com/txchat/dtalk/pkg/error"

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
	// todo: 通过短信服务验证
	var b bool
	if !b {
		err = xerror.ErrExportAddressEmailInconsistent
	}
	return
}
