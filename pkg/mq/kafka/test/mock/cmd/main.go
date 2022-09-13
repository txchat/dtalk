package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/txchat/dtalk/pkg/mq/kafka/test/mock"

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

func main() {
	flag.Parse()
	log.Info("service start", "broker", broker, "topic", topic, "group", group, "Consumers", Consumers, "Processors", Processors)

	consumer := mock.NewConsumer(time.Millisecond * 500)
	log.Info("dial kafka broker success")
	bc := xkafka.NewBatchConsumer(xkafka.BatchConsumerConf{
		CacheCapacity: 0,
		Consumers:     Consumers,
		Processors:    Processors,
	}, xkafka.WithHandle(func(key string, data []byte) error {
		log.Info("receive msg:", "value", data)
		return nil
	}), consumer)

	bc.Start()

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
