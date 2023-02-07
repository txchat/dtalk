package model

import "errors"

var (
	ErrRecordNotFind = errors.New("record not find")
	ErrConsumeRedo   = errors.New("process msg failed")
	ErrFrameType     = errors.New("error frame type")

	ErrAppID            = errors.New("app_id not compared")
	ErrCustomNotSupport = errors.New("biz not support")
)
