package config

import (
	"github.com/go-sql-driver/mysql"
	xkafka "github.com/oofpgDLD/kafka-go"
	xredis "github.com/txchat/dtalk/pkg/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	RedisDB                    xredis.Config
	AppID                      string
	SyncCache                  bool
	MySQL                      mysql.Config
	DeviceRPC                  zrpc.RpcClientConf
	PusherRPC                  zrpc.RpcClientConf
	GroupRPC                   zrpc.RpcClientConf
	StoreDealConsumerConfig    xkafka.ConsumerConfig
	StoreDealBatchConsumerConf xkafka.BatchConsumerConf
	SyncDealConsumerConfig     xkafka.ConsumerConfig
	SyncDealBatchConsumerConf  xkafka.BatchConsumerConf
}
