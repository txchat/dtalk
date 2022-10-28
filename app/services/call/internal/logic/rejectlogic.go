package logic

import (
	"context"

	xerror "github.com/txchat/dtalk/pkg/error"

	"github.com/txchat/dtalk/app/services/call/call"
	"github.com/txchat/dtalk/app/services/call/internal/model"
	"github.com/txchat/dtalk/app/services/call/internal/svc"
	xcall "github.com/txchat/dtalk/pkg/call"
	"github.com/zeromicro/go-zero/core/logx"
)

type RejectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRejectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RejectLogic {
	return &RejectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RejectLogic) Reject(in *call.RejectReq) (*call.RejectResp, error) {
	s, err := l.svcCtx.Repo.GetSession(in.GetTraceId())
	if err != nil {
		return nil, err
	}
	session := xcall.Session(*s)

	if !session.IsPrivate() {
		return nil, xerror.ErrFeaturesUnSupported
	}
	target := session.Caller
	if session.Caller == in.GetOperator() {
		target = session.GetPrivateInvitee()
	}
	pt := xcall.NewPrivateTask(l.ctx, l.svcCtx.SignalNotify, in.GetOperator(), target)
	switch in.GetRejectType() {
	case call.RejectType_Reject:
		err = pt.Reject(&session)
		if err != nil {
			return nil, err
		}
	case call.RejectType_Occupied:
		err = pt.Occupied(&session)
		if err != nil {
			return nil, err
		}
	}
	err = l.svcCtx.Repo.SaveSession(model.Session(session))
	if err != nil {
		return nil, err
	}
	return &call.RejectResp{}, nil
}
