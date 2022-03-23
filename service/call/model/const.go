package model

const (
	// READYTIME 拨号最长持续时间
	READYTIME = 60
	BUSY      = true
	FREE      = false

	READY      = 0
	INPROGRESS = 1
	FINISH     = 2

	MAXROOMID = 1000000000

	// SESSIONMAXTIME session 在 redis 中过期的时间
	SESSIONMAXTIME = 60 * 60 * 24
)

type StopType int32

const (
	Busy    StopType = 0
	Timeout StopType = 1
	Reject  StopType = 2
	Hangup  StopType = 3
	Cancel  StopType = 4
)
