package config

import (
	xkafka "github.com/txchat/dtalk/pkg/mq/kafka"
	xredis "github.com/txchat/dtalk/pkg/redis"
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
