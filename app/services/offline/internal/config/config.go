package config

import (
	xkafka "github.com/txchat/pkg/mq/kafka"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	AppID                 string
	Pushers               map[string]Pusher
	DealConsumerConfig    xkafka.ConsumerConfig
	DealBatchConsumerConf xkafka.BatchConsumerConf
}

type Pusher struct {
	Env             string
	AppKey          string
	AppMasterSecret string
	MiActivity      string
}
