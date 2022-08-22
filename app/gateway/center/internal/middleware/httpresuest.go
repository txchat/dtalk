package middleware

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/txchat/dtalk/pkg/auth"
	xerror "github.com/txchat/dtalk/pkg/error"
	api "github.com/txchat/dtalk/pkg/newapi"
)

type AppAuthMiddleware struct {
}

func NewAppAuthMiddleware() *AppAuthMiddleware {
	return &AppAuthMiddleware{}
}

// Handle 处理
func (m *AppAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sig := r.Header.Get(api.HeaderSignature)
		uuid := r.Header.Get(api.HeaderUUID)
		device := r.Header.Get(api.HeaderDeviceType)
		deviceName := r.Header.Get(api.HeaderDeviceName)
		version := r.Header.Get(api.HeaderVersion)

		// TODO MOCK
		server := auth.NewDefaultApiAuthenticator()
		uid, err := server.Auth(sig)
		if err != nil {
			switch err {
			case auth.ERR_SIGNATUREEXPIRED:
				err = xerror.NewError(xerror.SignatureExpired)
			default:
				err = xerror.NewError(xerror.SignatureInvalid)
			}
			context.Set(r, api.ReqError, err)
			return
		}
		//set val
		context.Set(r, api.Signature, sig)
		context.Set(r, api.Address, uid)
		context.Set(r, api.UUID, uuid)
		context.Set(r, api.DeviceType, device)
		context.Set(r, api.DeviceName, deviceName)
		context.Set(r, api.Version, version)

		next(w, r)
	}
}
