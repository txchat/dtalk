package config

import (
	"time"

	xkafka "github.com/oofpgDLD/kafka-go"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	AppID                 string                   `json:","`
	HandleTimeout         time.Duration            `json:",default=10m"`
	Pushers               map[string]Pusher        `json:","`
	DealConsumerConfig    xkafka.ConsumerConfig    `json:","`
	DealBatchConsumerConf xkafka.BatchConsumerConf `json:","`
}

type Pusher struct {
	Env             string
	AppKey          string
	AppMasterSecret string
	MiActivity      string
}
