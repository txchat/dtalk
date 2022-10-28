package rpcnotify

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/txchat/dtalk/app/services/answer/answerclient"
	xcall "github.com/txchat/dtalk/pkg/call"
	"github.com/txchat/dtalk/pkg/call/sign"
	"github.com/txchat/imparse/proto/signal"
)

type CallNotifyClient struct {
	answerClient answerclient.Answer
}

func NewCallNotifyClient(rpcCli answerclient.Answer) *CallNotifyClient {
	return &CallNotifyClient{
		answerClient: rpcCli,
	}
}

func (s *CallNotifyClient) SendStartCallSignal(ctx context.Context, target string, traceId int64) error {
	action := &signal.SignalStartCall{
		TraceId: traceId,
	}
	body, err := proto.Marshal(action)
	if err != nil {
		return errors.WithMessagef(err, "proto.Marshal, action=%+v", action)
	}

	_, err = s.answerClient.UniCastSignal(ctx, &answerclient.UniCastSignalReq{
		Type:   signal.SignalType_StartCall,
		Target: target,
		Body:   body,
	})
	return err
}

func (s *CallNotifyClient) SendAcceptCallSignal(ctx context.Context, target string, traceId int64, ticket xcall.Ticket) error {
	cloudTicket, err := sign.FromBytes(ticket)
	if err != nil {
		return err
	}
	action := &signal.SignalAcceptCall{
		TraceId:       traceId,
		RoomId:        int32(cloudTicket.RoomId),
		Uid:           target,
		UserSig:       cloudTicket.UserSig,
		PrivateMapKey: cloudTicket.PrivateMapKey,
		SkdAppId:      cloudTicket.SDKAppID,
	}
	body, err := proto.Marshal(action)
	if err != nil {
		return errors.WithMessagef(err, "proto.Marshal, action=%+v", action)
	}

	_, err = s.answerClient.UniCastSignal(ctx, &answerclient.UniCastSignalReq{
		Type:   signal.SignalType_AcceptCall,
		Target: target,
		Body:   body,
	})
	return err
}

func (s *CallNotifyClient) SendStopCallSignal(ctx context.Context, Target string, TraceId int64, stopType xcall.StopType) error {
	var StopCallType signal.StopCallType
	switch signal.StopCallType(stopType) {
	case signal.StopCallType_Busy:
		StopCallType = signal.StopCallType_Busy
	case signal.StopCallType_Timeout:
		StopCallType = signal.StopCallType_Timeout
	case signal.StopCallType_Reject:
		StopCallType = signal.StopCallType_Reject
	case signal.StopCallType_Hangup:
		StopCallType = signal.StopCallType_Hangup
	case signal.StopCallType_Cancel:
		StopCallType = signal.StopCallType_Cancel

	}
	action := &signal.SignalStopCall{
		TraceId: TraceId,
		Reason:  StopCallType,
	}

	body, err := proto.Marshal(action)
	if err != nil {
		return errors.WithMessagef(err, "proto.Marshal, action=%+v", action)
	}

	_, err = s.answerClient.UniCastSignal(ctx, &answerclient.UniCastSignalReq{
		Type:   signal.SignalType_StopCall,
		Target: Target,
		Body:   body,
	})
	return err
}
