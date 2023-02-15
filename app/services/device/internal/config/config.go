package config

import (
	xredis "github.com/txchat/dtalk/pkg/redis"
	xkafka "github.com/txchat/pkg/mq/kafka"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	RedisDB                   xredis.Config
	AppID                     string
	GroupRPC                  zrpc.RpcClientConf
	LogicRPC                  zrpc.RpcClientConf
	ConnDealConsumerConfig    xkafka.ConsumerConfig
	ConnDealBatchConsumerConf xkafka.BatchConsumerConf
}
