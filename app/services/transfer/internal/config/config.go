package config

import (
	xkafka "github.com/oofpgDLD/kafka-go"
	xredis "github.com/txchat/dtalk/pkg/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	AppID             string
	Node              int64 `json:",default=1"`
	RedisDB           xredis.Config
	PusherRPC         zrpc.RpcClientConf
	GroupRPC          zrpc.RpcClientConf
	Producer          xkafka.ProducerConfig
	ConsumerConfig    xkafka.ConsumerConfig
	BatchConsumerConf xkafka.BatchConsumerConf
}
