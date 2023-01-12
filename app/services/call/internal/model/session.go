package model

import xcall "github.com/txchat/dtalk/internal/call"

const (
	// SessionTimeout session 在 redis 中过期的时间
	SessionTimeout = 60 * 60 * 24
)

type Session xcall.Session
