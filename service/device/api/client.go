package device

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
	client DeviceSrvClient
}

func New(etcdAddr, schema, srvName string, dial time.Duration) *Client {
	rb := naming.NewResolver(etcdAddr, schema)
	resolver.Register(rb)

	addr := fmt.Sprintf("%s:///%s", schema, srvName) // "schema://[authority]/service"
	fmt.Println("device rpc client call addr:", addr)

	conn, err := xgrpc.NewGRPCConnWithOpts(addr, dial)
	if err != nil {
		panic(err)
	}
	return &Client{
		client: NewDeviceSrvClient(conn),
	}
}

func (c *Client) AddDevice(ctx context.Context, in *Device) error {
	_, err := c.client.AddDevice(ctx, in)
	return err
}

func (c *Client) EnableThreadPushDevice(ctx context.Context, in *EnableThreadPushDeviceRequest) error {
	_, err := c.client.EnableThreadPushDevice(ctx, in)
	return err
}

func (c *Client) GetUserAllDevices(ctx context.Context, in *GetUserAllDevicesRequest) (*GetUserAllDevicesReply, error) {
	return c.client.GetUserAllDevices(ctx, in)
}
