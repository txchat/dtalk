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
	AnswerRPC zrpc.RpcClientConf
	GroupRPC  zrpc.RpcClientConf
	LogicRPC  zrpc.RpcClientConf

	ProducerStorage           xkafka.ProducerConfig
	OffPushEnabled            bool
	ProducerOffPush           xkafka.ProducerConfig
	ConnDealConsumerConfig    xkafka.ConsumerConfig
	ConnDealBatchConsumerConf xkafka.BatchConsumerConf
	PushDealConsumerConfig    xkafka.ConsumerConfig
	PushDealBatchConsumerConf xkafka.BatchConsumerConf
}
