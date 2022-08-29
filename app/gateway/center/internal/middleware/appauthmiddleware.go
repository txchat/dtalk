package middleware

import (
	"net/http"

	"github.com/txchat/dtalk/pkg/auth"
	xerror "github.com/txchat/dtalk/pkg/error"
	xhttp "github.com/txchat/dtalk/pkg/net/http"
)

type AppAuthMiddleware struct {
}

func NewAppAuthMiddleware() *AppAuthMiddleware {
	return &AppAuthMiddleware{}
}

func (m *AppAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// parse header
		bizHeader := xhttp.ConvertCustom(r)

		// TODO MOCK
		server := auth.NewDefaultApiAuthenticator()
		uid, err := server.Auth(bizHeader.Signature)
		if err != nil {
			switch err {
			case auth.ERR_SIGNATUREEXPIRED:
				err = xerror.ErrSignatureExpired
			default:
				err = xerror.ErrSignatureInvalid
			}
			xhttp.Error(w, r, err)
			return
		}
		//set context values
		bizHeader.UID = uid
		next(w, bizHeader.SetWithRequest(r))
	}
}
