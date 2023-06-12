package model

import (
	"errors"
)

var (
	ErrorEnvType = errors.New("unsupported event type")
	ErrorChType  = errors.New("unsupported channel type")
	ErrorMsgType = errors.New("unsupported message type")

	ErrAppID            = errors.New("app_id not compared")
	ErrCustomNotSupport = errors.New("biz not support")
)
