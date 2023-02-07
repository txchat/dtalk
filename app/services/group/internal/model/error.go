package model

import "errors"

var (
	ErrPushMsgArg       = errors.New("rpc error: code = Unknown desc = rpc pushmsg arg error")
	ErrPermissionDenied = errors.New("permission denied")
)
