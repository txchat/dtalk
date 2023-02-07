package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/device/device"
	"github.com/txchat/dtalk/app/services/device/internal/model"
	"github.com/txchat/dtalk/app/services/device/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddDeviceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddDeviceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddDeviceLogic {
	return &AddDeviceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddDeviceLogic) AddDevice(in *device.DeviceInfo) (*device.Empty, error) {
	err := l.svcCtx.Repo.AddDeviceInfo(&model.Device{
		Uid:         in.Uid,
		ConnectId:   in.ConnectId,
		DeviceUuid:  in.DeviceUUid,
		DeviceType:  in.DeviceType,
		DeviceName:  in.DeviceName,
		Username:    in.Username,
		DeviceToken: in.DeviceToken,
		IsEnabled:   in.IsEnabled,
		AddTime:     in.AddTime,
		DTUid:       in.DTUid,
	})
	if err != nil {
		l.Errorf("AddDeviceInfo Repo err, err:%v", err)
		return nil, err
	}
	return &device.Empty{}, nil
}
