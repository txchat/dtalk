package config

import (
	"time"

	xkafka "github.com/oofpgDLD/kafka-go"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	AppID                 string
	Timeout               time.Duration
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
