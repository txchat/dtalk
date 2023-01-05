package logic

import (
	"context"

	"github.com/txchat/dtalk/internal/call/sign"

	xerror "github.com/txchat/dtalk/pkg/error"

	"github.com/txchat/dtalk/app/services/call/call"
	"github.com/txchat/dtalk/app/services/call/internal/model"
	"github.com/txchat/dtalk/app/services/call/internal/svc"
	xcall "github.com/txchat/dtalk/internal/call"
	"github.com/zeromicro/go-zero/core/logx"
)

type AcceptLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAcceptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AcceptLogic {
	return &AcceptLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AcceptLogic) Accept(in *call.AcceptReq) (*call.AcceptResp, error) {
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
	pt := xcall.NewPrivateTask(l.ctx, l.svcCtx.SignalHub, in.GetOperator(), target)
	inviteeTicket, err := pt.Accept(l.svcCtx.TicketCreator, &session)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.Repo.SaveSession(model.Session(session))
	if err != nil {
		return nil, err
	}
	cloudTicket, err := sign.FromBytes(inviteeTicket)
	if err != nil {
		return nil, err
	}
	return &call.AcceptResp{
		RoomId:        cloudTicket.RoomId,
		UserSign:      cloudTicket.UserSig,
		PrivateMapKey: cloudTicket.PrivateMapKey,
		SDKAppID:      cloudTicket.SDKAppID,
	}, nil
}
