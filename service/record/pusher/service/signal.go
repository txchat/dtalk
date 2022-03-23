package service

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/txchat/dtalk/service/record/pusher/logH"
	xproto "github.com/txchat/imparse/proto"
)

//just socket push
func (s *Service) UniCastSignalReceived(ctx context.Context, item *logH.ConnSeqItem) error {
	actionProto := &xproto.SignalReceived{
		Logs: item.Logs,
	}
	actionData, err := proto.Marshal(actionProto)
	if err != nil {
		return err
	}
	err = s.answerClient.UniCastSignal(ctx, xproto.SignalType_Received, item.Sender, actionData)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UniCastSignalEndpointLogin(ctx context.Context, uid string, actionProto *xproto.SignalEndpointLogin) error {
	actionData, err := proto.Marshal(actionProto)
	if err != nil {
		return err
	}
	err = s.answerClient.UniCastSignal(ctx, xproto.SignalType_EndpointLogin, uid, actionData)
	if err != nil {
		return err
	}
	return nil
}
