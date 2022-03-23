package model

import (
	"github.com/pkg/errors"
)

var (
	//ErrAdminPermission = errors.New("需要更高的权限")
	//ErrGroupIdNotExist = errors.New("该群号不存在")
	ErrMemberNotExist = errors.New("该用户不在本群中")
	//ErrPersonNotExist  = errors.New("你已不在本群中")
	ErrType = errors.New("参数不存在")
	//ErrAdminNum        = errors.New("管理员数量已满")

	ErrPushMsgArg     = errors.New("rpc error: code = Unknown desc = rpc pushmsg arg error")
	ErrRecordNotExist = errors.New("record not exist.")
)
