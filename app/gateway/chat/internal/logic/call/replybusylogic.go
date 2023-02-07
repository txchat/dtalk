package call

import (
	"context"

	"github.com/txchat/dtalk/app/services/call/call"
	"github.com/txchat/dtalk/app/services/call/callclient"
	xerror "github.com/txchat/dtalk/pkg/error"
	xhttp "github.com/txchat/dtalk/pkg/net/http"
	"github.com/txchat/dtalk/pkg/util"

	"github.com/txchat/dtalk/app/gateway/chat/internal/svc"
	"github.com/txchat/dtalk/app/gateway/chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	//xhttp "github.com/txchat/dtalk/pkg/net/http"
)

type ReplyBusyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	custom *xhttp.Custom
}

func NewReplyBusyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReplyBusyLogic {
	c, ok := xhttp.FromContext(ctx)
	if !ok {
		c = &xhttp.Custom{}
	}
	return &ReplyBusyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		custom: c,
	}
}

func (l *ReplyBusyLogic) ReplyBusy(req *types.ReplyBusyReq) (resp *types.ReplyBusyResp, err error) {
	traceId := req.TraceId
	if req.TraceIdStr != "" {
		traceId, err = util.ToInt64(req.TraceIdStr)
		if err != nil {
			err = xerror.ErrInvalidParams
			return
		}
	}
	_, err = l.svcCtx.CallRPC.Reject(l.ctx, &callclient.RejectReq{
		Operator:   l.custom.UID,
		TraceId:    traceId,
		RejectType: call.RejectType_Occupied,
	})
	return
}
