package model

const (
	Private = 0
	Group   = 1
)

const MuteForever = int64(^uint(0) >> 1) // 永久禁言的时间 9223372036854775807

const GroupManagerLimit = 10
