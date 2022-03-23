package dao

import (
	"context"
	"fmt"
	"gopkg.in/Shopify/sarama.v1"
)

// PushMsg push a message to databus.
func (d *Dao) PublishToSend(ctx context.Context, fromId string, data []byte) (err error) {
	appTopic := fmt.Sprintf("received-%s-topic", d.appId)
	m := &sarama.ProducerMessage{
		Key:   sarama.StringEncoder(fromId),
		Topic: appTopic,
		Value: sarama.ByteEncoder(data),
	}
	if _, _, err = d.mqPub.SendMessage(m); err != nil {
		return
	}
	return
}
