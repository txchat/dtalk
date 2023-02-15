package config

import (
	xkafka "github.com/txchat/pkg/mq/kafka"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	AppID             string
	Producer          xkafka.ProducerConfig
	ConsumerConfig    xkafka.ConsumerConfig
	BatchConsumerConf xkafka.BatchConsumerConf
}
