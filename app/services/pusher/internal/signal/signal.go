package signal

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/txchat/dtalk/app/services/answer/answer"
	"github.com/txchat/dtalk/app/services/answer/answerclient"
	"github.com/txchat/dtalk/app/services/pusher/internal/recordhelper"
	"github.com/txchat/imparse/proto/signal"
)

type Signal struct {
	conn answerclient.Answer
}

func NewSignal(conn answerclient.Answer) *Signal {
	return &Signal{conn: conn}
}

func (s *Signal) UniCastReceived(ctx context.Context, item *recordhelper.ConnSeqItem) error {
	actionProto := &signal.SignalReceived{
		Logs: item.Logs,
	}
	actionData, err := proto.Marshal(actionProto)
	if err != nil {
		return err
	}
	_, err = s.conn.UniCastSignal(ctx, &answerclient.UniCastSignalReq{
		Type:   answer.SignalType_Received,
		Target: item.Sender,
		Body:   actionData,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *Signal) UniCastEndpointLogin(ctx context.Context, uid string, actionProto *signal.SignalEndpointLogin) error {
	actionData, err := proto.Marshal(actionProto)
	if err != nil {
		return err
	}
	_, err = s.conn.UniCastSignal(ctx, &answerclient.UniCastSignalReq{
		Type:   answer.SignalType_EndpointLogin,
		Target: uid,
		Body:   actionData,
	})
	if err != nil {
		return err
	}
	return nil
}
