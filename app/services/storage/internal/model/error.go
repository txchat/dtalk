package model

import "errors"

var (
	ErrAppID         = errors.New("app_id not compared")
	ErrRecordNotFind = errors.New("record not find")
	ErrChannelType   = errors.New("channel type not supported")
	ErrTargetIsEmpty = errors.New("message target is empty")

	ErrConsumeRedo = errors.New("process msg failed")
	ErrFrameType   = errors.New("error frame type")
)
