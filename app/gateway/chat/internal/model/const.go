package model

const (
	Private = 0
	Group   = 1
)

const MuteForever = int64(^uint(0) >> 1) // 永久禁言的时间 9223372036854775807

const GroupManagerLimit = 10

const MaxPartSize = 1024 * 1024 * 1024 * 5

const (
	Oss_Aliyun    = "aliyun"
	Oss_Huaweiuyn = "huaweiyun"
	Oss_Minio     = "minio"
)
