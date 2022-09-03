package call

import (
	"context"
	"time"
)

type IDGenerator interface {
	GetID(ctx context.Context) (int64, error)
}

type SessionCreator struct {
	callingTimeout int64
	taskIDGen      IDGenerator
	roomIDGen      IDGenerator
}

func NewSessionCreator(callingTimeout int64, taskIDGen IDGenerator, roomIDGen IDGenerator) *SessionCreator {
	return &SessionCreator{
		callingTimeout: callingTimeout,
		taskIDGen:      taskIDGen,
		roomIDGen:      roomIDGen,
	}
}

func (sc *SessionCreator) InitSession(ctx context.Context, RTCType RTCType, caller string, invitees []string, groupID int64) (*Session, error) {
	//traceId
	traceID, err := sc.taskIDGen.GetID(ctx)
	if err != nil {
		return nil, err
	}
	//roomID
	roomID, err := sc.roomIDGen.GetID(ctx)
	if err != nil {
		return nil, err
	}

	session := &Session{
		RTCType:    RTCType,
		TraceId:    traceID,
		RoomId:     roomID,
		Timeout:    sc.callingTimeout * 1000,
		Deadline:   time.Now().Add(time.Duration(sc.callingTimeout) * time.Second).UnixMilli(),
		Status:     READY,
		Invitees:   invitees,
		Caller:     caller,
		CreateTime: time.Now().UnixMilli(),
		GroupId:    groupID,
	}
	return session, nil
}
