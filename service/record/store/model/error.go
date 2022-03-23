package model

import "errors"

var (
	ErrRecordNotFind    = errors.New("record not find")
	ErrAppId            = errors.New("appId not compared")
	ErrConsumeRedo      = errors.New("process msg failed")
	ErrCustomNotSupport = errors.New("biz not support")
	ErrFrameType        = errors.New("error frame type")
)
