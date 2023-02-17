package config

import (
	xredis "github.com/txchat/dtalk/pkg/redis"
	xkafka "github.com/txchat/pkg/mq/kafka"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	AppID     string
	RedisDB   xredis.Config
	DeviceRPC zrpc.RpcClientConf
	LogicRPC  zrpc.RpcClientConf

	OffPushEnabled            bool
	ProducerOffPush           xkafka.ProducerConfig
	PushDealConsumerConfig    xkafka.ConsumerConfig
	PushDealBatchConsumerConf xkafka.BatchConsumerConf
}
