package model

import (
	"errors"
)

var (
	ErrorEnvType = errors.New("unsupported event type")
	ErrorChType  = errors.New("unsupported channel type")
	ErrorMsgType = errors.New("unsupported message type")
)
