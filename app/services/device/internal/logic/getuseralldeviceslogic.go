package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/device/device"
	"github.com/txchat/dtalk/app/services/device/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAllDevicesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAllDevicesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAllDevicesLogic {
	return &GetUserAllDevicesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserAllDevicesLogic) GetUserAllDevices(in *device.GetUserAllDevicesRequest) (*device.GetUserAllDevicesReply, error) {
	devs, err := l.svcCtx.Repo.GetAllDevices(in.GetUid())
	if err != nil {
		l.Errorf("GetDevicesByConnID Repo err, err:%v", err)
		return nil, err
	}
	devices := make([]*device.DeviceInfo, len(devs))
	for i, dev := range devs {
		devices[i] = &device.DeviceInfo{
			Uid:         dev.Uid,
			ConnectId:   dev.ConnectId,
			DeviceUUid:  dev.DeviceUuid,
			DeviceType:  dev.DeviceType,
			DeviceName:  dev.DeviceName,
			Username:    dev.Username,
			DeviceToken: dev.DeviceToken,
			IsEnabled:   dev.IsEnabled,
			AddTime:     dev.AddTime,
			DTUid:       dev.DTUid,
		}
	}
	return &device.GetUserAllDevicesReply{
		Devices: devices,
	}, nil
}
