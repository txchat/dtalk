package mock

import (
	"context"
	"strconv"
	"time"

	"github.com/Shopify/sarama"
)

type Consumer struct {
	index  int64
	ticker *time.Ticker
	queue  chan *sarama.ConsumerMessage
	closer chan struct{}
}

func NewConsumer(dur time.Duration) *Consumer {
	c := &Consumer{
		ticker: time.NewTicker(dur),
		queue:  make(chan *sarama.ConsumerMessage, 100),
		closer: make(chan struct{}),
	}
	go func() {
		for {
			select {
			case <-c.ticker.C:
				c.index++
				c.queue <- &sarama.ConsumerMessage{
					Headers:        nil,
					Timestamp:      time.Time{},
					BlockTimestamp: time.Time{},
					Key:            nil,
					Value:          []byte(strconv.FormatInt(c.index, 10)),
					Topic:          "",
					Partition:      0,
					Offset:         0,
				}
			case <-c.closer:
				return
			}
		}
	}()
	return c
}

func (c *Consumer) Close() error {
	close(c.closer)
	return nil
}

func (c *Consumer) FetchMessage(ctx context.Context) (message *sarama.ConsumerMessage, err error) {
	return <-c.queue, nil
}
