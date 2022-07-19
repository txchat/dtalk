package mock

import (
	"context"
	"fmt"

	"github.com/txchat/dtalk/service/record/kafka/publisher"
	"gopkg.in/Shopify/sarama.v1"
	kafka "gopkg.in/Shopify/sarama.v1"
)

type Client struct {
	brokers    []string
	appId      string
	offPushPub kafka.SyncProducer
}

func NewClient(appId string, brokers []string) *Client {
	c := &Client{
		appId:      appId,
		brokers:    brokers,
		offPushPub: publisher.NewKafkaPub(brokers),
	}
	return c
}

func (c *Client) PublishOfflineMsg(ctx context.Context, fromId string, b []byte) error {
	if c.offPushPub == nil {
		return fmt.Errorf("kafka publish client not init")
	}
	appTopic := fmt.Sprintf("offpush-%s-topic", c.appId)
	m := &sarama.ProducerMessage{
		Key:   sarama.StringEncoder(fromId),
		Topic: appTopic,
		Value: sarama.ByteEncoder(b),
	}
	_, _, err := c.offPushPub.SendMessage(m)
	return err
}
