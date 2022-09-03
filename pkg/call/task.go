package call

import (
	"context"
	"errors"

	"github.com/txchat/dtalk/pkg/call/sign"
	"github.com/txchat/dtalk/pkg/util"
)

var (
	ErrUserBusy = errors.New("user is busy")
)

type Ticket struct {
	RoomId        int64
	UserSig       string
	PrivateMapKey string
	SDKAppID      int32
}

type PrivateTask struct {
	ctx      context.Context
	sc       *SessionCreator
	notify   SignalNotify
	rtc      sign.TLSSig
	operator string
	target   string
	RTCType  RTCType

	session *Session
}

func NewPrivateTask(ctx context.Context, sessionCreator *SessionCreator, notify SignalNotify, rtc sign.TLSSig, operator, target string, rtcType RTCType) *PrivateTask {
	return &PrivateTask{
		ctx:      ctx,
		sc:       sessionCreator,
		notify:   notify,
		rtc:      rtc,
		operator: operator,
		target:   target,
		RTCType:  rtcType,
	}
}

func (t *PrivateTask) GetSession() Session {
	return *t.session
}

func (t *PrivateTask) SetSession(s *Session) {
	t.session = s
}

func (t *PrivateTask) Offer() error {
	//init session
	s, err := t.sc.InitSession(t.ctx, t.RTCType, t.operator, []string{t.target}, 0)
	if err != nil {
		return err
	}
	t.SetSession(s)
	//给B发送Start Call通知
	return t.notify.SendStartCallSignal(t.ctx, t.target, s.TraceId)
}

func (t *PrivateTask) Occupied() error {
	stopType := Busy
	err := t.notify.SendStopCallSignal(t.ctx, t.target, t.session.TraceId, stopType)
	if err != nil {
		return err
	}
	t.session.Finish()
	return nil
}

func (t *PrivateTask) Reject() error {
	stopType := Hangup

	if t.session.IsReady() {
		stopType = Reject
		if t.operator == t.session.Caller {
			// 发起方主动取消
			stopType = Cancel
		}
	}
	err := t.notify.SendStopCallSignal(t.ctx, t.target, t.session.TraceId, stopType)
	if err != nil {
		return err
	}
	t.session.Finish()
	return nil
}

func (t *PrivateTask) Accept() (Ticket, error) {
	// TODO 判断是否在被接收方组内

	if !t.session.IsReady() {
		return Ticket{}, ErrUserBusy
	}

	t.session.Processing()

	// 生成接收方的 userSig 和 privateMapKey
	inviteeTicket, err := t.GetTicket(t.operator, t.session.RoomId)
	if err != nil {
		return Ticket{}, err
	}
	// 生成发起方的 userSig 和 privateMapKey
	callerTicket, err := t.GetTicket(t.target, t.session.RoomId)
	if err != nil {
		return Ticket{}, err
	}
	return inviteeTicket, t.notify.SendAcceptCallSignal(t.ctx, t.session.Caller, t.session.TraceId, callerTicket)
}

func (t *PrivateTask) GetTicket(user string, roomId int64) (Ticket, error) {
	sdkAppId := t.rtc.GetAppId()

	// 生成接收方的 userSig 和 privateMapKey
	userSig, err := t.rtc.GetUserSig(t.operator)
	if err != nil {
		return Ticket{}, err
	}
	privateMapKey, err := t.rtc.GenPrivateMapKey(t.operator, util.MustToInt32(t.session.RoomId), 255)
	if err != nil {
		return Ticket{}, err
	}
	return Ticket{
		RoomId:        roomId,
		UserSig:       userSig,
		PrivateMapKey: privateMapKey,
		SDKAppID:      sdkAppId,
	}, nil
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
