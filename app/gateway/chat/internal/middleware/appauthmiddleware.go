package middleware

import (
	"net/http"

	"github.com/txchat/dtalk/app/gateway/chat/internal/middleware/authmock"
	xerror "github.com/txchat/dtalk/pkg/error"
	xhttp "github.com/txchat/dtalk/pkg/net/http"
	"github.com/txchat/pkg/auth"
)

type AppAuthMiddleware struct {
	mock authmock.Mock
}

func NewAppAuthMiddleware(mock authmock.Mock) *AppAuthMiddleware {
	return &AppAuthMiddleware{
		mock: mock,
	}
}

func (m *AppAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// parse header
		custom := xhttp.ConvertCustom(r)

		// MOCK
		if m.mock != nil {
			uid := m.mock.Signature(custom.Signature)
			if uid != "" {
				//set context values
				custom.UID = uid
				next(w, custom.SetWithRequest(r))
				return
			}
		}
		server := auth.NewDefaultAPIAuthenticator()
		uid, err := server.Auth(custom.Signature)
		if err != nil {
			switch err {
			case auth.ErrSignatureExpired:
				err = xerror.ErrSignatureExpired
			default:
				err = xerror.ErrSignatureInvalid
			}
			xhttp.Error(w, r, err)
			return
		}
		//set context values
		custom.UID = uid
		next(w, custom.SetWithRequest(r))
	}
}
