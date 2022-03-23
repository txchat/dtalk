package grpc

import (
	"context"
	"time"

	xgrpc "github.com/txchat/dtalk/pkg/net/grpc"
	pb "github.com/txchat/dtalk/service/record/store/api"
	"github.com/txchat/dtalk/service/record/store/service"
	"github.com/txchat/im-pkg/trace"
	"google.golang.org/grpc"
)

func New(c *xgrpc.ServerConfig, svr *service.Service) *xgrpc.Server {
	connectionTimeout := grpc.ConnectionTimeout(time.Duration(c.Timeout))
	traceIcp := grpc.UnaryInterceptor(trace.OpentracingServerInterceptor)
	ws := xgrpc.NewServer(c, connectionTimeout, traceIcp)
	pb.RegisterStoreServer(ws.Server(), &server{pb.UnimplementedStoreServer{}, svr})
	ws, err := ws.Start()
	if err != nil {
		panic(err)
	}
	return ws
}

type server struct {
	pb.UnimplementedStoreServer
	svr *service.Service
}

func (s *server) GetRecord(ctx context.Context, req *pb.GetRecordReq) (*pb.GetRecordReply, error) {
	item, err := s.svr.GetSpecifyRecord(req.Tp, req.GetMid())
	if err != nil {
		return &pb.GetRecordReply{}, err
	}
	return &pb.GetRecordReply{
		Mid:        item.Mid,
		Seq:        item.Seq,
		SenderId:   item.SenderId,
		ReceiverId: item.ReceiverId,
		MsgType:    item.MsgType,
		Content:    item.Content,
		CreateTime: item.CreateTime,
		Source:     item.Source,
	}, nil
}

func (s *server) DelRecord(ctx context.Context, req *pb.DelRecordReq) (*pb.DelRecordReply, error) {
	return &pb.DelRecordReply{}, s.svr.DelRecord(req.Tp, req.GetMid())
}

func (s *server) AddRecordFocus(ctx context.Context, req *pb.AddRecordFocusReq) (*pb.AddRecordFocusReply, error) {
	err := s.svr.AddRecordFocus(req.GetUid(), req.GetMid(), req.GetTime())
	if err != nil {
		return &pb.AddRecordFocusReply{}, err
	}
	currentNum, err := s.svr.StatisticRecordFocusNumber(req.GetMid())
	if err != nil {
		return &pb.AddRecordFocusReply{}, err
	}
	return &pb.AddRecordFocusReply{
		CurrentNum: currentNum,
	}, nil
}

func (s *server) GetRecordsAfterMid(ctx context.Context, req *pb.GetRecordsAfterMidReq) (*pb.GetRecordsAfterMidReply, error) {
	items, err := s.svr.GetRecordsAfterMid(req.GetTp(), req.GetFrom(), req.GetTarget(), req.GetMid(), req.GetCount())
	if err != nil {
		return &pb.GetRecordsAfterMidReply{}, err
	}
	records := make([]*pb.GetRecordReply, len(items))
	for i, item := range items {
		records[i] = &pb.GetRecordReply{
			Mid:        item.Mid,
			Seq:        item.Seq,
			SenderId:   item.SenderId,
			ReceiverId: item.ReceiverId,
			MsgType:    item.MsgType,
			Content:    item.Content,
			CreateTime: item.CreateTime,
			Source:     item.Source,
		}
	}
	return &pb.GetRecordsAfterMidReply{
		Records: records,
	}, nil
}

func (s *server) GetSyncRecordsAfterMid(ctx context.Context, req *pb.GetSyncRecordsAfterMidReq) (*pb.GetSyncRecordsAfterMidReply, error) {
	items, err := s.svr.GetSyncMsg("", req.GetUid(), req.GetMid(), req.GetCount())
	if err != nil {
		return &pb.GetSyncRecordsAfterMidReply{}, err
	}
	return &pb.GetSyncRecordsAfterMidReply{
		Records: items,
	}, nil
}
