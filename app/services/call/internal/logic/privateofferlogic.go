package logic

import (
	"context"

	"github.com/txchat/dtalk/pkg/util"

	"github.com/txchat/dtalk/app/services/call/call"
	"github.com/txchat/dtalk/app/services/call/internal/model"
	"github.com/txchat/dtalk/app/services/call/internal/svc"
	xcall "github.com/txchat/dtalk/pkg/call"
	"github.com/zeromicro/go-zero/core/logx"
)

type PrivateOfferLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPrivateOfferLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PrivateOfferLogic {
	return &PrivateOfferLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PrivateOfferLogic) PrivateOffer(in *call.PrivateOfferReq) (*call.PrivateOfferResp, error) {
	pt := xcall.NewPrivateTask(l.ctx, l.svcCtx.SessionCreator, l.svcCtx.SignalNotify, l.svcCtx.RTC, in.GetOperator(), in.GetInvitee(), xcall.RTCType(in.GetRTCType()))
	err := pt.Offer()
	if err != nil {
		return nil, err
	}
	session := pt.GetSession()
	err = l.svcCtx.Repo.SaveSession(model.Session(session))
	if err != nil {
		return nil, err
	}
	return &call.PrivateOfferResp{
		Session: &call.Session{
			TraceId:    session.TraceId,
			RoomId:     session.RoomId,
			RTCType:    util.MustToInt32(session.RTCType),
			Deadline:   session.Deadline,
			Status:     util.MustToInt32(session.Status),
			Invitees:   session.Invitees,
			Caller:     session.Caller,
			Timeout:    util.MustToInt32(session.Timeout),
			CreateTime: session.CreateTime,
			GroupId:    session.GroupId,
		},
	}, nil
}
