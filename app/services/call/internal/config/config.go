package config

import (
	xredis "github.com/txchat/dtalk/pkg/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	RedisDB  xredis.Config
	RTC      RTCConfig
	IDGenRPC zrpc.RpcClientConf
}

type RTCConfig struct {
	CallingTimeout int64
	SDKAppId       int
	SecretKey      string
	Expire         int
}
