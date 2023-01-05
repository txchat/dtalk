package config

import (
	xkafka "github.com/txchat/dtalk/pkg/mq/kafka"
	xmysql "github.com/txchat/dtalk/pkg/mysql"
	xredis "github.com/txchat/dtalk/pkg/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	RedisDB                    xredis.Config
	AppID                      string
	SyncCache                  bool
	MySQL                      xmysql.Config
	DeviceRPC                  zrpc.RpcClientConf
	PusherRPC                  zrpc.RpcClientConf
	GroupRPC                   zrpc.RpcClientConf
	StoreDealConsumerConfig    xkafka.ConsumerConfig
	StoreDealBatchConsumerConf xkafka.BatchConsumerConf
	SyncDealConsumerConfig     xkafka.ConsumerConfig
	SyncDealBatchConsumerConf  xkafka.BatchConsumerConf
}
