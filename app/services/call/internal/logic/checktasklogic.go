package logic

import (
	"context"
	"time"

	"github.com/txchat/dtalk/app/services/call/call"
	"github.com/txchat/dtalk/app/services/call/internal/svc"
	xcall "github.com/txchat/dtalk/internal/call"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/zeromicro/go-zero/core/logx"
)

type CheckTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckTaskLogic {
	return &CheckTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckTaskLogic) CheckTask(in *call.CheckTaskReq) (*call.CheckTaskResp, error) {
	s, err := l.svcCtx.Repo.GetSession(in.GetTraceId())
	if err != nil {
		return nil, err
	}
	session := xcall.Session(*s)

	if !session.IsPrivate() {
		return nil, xerror.ErrFeaturesUnSupported
	}

	if !session.IsReady() {
		return nil, xerror.ErrCallTimeout
	}

	if session.IsTimeout(time.Now().UnixMilli()) {
		return nil, xerror.ErrCallTimeout
	}

	return &call.CheckTaskResp{
		Session: &call.Session{
			TraceId:    session.TaskID,
			RoomId:     session.RoomID,
			RTCType:    util.MustToInt32(session.RTCType),
			Deadline:   session.Deadline,
			Status:     util.MustToInt32(session.Status),
			Invitees:   session.Invitees,
			Caller:     session.Caller,
			Timeout:    util.MustToInt32(session.Timeout),
			CreateTime: session.CreateTime,
			GroupId:    session.GroupID,
		},
	}, nil
}
