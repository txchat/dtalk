package kafka

import (
	cluster "github.com/bsm/sarama-cluster"
	"github.com/golang/protobuf/proto"
	"github.com/inconshreveable/log15"
	offlinepush "github.com/txchat/dtalk/service/offline-push/api"
	"github.com/txchat/dtalk/service/offline-push/model"
)

var log = log15.New("model", "offputhtest")

type Process interface {
	Deal(m *offlinepush.OffPushMsg) error
}

type Consumer struct {
	*cluster.Consumer
}

func (c *Consumer) Listen(p Process) {
	for {
		select {
		case err := <-c.Errors():
			log.Error("consumer error", "err", err)
		case n := <-c.Notifications():
			log.Info("consumer rebalanced", "number", n)
		case msg, ok := <-c.Messages():
			if !ok {
				log.Debug("consume not ok", "topic", msg.Topic, "partition", msg.Partition, "offset", msg.Offset, "key", msg.Key)
				return
			}
			bizMsg := new(offlinepush.OffPushMsg)
			if err := proto.Unmarshal(msg.Value, bizMsg); err != nil {
				log.Error("proto.Unmarshal error", "err", err, "msg", msg)
				continue
			}
			log.Debug("consume process", "topic", msg.Topic, "partition", msg.Partition, "offset", msg.Offset, "key", msg.Key, "bizMsg", bizMsg)
			err := p.Deal(bizMsg)
			if err != nil {
				log.Debug("p.Deal error", "err", err, "bizMsg", bizMsg)
				if err == model.ErrConsumeRedo {
					//TODO redo consume message
				}
			}
			c.MarkOffset(msg, "")
		}
	}
}
