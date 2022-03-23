package generator

import (
	"context"
	"fmt"
	"github.com/txchat/dtalk/pkg/naming"
	"github.com/txchat/dtalk/pkg/net/grpc"
	"google.golang.org/grpc/resolver"
	"time"
)

// AppID unique app id for service discovery
//const AppID = "identify.service.generator"

type Client struct {
	client GeneratorClient
}

func New(etcdAddr, schema, srvName string, dial time.Duration) *Client {
	rb := naming.NewResolver(etcdAddr, schema)
	resolver.Register(rb)

	addr := fmt.Sprintf("%s:///%s", schema, srvName) // "schema://[authority]/service"
	fmt.Println("generator rpc client call addr:", addr)

	conn, err := grpc.NewGRPCConn(addr, dial)
	if err != nil {
		panic(err)
	}
	return &Client{
		client: NewGeneratorClient(conn),
	}
}

func (c *Client) GetID() (int64, error) {
	reply, err := c.client.GetID(context.Background(), &Empty{})
	if reply == nil {
		return 0, err
	}
	return reply.Id, err
}

/*

type defaultGenerator struct {
	client idgen.GeneratorClient
}

// NewDefaultGenerator .
// etcdAddr like "127.0.0.1:2379;127.0.0.2:2379;"
func NewDefaultGenerator(etcdAddr, schema, srvName string, dial time.Duration) *defaultGenerator {
	rb := naming.NewResolver(etcdAddr, schema)
	resolver.Register(rb)

	addr := fmt.Sprintf("%s:///%s", schema, srvName) // "schema://[authority]/service"
	fmt.Println("generator rpc client call addr:", addr)

	conn, err := common.NewGRPCConn(addr, dial)
	if err != nil {
		panic(err)
	}
	return &defaultGenerator{client: idgen.NewGeneratorClient(conn)}
}

func (g *defaultGenerator) GetId() (int64, error) {
	var (
		req   idgen.Empty
		reply *idgen.GetIDReply
	)
	reply, err := g.client.GetID(context.Background(), &req)
	if err != nil {
		return 0, errors.WithMessagef(err, "getLogId")
	}

	return reply.Id, nil
}

*/
