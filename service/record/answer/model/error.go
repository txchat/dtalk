package model

import "errors"

var (
	ErrAppId            = errors.New("appId not compared")
	ErrConsumeRedo      = errors.New("process msg failed")
	ErrCustomNotSupport = errors.New("biz not support")

	ErrGroupMemberNotExists = errors.New("group member not exists")
)

//check error
var (
	ErrorEnvType = errors.New("unsupported event type")
	ErrorChType  = errors.New("unsupported channel type")
	ErrorMsgType = errors.New("unsupported message type")
)
