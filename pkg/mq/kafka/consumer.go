package kafka

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	"github.com/inconshreveable/log15"
)

type ConsumerConfig struct {
	Version        string
	Brokers        []string
	Group          string
	Topic          string
	ConnectTimeout time.Duration
	realVersion    sarama.KafkaVersion
}

type Consumer struct {
	log log15.Logger
	cfg ConsumerConfig
	ctx context.Context

	client sarama.ConsumerGroup
	queue  chan *sarama.ConsumerMessage

	ready chan bool
}

func ensureConsumerConfig(cfg *ConsumerConfig) {
	//version
	version, err := sarama.ParseKafkaVersion(cfg.Version)
	if err != nil {
		cfg.realVersion = sarama.V2_0_0_0
	}
	cfg.realVersion = version

	if len(cfg.Brokers) == 0 {
		panic(errors.New("brokers is empty"))
	}
	if cfg.Topic == "" {
		panic(errors.New("topic is empty"))
	}
}

func NewConsumer(cfg ConsumerConfig, logger log15.Logger) *Consumer {
	ensureConsumerConfig(&cfg)
	c := sarama.NewConfig()
	c.Consumer.Return.Errors = true
	c.Version = cfg.realVersion
	//c.Consumer.Offsets.AutoCommit.Interval=time.Second

	if logger == nil {
		logger = log15.New()
	}

	//ctx, _ := context.WithTimeout(context.Background(), cfg.ConnectTimeout)
	consumer := &Consumer{
		log:    logger,
		cfg:    cfg,
		ctx:    context.Background(),
		client: nil,
		queue:  make(chan *sarama.ConsumerMessage, 100),
		ready:  make(chan bool),
	}

	client, err := sarama.NewConsumerGroup(cfg.Brokers, cfg.Group, c)
	if err != nil {
		panic(err)
	}
	consumer.client = client

	go func() {
		for {
			// `Consume` should be called inside an infinite loop, when a
			// server-side rebalance happens, the consumer session will need to be
			// recreated to get the new claims
			if err := client.Consume(consumer.ctx, []string{cfg.Topic}, consumer); err != nil {
				panic(fmt.Errorf("error from consumer: %v", err))
			}
			// check if context was cancelled, signaling that the consumer should stop
			if consumer.ctx.Err() != nil {
				return
			}
			//consumer.ready = make(chan bool)
		}
	}()

	//block create with timeout
	select {
	case <-consumer.ctx.Done():
		panic(consumer.ctx.Err())
	case <-consumer.ready:
	}
	return consumer
}

// FetchMessage 读取并返回message
func (c *Consumer) FetchMessage(ctx context.Context) (message *sarama.ConsumerMessage, err error) {
	return <-c.queue, nil
}

func (c *Consumer) Close() error {
	return c.client.Close()
}

// Setup is run at the beginning of a new session, before ConsumeClaim.
func (c *Consumer) Setup(sarama.ConsumerGroupSession) error {
	close(c.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
// but before the offsets are committed for the very last time.
func (c *Consumer) Cleanup(session sarama.ConsumerGroupSession) error {
	c.log.Info("consumer rebalanced", "member id", session.MemberID(), "id", session.GenerationID())
	c.ready = make(chan bool)
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// Once the Messages() channel is closed, the Handler must finish its processing
// loop and exit.
func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/main/consumer_group.go#L27-L29
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				c.log.Error("consumerMessage queue closed", "brokers", strings.Join(c.cfg.Brokers, ","),
					"group", c.cfg.Group)
				return errors.New("consumerMessage queue closed")
			}
			c.log.Info("Message claimed", "value", string(message.Value), "timestamp", message.Timestamp, "topic", message.Topic, "offset", message.Offset)
			c.queue <- message
			session.MarkMessage(message, "")

		// Should return when `session.Context()` is done.
		// If not, will raise `ErrRebalanceInProgress` or `read tcp <ip>:<port>: i/o timeout` when kafka rebalance. see:
		// https://github.com/Shopify/sarama/issues/1192
		case <-session.Context().Done():
			c.log.Info("consume claim down", "id", session.GenerationID())
			return nil
		}
	}
}
