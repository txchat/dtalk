package grpc

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	xgrpc "github.com/txchat/dtalk/pkg/net/grpc"
	pb "github.com/txchat/dtalk/service/vip/api"
	"github.com/txchat/dtalk/service/vip/model"
	"github.com/txchat/dtalk/service/vip/service"
	"google.golang.org/grpc"
)

func New(c *xgrpc.ServerConfig, svr *service.Service, logger zerolog.Logger) *xgrpc.Server {
	connectionTimeout := grpc.ConnectionTimeout(time.Duration(c.Timeout))
	ws := xgrpc.NewServer(
		c,
		connectionTimeout,
	)
	pb.RegisterVIPSrvServer(ws.Server(), &server{pb.UnimplementedVIPSrvServer{}, svr})
	ws, err := ws.Start()
	if err != nil {
		panic(err)
	}
	return ws
}

type server struct {
	pb.UnimplementedVIPSrvServer
	svr *service.Service
}

func (s *server) AddVIPs(ctx context.Context, req *pb.AddVIPsReq) (*pb.AddVIPsReply, error) {
	unsuccessful := make([]string, 0)
	for _, uid := range req.GetUid() {
		if err := s.svr.AddVIP(ctx, uid); err != nil {
			unsuccessful = append(unsuccessful, uid)
		}
	}
	return &pb.AddVIPsReply{
		Uid: unsuccessful,
	}, nil
}

func (s *server) GetVIPs(ctx context.Context, req *pb.GetVIPsReq) (*pb.GetVIPsReply, error) {
	totalCount, err := s.svr.GetVIPCount(ctx)
	if err != nil {
		return &pb.GetVIPsReply{}, err
	}
	vipEntities, err := s.svr.GetScopeVIP(ctx, req.GetStart(), req.GetLimit())
	vips := make([]*pb.VIP, len(vipEntities))
	for i, entity := range vipEntities {
		vips[i] = convertVIP(entity)
	}
	return &pb.GetVIPsReply{
		Vip:        vips,
		TotalCount: totalCount,
	}, err
}

func (s *server) GetVIP(ctx context.Context, req *pb.GetVIPReq) (*pb.GetVIPReply, error) {
	v, err := s.svr.GetVIP(ctx, req.GetUid())
	return &pb.GetVIPReply{
		Vip: convertVIP(v),
	}, err
}

func convertVIP(v *model.VIPEntity) *pb.VIP {
	if v == nil {
		return nil
	}
	return &pb.VIP{
		Uid: v.UID,
	}
}
