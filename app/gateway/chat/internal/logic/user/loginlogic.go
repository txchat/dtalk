package user

import (
	"context"

	"github.com/txchat/dtalk/app/gateway/chat/internal/model"
	"github.com/txchat/dtalk/app/gateway/chat/internal/svc"
	"github.com/txchat/dtalk/app/gateway/chat/internal/types"
	"github.com/txchat/dtalk/app/services/device/deviceclient"
	xerror "github.com/txchat/dtalk/pkg/error"
	xhttp "github.com/txchat/dtalk/pkg/net/http"

	"github.com/txchat/imparse/proto/auth"
	"github.com/zeromicro/go-zero/core/logx"
	//xhttp "github.com/txchat/dtalk/pkg/net/http"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	custom *xhttp.Custom
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	c, ok := xhttp.FromContext(ctx)
	if !ok {
		c = &xhttp.Custom{}
	}
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		custom: c,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	uid := l.custom.UID
	deviceType := l.custom.Device
	uuid := l.custom.UUID

	if auth.Login_ConnType(req.ConnType) == auth.Login_Reconnect {
		dev, err := l.getConflictDevice(uid, deviceType, uuid)
		if err != nil {
			return nil, err
		}
		if dev != nil {
			return nil, xerror.NewCustomError(xerror.ErrReconnectRejected, model.LoginNotAllowedErr{
				Code: xerror.ErrReconnectRejected.Code(),
				Message: struct {
					Datetime   uint64 `json:"datetime"`
					Device     int32  `json:"device"`
					DeviceName string `json:"deviceName"`
					Uuid       string `json:"uuid"`
				}{
					Datetime:   dev.GetAddTime(),
					Device:     dev.GetDeviceType(),
					DeviceName: dev.GetDeviceName(),
					Uuid:       dev.GetDeviceUUid(),
				},
				Service: "",
			})
		}
	}
	resp = &types.LoginResp{
		Address: uid,
	}
	return
}

func (l *LoginLogic) getConflictDevice(uid, deviceType, uuid string) (*deviceclient.DeviceInfo, error) {
	switch deviceType {
	case auth.Device_Android.String(), auth.Device_IOS.String():
		//get device uuid
		devices, err := l.userAllDevices(uid)
		if err != nil {
			return nil, err
		}
		if len(devices) < 1 {
			break
		}
		var latestConn = &deviceclient.DeviceInfo{
			DeviceUUid: uuid,
		}
		for _, d := range devices {
			switch auth.Device(d.DeviceType) {
			case auth.Device_Android, auth.Device_IOS:
			default:
				continue
			}
			if latestConn.GetAddTime() < d.GetAddTime() {
				latestConn = d
			}
		}
		if latestConn.GetDeviceUUid() != uuid {
			return latestConn, nil
		}
	default:
	}
	return nil, nil
}

func (l *LoginLogic) userAllDevices(uid string) ([]*deviceclient.DeviceInfo, error) {
	resp, err := l.svcCtx.DeviceRPC.GetUserAllDevices(l.ctx, &deviceclient.GetUserAllDevicesRequest{
		Uid: uid,
	})
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, nil
	}

	return resp.GetDevices(), nil
}
