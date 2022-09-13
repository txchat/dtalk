package kafka

import (
	"context"
	"io"
	"testing"
	"time"

	"github.com/Shopify/sarama"
	"github.com/inconshreveable/log15"
	"github.com/stretchr/testify/assert"
)

type mockConsumeGroup struct {
	sarama.ConsumerGroup
	handler sarama.ConsumerGroupHandler
}

func NewMockConsumeGroup(handler sarama.ConsumerGroupHandler) *mockConsumeGroup {
	return &mockConsumeGroup{
		handler: handler,
	}
}

// Close() override
func (mc *mockConsumeGroup) Close() error {
	mc.handler.Cleanup(nil)
	return nil
}

func TestCloseConsumer(t *testing.T) {
	c := &Consumer{
		log: log15.New(),
		cfg: ConsumerConfig{
			Version:        "",
			Brokers:        []string{"broker"},
			Group:          "group",
			Topic:          "topic",
			ConnectTimeout: time.Second * 7,
		},
		ctx:   context.Background(),
		queue: make(chan *sarama.ConsumerMessage, 100),
		ready: make(chan bool),
	}
	c.client = NewMockConsumeGroup(c)
	c.Close()

	msg, err := c.FetchMessage(context.Background())
	assert.ErrorIs(t, err, io.EOF)
	assert.Nil(t, msg)
}
