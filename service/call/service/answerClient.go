package service

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/txchat/dtalk/service/call/model"
	xproto "github.com/txchat/imparse/proto"
)

func (s *Service) noticeStartCall(target string, traceId int64) error {
	action := &xproto.SignalStartCall{
		TraceId: traceId,
	}
	body, err := proto.Marshal(action)
	if err != nil {
		return errors.WithMessagef(err, "proto.Marshal, action=%+v", action)
	}

	return s.answerClient.UniCastSignal(context.Background(), xproto.SignalType_StartCall, target, body)
}

func (s *Service) noticeAcceptCall(target string, traceId int64, roomId int32, userSig, privateMapKey string, sdkAppId int32) error {
	action := &xproto.SignalAcceptCall{
		TraceId:       traceId,
		RoomId:        roomId,
		Uid:           target,
		UserSig:       userSig,
		PrivateMapKey: privateMapKey,
		SkdAppId:      sdkAppId,
	}
	body, err := proto.Marshal(action)
	if err != nil {
		return errors.WithMessagef(err, "proto.Marshal, action=%+v", action)
	}

	return s.answerClient.UniCastSignal(context.Background(), xproto.SignalType_AcceptCall, target, body)
}

func (s *Service) noticeStopCall(Target string, TraceId int64, stopType model.StopType) error {
	var StopCallType xproto.StopCallType
	switch xproto.StopCallType(stopType) {
	case xproto.StopCallType_Busy:
		StopCallType = xproto.StopCallType_Busy
	case xproto.StopCallType_Timeout:
		StopCallType = xproto.StopCallType_Timeout
	case xproto.StopCallType_Reject:
		StopCallType = xproto.StopCallType_Reject
	case xproto.StopCallType_Hangup:
		StopCallType = xproto.StopCallType_Hangup
	case xproto.StopCallType_Cancel:
		StopCallType = xproto.StopCallType_Cancel

	}
	action := &xproto.SignalStopCall{
		TraceId: TraceId,
		Reason:  StopCallType,
	}

	body, err := proto.Marshal(action)
	if err != nil {
		return errors.WithMessagef(err, "proto.Marshal, action=%+v", action)
	}

	return s.answerClient.UniCastSignal(context.Background(), xproto.SignalType_StopCall, Target, body)
}
