package call

import (
	"context"
	"errors"
)

var (
	ErrUserBusy = errors.New("user is busy")
)

type PrivateTask struct {
	ctx      context.Context
	notify   SignalNotify
	operator string
	target   string
}

func NewPrivateTask(ctx context.Context, notify SignalNotify, operator, target string) *PrivateTask {
	return &PrivateTask{
		ctx:      ctx,
		notify:   notify,
		operator: operator,
		target:   target,
	}
}

func (t *PrivateTask) Offer(sc *SessionCreator, rtcType RTCType) (*Session, error) {
	//init session
	session, err := sc.InitSession(t.ctx, rtcType, t.operator, []string{t.target}, 0)
	if err != nil {
		return nil, err
	}
	//给B发送Start Call通知
	err = t.notify.SendStartCallSignal(t.ctx, t.target, session.TaskID)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (t *PrivateTask) Occupied(session *Session) error {
	stopType := Busy
	err := t.notify.SendStopCallSignal(t.ctx, t.target, session.TaskID, stopType)
	if err != nil {
		return err
	}
	session.Finish()
	return nil
}

func (t *PrivateTask) Reject(session *Session) error {
	stopType := Hangup

	if session.IsReady() {
		stopType = Reject
		if t.operator == session.Caller {
			// 发起方主动取消
			stopType = Cancel
		}
	}
	err := t.notify.SendStopCallSignal(t.ctx, t.target, session.TaskID, stopType)
	if err != nil {
		return err
	}
	session.Finish()
	return nil
}

func (t *PrivateTask) Accept(createTicket TicketCreator, session *Session) (Ticket, error) {
	// TODO 判断是否在被接收方组内
	if !session.IsReady() {
		return nil, ErrUserBusy
	}

	session.Processing()

	// 生成接收方的 userSig 和 privateMapKey
	inviteeTicket, err := createTicket(t.operator, session.RoomID)
	if err != nil {
		return nil, err
	}
	// 生成发起方的 userSig 和 privateMapKey
	callerTicket, err := createTicket(t.target, session.RoomID)
	if err != nil {
		return nil, err
	}
	return inviteeTicket, t.notify.SendAcceptCallSignal(t.ctx, session.Caller, session.TaskID, callerTicket)
}

//
//type Group struct {
//	SessionCreator
//
//	initiator string
//	invitees []string
//	groupID int64
//}
//
//
//func (t *Group) Start() {
//	//init session
//	t.InitSession(t.initiator, t.invitees, t.groupID)
//
//	//发送Start Call通知
//
//}
//
//func (t *Group) Reject() {
//
//}
//
//func (t *Group) Accept() {
//
//}
