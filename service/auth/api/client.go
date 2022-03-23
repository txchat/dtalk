package auth

import (
	"context"
	"fmt"
	"github.com/txchat/dtalk/pkg/naming"
	"github.com/txchat/dtalk/pkg/net/grpc"
	"google.golang.org/grpc/resolver"
	"time"
)

type Client struct {
	client AuthClient
}

func New(etcdAddr, schema, srvName string, dial time.Duration) *Client {
	rb := naming.NewResolver(etcdAddr, schema)
	resolver.Register(rb)

	addr := fmt.Sprintf("%s:///%s", schema, srvName) // "schema://[authority]/service"
	fmt.Println("group rpc client call addr:", addr)

	conn, err := grpc.NewGRPCConn(addr, dial)
	if err != nil {
		panic(err)
	}
	return &Client{
		client: NewAuthClient(conn),
	}
}

func (c *Client) DoAuth(ctx context.Context, appId string, token string) (string, error) {
	request := &AuthRequest{
		Appid: appId,
		Token: token,
	}
	reply, err := c.client.DoAuth(ctx, request)
	if err != nil {
		return "", err
	}
	return reply.Uid, nil
}
