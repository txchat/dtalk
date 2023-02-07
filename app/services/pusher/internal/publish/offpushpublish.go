package publish

import (
	"context"
	"fmt"

	xkafka "github.com/txchat/pkg/mq/kafka"
)

type OffPush struct {
	AppID string
	topic string
	conn  *xkafka.Producer
}

func NewOffPushPublish(appID string, cfg xkafka.ProducerConfig) *OffPush {
	conn := xkafka.NewProducer(cfg)
	return &OffPush{
		AppID: appID,
		topic: fmt.Sprintf("offpush-%s-topic", appID),
		conn:  conn,
	}
}

func (p *OffPush) PublishOfflineMsg(ctx context.Context, fromId string, b []byte) (err error) {
	_, _, err = p.conn.Publish(p.topic, fromId, b)
	return
}
