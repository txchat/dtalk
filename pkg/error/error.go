package error

import (
	"fmt"
	"reflect"
	"regexp"

	"github.com/txchat/dtalk/pkg/util"
	"google.golang.org/grpc/status"
)

var (
	descRegex = regexp.MustCompile(`code: ([-|+]?\d+), msg: (.+)`)
)

func NewError(code int64, msg string) *Error {
	return &Error{
		code: code,
		msg:  msg,
	}
}

type Error struct {
	code int64
	// message
	msg string
}

func (e *Error) Code() int64 {
	return e.code
}

func (e *Error) Error() string {
	return e.msg
}

func (e *Error) String() string {
	return fmt.Sprintf("code: %d, msg: %s",
		e.code, e.msg)
}

func (e *Error) Equal(err error) bool {
	if e != err {
		return reflect.DeepEqual(e, err)
	}
	return true
}

// ParseError 解析业务错误
func ParseError(err error) *Error {
	if err == nil {
		return NoErr
	}

	if e, ok := err.(*Error); ok {
		return e
	}

	s, _ := status.FromError(err)
	c, m := uint32(s.Code()), s.Message()
	if c < grpcMaxCode {
		return ErrUnexpected
	}
	if e, ok := ConvertError(m); ok {
		return e
	}
	return ErrUnexpected
}

// ConvertError 解析业务消息对应的业务错误
func ConvertError(msg string) (*Error, bool) {
	gs := descRegex.FindStringSubmatch(msg)
	if len(gs) == 3 {
		return &Error{
			code: util.MustToInt64(gs[1]),
			msg:  gs[2],
		}, true
	}

	return ErrUnexpected, false
}
