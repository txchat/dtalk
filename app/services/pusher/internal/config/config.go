package config

import (
	xkafka "github.com/txchat/dtalk/pkg/mq/kafka"
	xgrpc "github.com/txchat/dtalk/pkg/net/grpc"
	xredis "github.com/txchat/dtalk/pkg/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	AppID     string
	RedisDB   xredis.Config
	DeviceRPC zrpc.RpcClientConf
	AnswerRPC zrpc.RpcClientConf

	ProducerStorage           xkafka.ProducerConfig
	OffPushEnabled            bool
	ProducerOffPush           xkafka.ProducerConfig
	ConnDealConsumerConfig    xkafka.ConsumerConfig
	ConnDealBatchConsumerConf xkafka.BatchConsumerConf
	PushDealConsumerConfig    xkafka.ConsumerConfig
	PushDealBatchConsumerConf xkafka.BatchConsumerConf

	LogicRPCClient xgrpc.RPCClient
	GroupRPCClient xgrpc.RPCClient
}
