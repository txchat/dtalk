package grpc

import (
	"context"
	"time"

	"github.com/txchat/dtalk/service/device/model"

	"github.com/rs/zerolog"
	xgrpc "github.com/txchat/dtalk/pkg/net/grpc"
	pb "github.com/txchat/dtalk/service/device/api"
	"github.com/txchat/dtalk/service/device/service"
	"google.golang.org/grpc"
)

func New(c *xgrpc.ServerConfig, svr *service.Service, logger zerolog.Logger) *xgrpc.Server {
	connectionTimeout := grpc.ConnectionTimeout(time.Duration(c.Timeout))
	ws := xgrpc.NewServer(
		c,
		connectionTimeout,
	)
	pb.RegisterDeviceSrvServer(ws.Server(), &server{pb.UnimplementedDeviceSrvServer{}, svr})
	ws, err := ws.Start()
	if err != nil {
		panic(err)
	}
	return ws
}

type server struct {
	pb.UnimplementedDeviceSrvServer
	svr *service.Service
}

func (s *server) AddDevice(ctx context.Context, req *pb.Device) (*pb.Empty, error) {
	return &pb.Empty{}, s.svr.AddDevice(ctx, &model.Device{
		Uid:         req.Uid,
		ConnectId:   req.ConnectId,
		DeviceUuid:  req.DeviceUUid,
		DeviceType:  req.DeviceType,
		DeviceName:  req.DeviceName,
		Username:    req.Username,
		DeviceToken: req.DeviceToken,
		IsEnabled:   req.IsEnabled,
		AddTime:     req.AddTime,
		DTUid:       req.DTUid,
	})
}

func (s *server) EnableThreadPushDevice(ctx context.Context, req *pb.EnableThreadPushDeviceRequest) (*pb.Empty, error) {
	return &pb.Empty{}, s.svr.EnableThreadPushDevice(ctx, req.GetUid(), req.GetConnId())
}

func (s *server) GetUserAllDevices(ctx context.Context, req *pb.GetUserAllDevicesRequest) (*pb.GetUserAllDevicesReply, error) {
	devices, err := s.svr.GetUserAllDevices(ctx, req.GetUid())
	if err != nil {
		return &pb.GetUserAllDevicesReply{}, err
	}
	devicesRlt := make([]*pb.Device, len(devices))
	for i, device := range devices {
		devicesRlt[i] = &pb.Device{
			Uid:         device.Uid,
			ConnectId:   device.ConnectId,
			DeviceUUid:  device.DeviceUuid,
			DeviceType:  device.DeviceType,
			DeviceName:  device.DeviceName,
			Username:    device.Username,
			DeviceToken: device.DeviceToken,
			IsEnabled:   device.IsEnabled,
			AddTime:     device.AddTime,
			DTUid:       device.DTUid,
		}
	}
	return &pb.GetUserAllDevicesReply{
		Devices: devicesRlt,
	}, nil
}

func (s *server) GetDeviceByConnId(ctx context.Context, req *pb.GetDeviceByConnIdRequest) (*pb.Device, error) {
	device, err := s.svr.GetDeviceByConnId(ctx, req.GetUid(), req.GetConnID())
	if err != nil {
		return &pb.Device{}, err
	}
	return &pb.Device{
		Uid:         device.Uid,
		ConnectId:   device.ConnectId,
		DeviceUUid:  device.DeviceUuid,
		DeviceType:  device.DeviceType,
		DeviceName:  device.DeviceName,
		Username:    device.Username,
		DeviceToken: device.DeviceToken,
		IsEnabled:   device.IsEnabled,
		AddTime:     device.AddTime,
		DTUid:       device.DTUid,
	}, nil
}
