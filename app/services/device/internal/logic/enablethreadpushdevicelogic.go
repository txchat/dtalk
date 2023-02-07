package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/device/device"
	"github.com/txchat/dtalk/app/services/device/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type EnableThreadPushDeviceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEnableThreadPushDeviceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EnableThreadPushDeviceLogic {
	return &EnableThreadPushDeviceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EnableThreadPushDeviceLogic) EnableThreadPushDevice(in *device.EnableThreadPushDeviceRequest) (*device.Empty, error) {
	err := l.svcCtx.Repo.EnableDevice(in.GetUid(), in.GetConnId())
	if err != nil {
		l.Errorf("EnableDevice Repo err, err:%v", err)
		return nil, err
	}
	return &device.Empty{}, nil
}
