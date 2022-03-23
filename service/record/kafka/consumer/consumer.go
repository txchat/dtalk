package consumer

import (
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"github.com/rs/zerolog/log"
)

type DealHandler func(msg *sarama.ConsumerMessage) error

type Process interface {
	Deal(msg *sarama.ConsumerMessage) error
}

type Consumer struct {
	*cluster.Consumer
}

func (c *Consumer) Listen(deal DealHandler) {
	for {
		select {
		case err := <-c.Errors():
			log.Error().Err(err).Msg("consumer error")
		case n := <-c.Notifications():
			log.Info().Interface("number", n).Msg("consumer rebalanced")
		case msg, ok := <-c.Messages():
			if !ok {
				log.Debug().Str("topic", msg.Topic).
					Int32("partition", msg.Partition).
					Int64("offset", msg.Offset).
					Bytes("key", msg.Key).
					Msg("consume not ok")
				return
			}
			deal(msg)
			c.MarkOffset(msg, "")
		}
	}
}
