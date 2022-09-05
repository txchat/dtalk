package config

import (
	xredis "github.com/txchat/dtalk/pkg/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	RedisDB   xredis.Config
	IDGenRPC  zrpc.RpcClientConf
	AnswerRPC zrpc.RpcClientConf
	RTC       RTCConfig
}

type RTCConfig struct {
	CallingTimeout int64
	SDKAppId       int
	SecretKey      string
	Expire         int
}
