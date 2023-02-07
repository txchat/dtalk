package logic

import (
	"context"
	"time"

	"github.com/txchat/dtalk/app/services/version/internal/svc"
	"github.com/txchat/dtalk/app/services/version/version"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReleaseSpecificVersionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReleaseSpecificVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReleaseSpecificVersionLogic {
	return &ReleaseSpecificVersionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ReleaseSpecificVersionLogic) ReleaseSpecificVersion(in *version.ReleaseSpecificVersionReq) (*version.ReleaseSpecificVersionResp, error) {
	err := l.svcCtx.Repo.ReleaseSpecificVersion(l.ctx, in.GetVid(), time.Now().UnixNano()/1e6, in.GetOperator())
	return &version.ReleaseSpecificVersionResp{}, err
}
