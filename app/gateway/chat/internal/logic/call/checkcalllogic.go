package call

import (
	"context"

	"github.com/txchat/dtalk/app/services/call/callclient"
	xerror "github.com/txchat/dtalk/pkg/error"
	xhttp "github.com/txchat/dtalk/pkg/net/http"
	"github.com/txchat/dtalk/pkg/util"

	"github.com/txchat/dtalk/app/gateway/chat/internal/svc"
	"github.com/txchat/dtalk/app/gateway/chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	//xhttp "github.com/txchat/dtalk/pkg/net/http"
)

type CheckCallLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	custom *xhttp.Custom
}

func NewCheckCallLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckCallLogic {
	c, ok := xhttp.FromContext(ctx)
	if !ok {
		c = &xhttp.Custom{}
	}
	return &CheckCallLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		custom: c,
	}
}

func (l *CheckCallLogic) CheckCall(req *types.CheckCallReq) (resp *types.CheckCallResp, err error) {
	traceId := req.TraceId
	if req.TraceIdStr != "" {
		traceId, err = util.ToInt64(req.TraceIdStr)
		if err != nil {
			err = xerror.ErrInvalidParams
			return
		}
	}
	var rpcResp *callclient.CheckTaskResp
	rpcResp, err = l.svcCtx.CallRPC.CheckTask(l.ctx, &callclient.CheckTaskReq{
		Operator: l.custom.UID,
		TraceId:  traceId,
	})
	if err != nil {
		return
	}
	session := rpcResp.GetSession()
	resp = &types.CheckCallResp{
		TraceId:    session.GetTraceId(),
		TraceIdStr: util.MustToString(session.GetTraceId()),
		RTCType:    session.GetRTCType(),
		Invitees:   session.GetInvitees(),
		Caller:     session.GetCaller(),
		CreateTime: session.GetCreateTime(),
		Timeout:    session.GetTimeout(),
		Deadline:   session.GetDeadline(),
		GroupId:    util.MustToString(session.GetGroupId()),
	}
	return
}
