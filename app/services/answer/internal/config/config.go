package config

import (
	xredis "github.com/txchat/dtalk/pkg/redis"
	xkafka "github.com/txchat/pkg/mq/kafka"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	AppID             string
	RedisDB           xredis.Config
	IDGenRPC          zrpc.RpcClientConf
	LogicRPC          zrpc.RpcClientConf
	GroupRPC          zrpc.RpcClientConf
	Producer          xkafka.ProducerConfig
	ConsumerConfig    xkafka.ConsumerConfig
	BatchConsumerConf xkafka.BatchConsumerConf
}
