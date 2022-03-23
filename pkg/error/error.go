package error

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
)

const (
	DispMsgExt  = 0
	DispJustExt = 1
)

//错误信息组合
//1. message + : + extMsg
//2. extMsg
func NewError(code int) *Error {
	return &Error{
		code: code,
		data: make(map[string]interface{}),
	}
}

type Error struct {
	code int
	//额外错误信息
	extMsg string
	//暴露给接口，但客户端不显示
	data map[string]interface{}
	//显示方式
	displayType int
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Data() map[string]interface{} {
	return e.data
}

//策略返回显示的错误信息
func (e *Error) Error() string {
	switch e.displayType {
	case DispJustExt:
		return e.extMsg
	default:
		return strings.TrimRight(fmt.Sprintf("%s%s%s", errMsg(e.code), ":", e.extMsg), `:`)
	}
}

//external
func (e *Error) JustShowExtMsg() *Error {
	e.displayType = DispJustExt
	return e
}

func (e *Error) SetExtMessage(extMsg string) *Error {
	e.extMsg = extMsg
	return e
}

//data["service"]=service name
//data["code"]=code
//data["message"]=message
func (e *Error) SetChildErr(name string, code interface{}, message interface{}) *Error {
	if e.data == nil {
		e.data = make(map[string]interface{})
	}
	e.data["service"] = name
	e.WriteCode(code)
	e.WriteMessage(message)
	return e
}

func (e *Error) WriteMessage(msg interface{}) *Error {
	if e.data == nil {
		e.data = make(map[string]interface{})
	}
	e.data["message"] = msg
	return e
}

func (e *Error) WriteCode(code interface{}) *Error {
	if e.data == nil {
		e.data = make(map[string]interface{})
	}
	e.data["code"] = code
	return e
}

func errMsg(code int) string {
	errStr, ok := errorMsg[code]
	if !ok {
		log.Error().Err(fmt.Errorf("error code not find"))
		return ""
	}
	return errStr
}
