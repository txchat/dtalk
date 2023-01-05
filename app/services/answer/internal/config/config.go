package config

import (
	xkafka "github.com/txchat/dtalk/pkg/mq/kafka"
	xgrpc "github.com/txchat/dtalk/pkg/net/grpc"
	xredis "github.com/txchat/dtalk/pkg/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	RedisDB           xredis.Config
	IDGenRPC          zrpc.RpcClientConf
	AppID             string
	LogicRPCClient    xgrpc.RPCClient
	GroupRPC          zrpc.RpcClientConf
	Producer          xkafka.ProducerConfig
	ConsumerConfig    xkafka.ConsumerConfig
	BatchConsumerConf xkafka.BatchConsumerConf
}
