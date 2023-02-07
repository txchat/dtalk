package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/device/device"
	"github.com/txchat/dtalk/app/services/device/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetDeviceByConnIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDeviceByConnIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDeviceByConnIdLogic {
	return &GetDeviceByConnIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDeviceByConnIdLogic) GetDeviceByConnId(in *device.GetDeviceByConnIdRequest) (*device.DeviceInfo, error) {
	dev, err := l.svcCtx.Repo.GetDevicesByConnID(in.GetUid(), in.GetConnID())
	if err != nil {
		l.Errorf("GetDevicesByConnID Repo err, err:%v", err)
		return nil, err
	}
	return &device.DeviceInfo{
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
	}, nil
}
