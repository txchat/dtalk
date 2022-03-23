package api

import (
	"context"
	"time"
)

const (
	RespMiddleWareDisabled = "RespMiddleWareDisabled"

	ReqError  = "error"
	ReqResult = "result"
)

const (
	Address    = "address"
	Signature  = "signature"
	DeviceName = "deviceName"
	DeviceType = "deviceType"
	Uuid       = "uuid"
	Version    = "version"
)

const HeaderTimeOut = 120 * time.Second

func NewAddrWithContext(ctx context.Context) string {
	addr, ok := ctx.Value(Address).(string)
	if !ok {
		addr = ""
	}

	return addr
}
