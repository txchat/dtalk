package rpcnotify

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	xcall "github.com/txchat/dtalk/pkg/call"
	answer "github.com/txchat/dtalk/service/record/answer/api"
	xproto "github.com/txchat/imparse/proto"
)

type CallNotifyClient struct {
	answerClient *answer.Client
}

func NewCallNotifyClient(rpcCli *answer.Client) *CallNotifyClient {
	return &CallNotifyClient{
		answerClient: rpcCli,
	}
}

func (s *CallNotifyClient) SendStartCallSignal(ctx context.Context, target string, traceId int64) error {
	action := &xproto.SignalStartCall{
		TraceId: traceId,
	}
	body, err := proto.Marshal(action)
	if err != nil {
		return errors.WithMessagef(err, "proto.Marshal, action=%+v", action)
	}

	return s.answerClient.UniCastSignal(ctx, xproto.SignalType_StartCall, target, body)
}

func (s *CallNotifyClient) SendAcceptCallSignal(ctx context.Context, target string, traceId int64, ticket xcall.Ticket) error {
	action := &xproto.SignalAcceptCall{
		TraceId:       traceId,
		RoomId:        int32(ticket.RoomId),
		Uid:           target,
		UserSig:       ticket.UserSig,
		PrivateMapKey: ticket.PrivateMapKey,
		SkdAppId:      ticket.SDKAppID,
	}
	body, err := proto.Marshal(action)
	if err != nil {
		return errors.WithMessagef(err, "proto.Marshal, action=%+v", action)
	}

	return s.answerClient.UniCastSignal(ctx, xproto.SignalType_AcceptCall, target, body)
}

func (s *CallNotifyClient) SendStopCallSignal(ctx context.Context, Target string, TraceId int64, stopType xcall.StopType) error {
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

	return s.answerClient.UniCastSignal(ctx, xproto.SignalType_StopCall, Target, body)
}
