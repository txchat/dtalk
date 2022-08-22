package api

import (
	"reflect"

	xerror "github.com/txchat/dtalk/pkg/error"
)

// ParseHandlerResult 将http handler执行结果解析为http的code和error message
func ParseHandlerResult(result interface{}, err interface{}) (code int, msg string, data interface{}) {
	if err != nil {
		switch errType := err.(type) {
		case *xerror.Error:
			code = errType.Code()
			msg = errType.Error()
			data = errType.Data()
		default:
			e := xerror.NewError(xerror.CodeInnerError)
			code = e.Code()
			msg = e.Error()
			data = err
		}
		return
	}

	code = xerror.CodeOK
	if isNil(result) {
		data = make(map[string]interface{})
	} else {
		data = result
	}
	return
}

func isNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}
