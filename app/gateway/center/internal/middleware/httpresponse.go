package middleware

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/txchat/dtalk/app/gateway/center/internal/types"
	api "github.com/txchat/dtalk/pkg/newapi"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type ResponseMiddleware struct {
}

func NewResponseMiddleware() *ResponseMiddleware {
	return &ResponseMiddleware{}
}

// Handle 处理
func (m *ResponseMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
		//获取 resp，error
		// resp := r.Context().Value(api.ReqResult)
		// err := r.Context().Value(api.ReqError)
		resp := context.Get(r, api.ReqResult)
		err := context.Get(r, api.ReqError)

		ret := convertHTTPResp(api.ParseHandlerResult(resp, err))
		//设置并返回结果
		httpx.OkJson(w, ret)
	}
}

func convertHTTPResp(code int, msg string, data interface{}) *types.HTTPBase {
	var ret types.HTTPBase
	ret.Result = code
	ret.Message = msg
	ret.Data = data
	return &ret
}
