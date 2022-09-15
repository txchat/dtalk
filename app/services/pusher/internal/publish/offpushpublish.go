package publish

import (
	"context"
	"fmt"

	xkafka "github.com/txchat/dtalk/pkg/mq/kafka"
)

type OffPush struct {
	AppID string
	conn  *xkafka.Producer
}

func NewOffPushPublish(appID string, cfg xkafka.ProducerConfig) *OffPush {
	appTopic := fmt.Sprintf("offpush-%s-topic", appID)
	cfg.Topic = appTopic
	conn := xkafka.NewProducer(cfg)
	return &OffPush{
		AppID: appID,
		conn:  conn,
	}
}

func (p *OffPush) PublishOfflineMsg(ctx context.Context, fromId string, b []byte) (err error) {
	_, _, err = p.conn.Publish(fromId, b)
	return
}
