package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	log "github.com/inconshreveable/log15"
)

var (
	group string
)

func init() {
	flag.StringVar(&group, "g", "test-group", "")
}

func main() {
	flag.Parse()
	fmt.Println(group)
	go Listen(newKafkaSub())
	// init signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			time.Sleep(time.Second * 2)
			return
		case syscall.SIGHUP:
			// TODO reload
		default:
			return
		}
	}
}

func deal(msg *sarama.ConsumerMessage) error {
	fmt.Println(msg)
	fmt.Println(msg.Value)
	return nil
}

func Listen(c *cluster.Consumer) {
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
			deal(msg)
			c.MarkOffset(msg, "")
		}
	}
}

func newKafkaSub() *cluster.Consumer {
	c := cluster.NewConfig()
	c.Consumer.Return.Errors = true
	c.Group.Return.Notifications = true

	topic := "test-topic"
	consumer, err := cluster.NewConsumer([]string{"127.0.0.1:9092"}, group, []string{topic}, c)
	if err != nil {
		panic(err)
	}
	return consumer
}
