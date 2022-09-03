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

type HandleCallLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	custom *xhttp.Custom
}

func NewHandleCallLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandleCallLogic {
	c, ok := xhttp.FromContext(ctx)
	if !ok {
		c = &xhttp.Custom{}
	}
	return &HandleCallLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		custom: c,
	}
}

func (l *HandleCallLogic) HandleCall(req *types.HandleCallReq) (resp *types.HandleCallResp, err error) {
	traceId := req.TraceId
	if req.TraceIdStr != "" {
		traceId, err = util.ToInt64(req.TraceIdStr)
		if err != nil {
			err = xerror.ErrInvalidParams
			return
		}
	}

	switch req.Answer {
	case false:
		_, err = l.svcCtx.CallRPC.Reject(l.ctx, &callclient.RejectReq{
			Operator:   l.custom.UID,
			TraceId:    traceId,
			RejectType: call.RejectType_Reject,
		})
		if err != nil {
			return
		}
	case true:
		var rpcResp *callclient.AcceptResp
		rpcResp, err = l.svcCtx.CallRPC.Accept(l.ctx, &callclient.AcceptReq{
			Operator: l.custom.UID,
			TraceId:  traceId,
		})
		if err != nil {
			return
		}
		resp = &types.HandleCallResp{
			RoomId:        util.MustToInt32(rpcResp.GetRoomId()),
			UserSig:       rpcResp.GetUserSign(),
			PrivateMapKey: rpcResp.GetPrivateMapKey(),
			SDKAppId:      rpcResp.GetSDKAppID(),
		}
	}
	return
}
