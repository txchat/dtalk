package middleware

import (
	"net/http"

	xhttp "github.com/txchat/dtalk/pkg/net/http"
)

type AppParseHeaderMiddleware struct {
}

func NewAppParseHeaderMiddleware() *AppParseHeaderMiddleware {
	return &AppParseHeaderMiddleware{}
}

func (m *AppParseHeaderMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// parse header
		bizHeader := xhttp.ConvertCustom(r)
		//set context values
		next(w, bizHeader.SetWithRequest(r))
	}
}
