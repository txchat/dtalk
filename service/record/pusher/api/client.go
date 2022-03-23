package pusher

import (
	"context"
	"fmt"
	"time"

	"github.com/txchat/dtalk/pkg/naming"
	xgrpc "github.com/txchat/dtalk/pkg/net/grpc"
	"github.com/txchat/im-pkg/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

// AppID unique app id for service discovery
//const AppID = "identify.service.pusher"

type Client struct {
	client PusherClient
}

func New(etcdAddr, schema, srvName string, dial time.Duration) *Client {
	rb := naming.NewResolver(etcdAddr, schema)
	resolver.Register(rb)

	addr := fmt.Sprintf("%s:///%s", schema, srvName) // "schema://[authority]/service"
	fmt.Println("pusher rpc client call addr:", addr)

	conn, err := xgrpc.NewGRPCConnWithOpts(addr, dial, grpc.WithUnaryInterceptor(trace.OpentracingClientInterceptor))
	if err != nil {
		panic(err)
	}
	return &Client{
		client: NewPusherClient(conn),
	}
}

func (c *Client) PushClient(ctx context.Context, key, from string, data []byte) error {
	_, err := c.client.PushClient(ctx, &PushReq{
		Key:  key,
		From: from,
		Data: data,
	})
	return err
}
