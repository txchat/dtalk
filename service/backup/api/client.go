package backup

import (
	"context"
	"fmt"
	"github.com/txchat/dtalk/pkg/naming"
	"github.com/txchat/dtalk/pkg/net/grpc"
	"google.golang.org/grpc/resolver"
	"time"
)

// AppID unique app id for service discovery
//const AppID = "identify.service.pusher"

type Client struct {
	client BackupClient
}

func New(etcdAddr, schema, srvName string, dial time.Duration) *Client {
	rb := naming.NewResolver(etcdAddr, schema)
	resolver.Register(rb)

	addr := fmt.Sprintf("%s:///%s", schema, srvName) // "schema://[authority]/service"
	fmt.Println("backup rpc client call addr:", addr)

	conn, err := grpc.NewGRPCConn(addr, dial)
	if err != nil {
		panic(err)
	}
	return &Client{
		client: NewBackupClient(conn),
	}
}

func (c *Client) Retrieve(ctx context.Context, tp QueryType, val string) (*RetrieveReply, error) {
	reply, err := c.client.Retrieve(ctx, &RetrieveReq{
		Type: tp,
		Val:  val,
	})
	return reply, err
}
