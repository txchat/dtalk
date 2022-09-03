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

type StartCallLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	custom *xhttp.Custom
}

func NewStartCallLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartCallLogic {
	c, ok := xhttp.FromContext(ctx)
	if !ok {
		c = &xhttp.Custom{}
	}
	return &StartCallLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		custom: c,
	}
}

func (l *StartCallLogic) StartCall(req *types.StartCallReq) (resp *types.StartCallResp, err error) {
	//当req.GroupID 为空或者0时为私聊
	var groupId int64
	if req.GroupId != "" {
		groupId, err = util.ToInt64(req.GroupId)
		if err != nil {
			return nil, err
		}
	}

	if groupId != 0 && len(req.Invitees) != 1 {
		err = xerror.ErrFeaturesUnSupported
		return
	}

	var rpcResp *callclient.PrivateOfferResp
	rpcResp, err = l.svcCtx.CallRPC.PrivateOffer(l.ctx, &callclient.PrivateOfferReq{
		Operator: l.custom.UID,
		Invitee:  req.Invitees[0],
		RTCType:  call.RTCType(req.RTCType),
	})
	if err != nil {
		return
	}
	session := rpcResp.GetSession()
	resp = &types.StartCallResp{
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
