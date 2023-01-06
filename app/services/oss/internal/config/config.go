package config

import (
	xredis "github.com/txchat/dtalk/pkg/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	RedisDB xredis.Config
	Oss     []Oss
}

type Oss struct {
	AppId   string
	OssType string
	// RegionId 决定获得 tempKey 的速度快不快
	RegionId        string
	AccessKeyId     string
	AccessKeySecret string
	Role            string
	Policy          string
	DurationSeconds int
	Bucket          string
	// EndPoint 资源终端
	EndPoint  string
	PublicUrl string
}
