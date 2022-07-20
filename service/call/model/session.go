package model

import (
	"time"

	idgen "github.com/txchat/dtalk/service/generator/api"
)

type Session struct {
	TraceId int64
	RTCType int32
	RoomId  int32
	// 超出 Deadline 对方未接就结束通话
	Deadline int64
	// 0=对方未接通, 1=双方正在通话中, 2=通话结束
	Status     int32
	Invitees   []string
	Caller     string
	Timeout    int32
	CreateTime int64
	GroupId    int64
}

func NewSession(RTCType int32, caller string, invitees []string, groupId int64,
	idgenClient *idgen.Client, room *Room) (*Session, error) {
	traceId, err := idgenClient.GetID()
	if err != nil {
		return nil, err
	}
	roomId := room.GetID()
	session := &Session{
		RTCType:    RTCType,
		TraceId:    traceId,
		RoomId:     roomId,
		Timeout:    READYTIME * 1000,
		Deadline:   time.Now().Add(READYTIME*time.Second).UnixNano() / 1e6,
		Status:     READY,
		Invitees:   invitees,
		Caller:     caller,
		CreateTime: time.Now().UnixNano() / 1e6,
		GroupId:    groupId,
	}
	return session, nil
}
