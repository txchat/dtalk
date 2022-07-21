package vip

import (
	"context"
	"fmt"
	"time"

	"github.com/txchat/dtalk/pkg/naming"
	xgrpc "github.com/txchat/dtalk/pkg/net/grpc"
	"google.golang.org/grpc/resolver"
)

// AppID unique app id for service discovery
//const AppID = "identify.service.pusher"

type Client struct {
	client VIPSrvClient
}

func New(etcdAddr, schema, srvName string, dial time.Duration) *Client {
	rb := naming.NewResolver(etcdAddr, schema)
	resolver.Register(rb)

	addr := fmt.Sprintf("%s:///%s", schema, srvName) // "schema://[authority]/service"
	fmt.Println("vip rpc client call addr:", addr)

	conn, err := xgrpc.NewGRPCConnWithOpts(addr, dial)
	if err != nil {
		panic(err)
	}
	return &Client{
		client: NewVIPSrvClient(conn),
	}
}

func (c *Client) AddVIPs(ctx context.Context, in *AddVIPsReq) (*AddVIPsReply, error) {
	return c.client.AddVIPs(ctx, in)
}

func (c *Client) GetVIPs(ctx context.Context, in *GetVIPsReq) (*GetVIPsReply, error) {
	return c.client.GetVIPs(ctx, in)
}

func (c *Client) GetVIP(ctx context.Context, in *GetVIPReq) (*GetVIPReply, error) {
	return c.client.GetVIP(ctx, in)
}
