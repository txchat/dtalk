package store

import (
	"context"
	"fmt"
	"time"

	"github.com/txchat/dtalk/pkg/naming"
	xgrpc "github.com/txchat/dtalk/pkg/net/grpc"
	"github.com/txchat/dtalk/service/record/store/model"
	"github.com/txchat/im-pkg/trace"
	xproto "github.com/txchat/imparse/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

type Client struct {
	client StoreClient
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
		client: NewStoreClient(conn),
	}
}

func (c *Client) DelRecord(ctx context.Context, tp xproto.Channel, mid int64) error {
	in := &DelRecordReq{
		Tp:  tp,
		Mid: mid,
	}
	_, err := c.client.DelRecord(ctx, in)
	return err
}

func (c *Client) GetRecord(ctx context.Context, tp xproto.Channel, mid int64) (*model.MsgContent, error) {
	in := &GetRecordReq{
		Tp:  tp,
		Mid: mid,
	}
	reply, err := c.client.GetRecord(ctx, in)
	if reply == nil {
		return nil, model.ErrRecordNotFind
	}
	return &model.MsgContent{
		Mid:        reply.Mid,
		Seq:        reply.Seq,
		SenderId:   reply.SenderId,
		ReceiverId: reply.ReceiverId,
		MsgType:    reply.MsgType,
		Content:    reply.Content,
		CreateTime: reply.CreateTime,
		Source:     reply.Source,
	}, err
}

func (c *Client) AddRecordFocus(ctx context.Context, uid string, mid int64, time uint64) (int32, error) {
	in := &AddRecordFocusReq{
		Uid:  uid,
		Mid:  mid,
		Time: time,
	}
	reply, err := c.client.AddRecordFocus(ctx, in)
	if reply == nil {
		return 0, model.ErrRecordNotFind
	}
	return reply.CurrentNum, err
}

func (c *Client) GetRecordsAfterMid(ctx context.Context, req *GetRecordsAfterMidReq) (*GetRecordsAfterMidReply, error) {
	reply, err := c.client.GetRecordsAfterMid(ctx, req)
	if reply == nil {
		return nil, model.ErrRecordNotFind
	}
	return reply, err
}

func (c *Client) GetSyncRecordsAfterMid(ctx context.Context, req *GetSyncRecordsAfterMidReq) (*GetSyncRecordsAfterMidReply, error) {
	reply, err := c.client.GetSyncRecordsAfterMid(ctx, req)
	if reply == nil {
		return nil, model.ErrRecordNotFind
	}
	return reply, err
}
