package http

import (
	"net/http"
	"reflect"

	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type Response struct {
	Result  int64       `json:"result"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// OkJSON 成功json响应返回
func OkJSON(w http.ResponseWriter, r *http.Request, v interface{}) {
	if isNil(v) {
		v = struct{}{}
	}
	httpx.WriteJson(w, http.StatusOK, &Response{
		Result:  xerror.CodeOK,
		Message: xerror.MsgOK,
		Data:    v,
	})
}

// Error 错误响应返回
func Error(w http.ResponseWriter, r *http.Request, err error) {
	ctx := r.Context()
	logx.WithContext(ctx).Errorf("request handle err, err: %+v", err)

	e := xerror.ParseError(err)

	httpx.WriteJson(w, http.StatusOK, &Response{
		Result:  e.Code(),
		Message: e.Error(),
		Data:    e.Data(),
	})
}

func isNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}
