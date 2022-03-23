package grpc

import (
	"context"
	"time"

	xgrpc "github.com/txchat/dtalk/pkg/net/grpc"
	pb "github.com/txchat/dtalk/service/generator/api"
	"github.com/txchat/dtalk/service/generator/service"
	"google.golang.org/grpc"
)

func New(c *xgrpc.ServerConfig, svr *service.Service) *xgrpc.Server {
	connectionTimeout := grpc.ConnectionTimeout(time.Duration(c.Timeout))
	ws := xgrpc.NewServer(c, connectionTimeout)
	pb.RegisterGeneratorServer(ws.Server(), &server{pb.UnimplementedGeneratorServer{}, svr})
	ws, err := ws.Start()
	if err != nil {
		panic(err)
	}
	return ws
}

type server struct {
	pb.UnimplementedGeneratorServer
	svr *service.Service
}

func (s *server) GetID(ctx context.Context, req *pb.Empty) (*pb.GetIDReply, error) {
	return &pb.GetIDReply{
		Id: s.svr.GetID(),
	}, nil
}
