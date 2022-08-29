package backend

import (
	"context"

	"github.com/txchat/dtalk/app/services/version/versionclient"
	xhttp "github.com/txchat/dtalk/pkg/net/http"

	"github.com/txchat/dtalk/app/gateway/center/internal/svc"
	"github.com/txchat/dtalk/app/gateway/center/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeVersionStateLogic struct {
	logx.Logger
	ctx      context.Context
	svcCtx   *svc.ServiceContext
	username string
}

func NewChangeVersionStateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeVersionStateLogic {
	return &ChangeVersionStateLogic{
		Logger:   logx.WithContext(ctx),
		ctx:      ctx,
		svcCtx:   svcCtx,
		username: ctx.Value(xhttp.BackendJWTUsername).(string),
	}
}

func (l *ChangeVersionStateLogic) ChangeVersionState(req *types.ChangeVersionStateReq) (resp *types.ChangeVersionStateResp, err error) {
	_, err = l.svcCtx.VersionRPC.ReleaseSpecificVersion(l.ctx, &versionclient.ReleaseSpecificVersionReq{
		Vid:      req.Id,
		Operator: l.username,
	})
	return
}
