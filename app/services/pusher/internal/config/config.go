package config

import (
	xkafka "github.com/oofpgDLD/kafka-go"
	xredis "github.com/txchat/dtalk/pkg/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	AppID                     string
	RedisDB                   xredis.Config
	DeviceRPC                 zrpc.RpcClientConf
	LogicRPC                  zrpc.RpcClientConf
	OffPushEnabled            bool
	Producer                  xkafka.ProducerConfig
	PushDealConsumerConfig    xkafka.ConsumerConfig
	PushDealBatchConsumerConf xkafka.BatchConsumerConf
}
