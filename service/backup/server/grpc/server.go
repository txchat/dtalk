package grpc

import (
	"context"
	"errors"
	"time"

	xgrpc "github.com/txchat/dtalk/pkg/net/grpc"
	pb "github.com/txchat/dtalk/service/backup/api"
	"github.com/txchat/dtalk/service/backup/model"
	"github.com/txchat/dtalk/service/backup/service"
	"google.golang.org/grpc"
)

func New(c *xgrpc.ServerConfig, svr *service.Service) *xgrpc.Server {
	connectionTimeout := grpc.ConnectionTimeout(time.Duration(c.Timeout))
	ws := xgrpc.NewServer(c, connectionTimeout)
	pb.RegisterBackupServer(ws.Server(), &server{pb.UnimplementedBackupServer{}, svr})
	ws, err := ws.Start()
	if err != nil {
		panic(err)
	}
	return ws
}

type server struct {
	pb.UnimplementedBackupServer
	svr *service.Service
}

func (s *server) Retrieve(ctx context.Context, req *pb.RetrieveReq) (*pb.RetrieveReply, error) {
	if req == nil {
		return &pb.RetrieveReply{}, errors.New("bad request")
	}
	var addrBackup *model.AddrBackup
	var err error
	switch req.Type {
	case pb.QueryType_Address:
		addrBackup, err = s.svr.AddressRetrieve(req.Val)
	case pb.QueryType_Phone:
		addrBackup, err = s.svr.AddressRetrieve(req.Val)
	case pb.QueryType_Email:
		addrBackup, err = s.svr.AddressRetrieve(req.Val)
	}
	if err != nil {
		return &pb.RetrieveReply{}, err
	}
	if addrBackup == nil {
		return nil, errors.New("record not found")
	}
	return toRetrieve(addrBackup), err
}

// toBindInfo Account, Bind类型转 BindInfo
func toRetrieve(src *model.AddrBackup) *pb.RetrieveReply {
	return &pb.RetrieveReply{
		Address:    src.Address,
		Area:       src.Area,
		Phone:      src.Phone,
		Email:      src.Email,
		Mnemonic:   src.Mnemonic,
		PrivateKey: src.PrivateKey,
		UpdateTime: src.UpdateTime.Unix(),
		CreateTime: src.UpdateTime.Unix(),
	}
}
