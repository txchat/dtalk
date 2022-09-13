package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/inconshreveable/log15"
	xkafka "github.com/txchat/dtalk/pkg/mq/kafka"
)

var (
	topic, group string
	broker       string
	Consumers    int
	Processors   int
)

func init() {
	flag.IntVar(&Consumers, "cons", 1, "")
	flag.IntVar(&Processors, "pros", 1, "")
	flag.StringVar(&topic, "topic", "test-mytest-topic", "")
	flag.StringVar(&group, "group", "", "")
	flag.StringVar(&broker, "broker", "127.0.0.1:9092", "")
}

//
func main() {
	flag.Parse()
	log.Info("service start", "broker", broker, "topic", topic, "group", group, "Consumers", Consumers, "Processors", Processors)

	consumer := xkafka.NewConsumer(xkafka.ConsumerConfig{
		Version:        "",
		Brokers:        []string{broker},
		Group:          group,
		Topic:          topic,
		ConnectTimeout: time.Second * 20,
	}, nil)
	log.Info("dial kafka broker success")
	bc := xkafka.NewBatchConsumer(xkafka.BatchConsumerConf{
		CacheCapacity: 0,
		Consumers:     Consumers,
		Processors:    Processors,
	}, xkafka.WithHandle(func(key, value string) error {
		log.Info("receive msg:", "value", value)
		time.Sleep(time.Millisecond * 300)
		return nil
	}), consumer)

	go func() {
		bc.Start()
	}()

	// init signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			bc.Stop()
			time.Sleep(time.Second * 2)
			return
		case syscall.SIGHUP:
			// TODO reload
		default:
			return
		}
	}
}
