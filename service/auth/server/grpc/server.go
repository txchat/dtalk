package grpc

import (
	"context"
	"github.com/txchat/dtalk/service/auth/model"
	"time"

	xgrpc "github.com/txchat/dtalk/pkg/net/grpc"
	pb "github.com/txchat/dtalk/service/auth/api"
	"github.com/txchat/dtalk/service/auth/service"
	"google.golang.org/grpc"
)

func New(cfg *xgrpc.ServerConfig, svr *service.Service) *xgrpc.Server {
	connectionTimeout := grpc.ConnectionTimeout(time.Duration(cfg.Timeout))
	ws := xgrpc.NewServer(cfg, connectionTimeout)
	pb.RegisterAuthServer(ws.Server(), &server{svr: svr})
	ws, err := ws.Start()
	if err != nil {
		panic(err)
	}
	return ws
}

type server struct {
	pb.UnimplementedAuthServer
	svr *service.Service
}

func (s *server) DoAuth(ctx context.Context, in *pb.AuthRequest) (*pb.AuthReply, error) {
	req := &model.AuthRequest{
		AppId: in.Appid,
		Token: in.Token,
	}
	reply, err := s.svr.Auth(req)
	if err != nil {
		return &pb.AuthReply{}, err
	}
	return &pb.AuthReply{
		Uid: reply.Uid,
	}, nil
}
