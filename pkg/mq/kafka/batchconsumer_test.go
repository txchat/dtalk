package kafka

import (
	"context"
	"io"
	"strconv"
	"sync/atomic"
	"testing"
	"time"

	"github.com/Shopify/sarama"
	log "github.com/inconshreveable/log15"
	"github.com/stretchr/testify/assert"
)

type mockLegitimateConsumer struct {
	isClosed   int32
	queueEmpty chan *sarama.ConsumerMessage
}

func newMockLegitimateConsumer() *mockLegitimateConsumer {
	return &mockLegitimateConsumer{
		isClosed:   0,
		queueEmpty: make(chan *sarama.ConsumerMessage),
	}
}

func (c *mockLegitimateConsumer) FetchMessage(ctx context.Context) (message *sarama.ConsumerMessage, err error) {
	data, ok := <-c.queueEmpty
	if c.isClosed == 1 && !ok {
		return data, io.EOF
	}
	return data, nil
}

func (c *mockLegitimateConsumer) Close() error {
	var err error
	if atomic.CompareAndSwapInt32(&c.isClosed, 0, 1) {
		close(c.queueEmpty)
	}
	return err
}

func TestMockLegitimateConsumer(t *testing.T) {
	mc := newMockLegitimateConsumer()
	mc.Close()
	msg, err := mc.FetchMessage(context.Background())
	assert.ErrorIs(t, err, io.EOF)
	assert.Nil(t, msg)
}

func TestGracefulStopBatchConsumer(t *testing.T) {
	cacheSize := 60
	bc := NewBatchConsumer(BatchConsumerConf{
		CacheCapacity: cacheSize,
		Consumers:     1,
		Processors:    1,
	}, WithHandle(func(key string, data []byte) error {
		time.Sleep(time.Second)
		log.Debug("receive msg", "data", string(data))
		return nil
	}), newMockLegitimateConsumer())

	// 预先将缓冲填充满
	for i := 0; i < cacheSize; i++ {
		bc.channel <- &sarama.ConsumerMessage{
			Key:   nil,
			Value: []byte(strconv.FormatInt(int64(i), 10)),
		}
	}

	bc.Start()
	bc.GracefulStop(context.Background())
}

func TestForceStopBatchConsumer(t *testing.T) {
	cacheSize := 60
	bc := NewBatchConsumer(BatchConsumerConf{
		CacheCapacity: cacheSize,
		Consumers:     1,
		Processors:    1,
	}, WithHandle(func(key string, data []byte) error {
		time.Sleep(time.Second)
		log.Debug("receive msg", "data", string(data))
		return nil
	}), newMockLegitimateConsumer())

	// 预先将缓冲填充满
	for i := 0; i < cacheSize; i++ {
		bc.channel <- &sarama.ConsumerMessage{
			Key:   nil,
			Value: []byte(strconv.FormatInt(int64(i), 10)),
		}
	}

	bc.Start()
	bc.Stop()
}

func TestGracefulStopBatchConsumerWithTimeout(t *testing.T) {
	cacheSize := 60
	bc := NewBatchConsumer(BatchConsumerConf{
		CacheCapacity: cacheSize,
		Consumers:     1,
		Processors:    1,
	}, WithHandle(func(key string, data []byte) error {
		time.Sleep(time.Second)
		log.Debug("receive msg", "data", string(data))
		return nil
	}), newMockLegitimateConsumer())

	// 预先将缓冲填充满
	for i := 0; i < cacheSize; i++ {
		bc.channel <- &sarama.ConsumerMessage{
			Key:   nil,
			Value: []byte(strconv.FormatInt(int64(i), 10)),
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	bc.Start()
	bc.GracefulStop(ctx)
}
