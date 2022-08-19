package dao

import "github.com/txchat/dtalk/app/services/device/internal/model"

type DeviceRepository interface {
	AddDeviceInfo(device *model.Device) error
	EnableDevice(uid, connId string) error
	GetAllDevices(uid string) ([]*model.Device, error)
	GetDevicesByConnID(uid, connID string) (*model.Device, error)
}
