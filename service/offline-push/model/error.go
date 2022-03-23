package model

import "errors"

var (
	ErrAppId            = errors.New("appId not compared")
	ErrConsumeRedo      = errors.New("process msg failed")
	ErrCustomNotSupport = errors.New("pusher device type not support")
)
