package middleware

import (
	"net/http"

	"github.com/gorilla/context"
	api "github.com/txchat/dtalk/pkg/newapi"
)

type AppParseHeaderMiddleware struct {
}

func NewAppParseHeaderMiddleware() *AppParseHeaderMiddleware {
	return &AppParseHeaderMiddleware{}
}

// Handle 处理
func (m *AppParseHeaderMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uuid := r.Header.Get(api.HeaderUUID)
		device := r.Header.Get(api.HeaderDeviceType)
		deviceName := r.Header.Get(api.HeaderDeviceName)
		version := r.Header.Get(api.HeaderVersion)

		//set val
		context.Set(r, api.UUID, uuid)
		context.Set(r, api.DeviceType, device)
		context.Set(r, api.DeviceName, deviceName)
		context.Set(r, api.Version, version)

		next(w, r)
	}
}
