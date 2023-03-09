package config

import (
	xkafka "github.com/oofpgDLD/kafka-go"
	xredis "github.com/txchat/dtalk/pkg/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	RedisDB                   xredis.Config
	AppID                     string
	GroupRPC                  zrpc.RpcClientConf
	LogicRPC                  zrpc.RpcClientConf
	PusherRPC                 zrpc.RpcClientConf
	ConnDealConsumerConfig    xkafka.ConsumerConfig
	ConnDealBatchConsumerConf xkafka.BatchConsumerConf
}
