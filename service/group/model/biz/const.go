package biz

const (
	DisPlayNum = 10 // 查看群信息默认显示的人数

	MuteMaximum = int64(^uint(0) >> 1) // 永久禁言的时间 9223372036854775807
)

var (
	GroupMaximum int32 = 2000
	AdminNum     int32 = 10
)
