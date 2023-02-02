package kafka

import (
	"context"
	"io"
	"sync/atomic"

	"github.com/Shopify/sarama"
	"github.com/gammazero/workerpool"
	"github.com/inconshreveable/log15"
)

type IConsumer interface {
	FetchMessage(ctx context.Context) (message *sarama.ConsumerMessage, err error)
	Close() error
}

type ConsumeHandle func(key string, data []byte) error

type ConsumeHandler interface {
	Consume(key string, data []byte) error
}

const (
	defaultQueueCapacity = 1000
)

type (
	BatchConsumerConf struct {
		CacheCapacity int `json:",optional"`
		Consumers     int `json:",optional"`
		Processors    int `json:",optional"`
	}

	batchConsumerOptions struct {
		logger log15.Logger
	}
	BatchConsumerOption func(*batchConsumerOptions)

	BatchConsumer struct {
		cfg             BatchConsumerConf
		log             log15.Logger
		consumer        IConsumer
		channel         chan *sarama.ConsumerMessage
		handler         ConsumeHandler
		producerWorkers *workerpool.WorkerPool
		consumerWorkers *workerpool.WorkerPool
		isClosed        int32
	}
)

func ensureConfigOptions(cfg *BatchConsumerConf, options *batchConsumerOptions) {
	if options.logger == nil {
		options.logger = log15.New()
	}
	if cfg.CacheCapacity <= 0 {
		cfg.CacheCapacity = defaultQueueCapacity
	}
	if cfg.Consumers == 0 {
		cfg.Consumers = 8
	}
	if cfg.Processors == 0 {
		cfg.Processors = 8
	}
}

func NewBatchConsumer(cfg BatchConsumerConf, handler ConsumeHandler, consumer IConsumer, opts ...BatchConsumerOption) *BatchConsumer {
	var options batchConsumerOptions
	for _, opt := range opts {
		opt(&options)
	}
	ensureConfigOptions(&cfg, &options)

	return &BatchConsumer{
		log:             options.logger,
		consumer:        consumer,
		channel:         make(chan *sarama.ConsumerMessage, cfg.CacheCapacity),
		handler:         handler,
		cfg:             cfg,
		consumerWorkers: workerpool.New(cfg.Processors),
		producerWorkers: workerpool.New(cfg.Consumers),
	}
}

func (bc *BatchConsumer) startProducers() {
	for i := 0; i < bc.cfg.Consumers; i++ {
		bc.producerWorkers.Submit(func() {
			defer bc.close()

			for {
				msg, err := bc.consumer.FetchMessage(context.TODO())
				// io.EOF means consumer closed
				// io.ErrClosedPipe means committing messages on the consumer,
				// kafka will refire the messages on uncommitted messages, ignore
				if err == io.EOF || err == io.ErrClosedPipe {
					bc.log.Info("fetchMessage io.EOF or io.ErrClosedPipe")
					return
				}
				if err != nil {
					bc.log.Error("Error on reading message", "err", err.Error())
					continue
				}
				if msg == nil {
					continue
				}
				bc.log.Debug("fetchMessage", "msg", string(msg.Value))
				bc.channel <- msg
			}
		})
	}
}

func (bc *BatchConsumer) startConsumers() {
	for i := 0; i < bc.cfg.Processors; i++ {
		bc.consumerWorkers.Submit(func() {
			for msg := range bc.channel {
				if err := bc.consumeOne(string(msg.Key), msg.Value); err != nil {
					bc.log.Error("Error on consuming message", "msg", string(msg.Value), "err", err)
				}
			}
		})
	}
}

func (bc *BatchConsumer) consumeOne(key string, data []byte) error {
	return bc.handler.Consume(key, data)
}

func (bc *BatchConsumer) close() {
	if atomic.CompareAndSwapInt32(&bc.isClosed, 0, 1) {
		bc.consumer.Close()
		close(bc.channel)
	}
}

func (bc *BatchConsumer) Start() {
	bc.startConsumers()
	bc.startProducers()
}

func (bc *BatchConsumer) Stop() {
	bc.close()
}

func (bc *BatchConsumer) GracefulStop(ctx context.Context) {
	down := make(chan struct{})
	go func() {
		bc.close()
		bc.consumerWorkers.StopWait()
		close(down)
	}()

	select {
	case <-ctx.Done():
		return
	case <-down:
		return
	}
}

func WithLogger(logger log15.Logger) BatchConsumerOption {
	return func(options *batchConsumerOptions) {
		options.logger = logger
	}
}

func WithHandle(handle ConsumeHandle) ConsumeHandler {
	return innerConsumeHandler{
		handle: handle,
	}
}

type innerConsumeHandler struct {
	handle ConsumeHandle
}

func (ch innerConsumeHandler) Consume(k string, v []byte) error {
	return ch.handle(k, v)
}
