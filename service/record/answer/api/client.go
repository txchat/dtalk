package answer

import (
	"context"
	"fmt"
	"time"

	"github.com/txchat/dtalk/pkg/naming"
	xgrpc "github.com/txchat/dtalk/pkg/net/grpc"
	"github.com/txchat/im-pkg/trace"
	xproto "github.com/txchat/imparse/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

type Client struct {
	client AnswerClient
}

func New(etcdAddr, schema, srvName string, dial time.Duration) *Client {
	rb := naming.NewResolver(etcdAddr, schema)
	resolver.Register(rb)

	addr := fmt.Sprintf("%s:///%s", schema, srvName) // "schema://[authority]/service"
	fmt.Println("answer rpc client call addr:", addr)

	conn, err := xgrpc.NewGRPCConnWithOpts(addr, dial, grpc.WithUnaryInterceptor(trace.OpentracingClientInterceptor))
	if err != nil {
		panic(err)
	}
	return &Client{
		client: NewAnswerClient(conn),
	}
}

func (c *Client) PushCommonMsg(ctx context.Context, key, From string, body []byte) (int64, uint64, error) {
	in := &PushCommonMsgReq{
		Key:  key,
		From: From,
		Body: body,
	}
	res, err := c.client.PushCommonMsg(ctx, in)
	if err != nil {
		return 0, 0, err
	}
	return res.Mid, res.Time, err
}

func (c *Client) PushNoticeMsg(ctx context.Context, Seq string, ChannelType int32, From, Target string, Data []byte) (int64, error) {
	in := &PushNoticeMsgReq{
		Seq:         Seq,
		ChannelType: ChannelType,
		From:        From,
		Target:      Target,
		Data:        Data,
	}
	res, err := c.client.PushNoticeMsg(ctx, in)
	if err != nil {
		return 0, err
	}
	return res.Mid, err
}

func (c *Client) UniCastSignal(ctx context.Context, Action xproto.SignalType, Target string, Body []byte) error {
	in := &UniCastSignalReq{
		Type:   Action,
		Target: Target,
		Body:   Body,
	}
	_, err := c.client.UniCastSignal(ctx, in)
	return err
}

func (c *Client) GroupCastSignal(ctx context.Context, Action xproto.SignalType, Target string, Body []byte) error {
	in := &GroupCastSignalReq{
		Type:   Action,
		Target: Target,
		Body:   Body,
	}
	_, err := c.client.GroupCastSignal(ctx, in)
	return err
}
