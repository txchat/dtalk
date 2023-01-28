package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/inconshreveable/log15"
	xkafka "github.com/txchat/dtalk/pkg/mq/kafka"
)

var (
	topic, group          string
	broker                string
	Consumers             int
	Processors            int
	connTimeout           int
	batchCacheCapacity    int
	consumerCacheCapacity int
	graceful              bool
)

func init() {
	flag.IntVar(&Consumers, "cons", 1, "装填到缓冲区的工作协程数量")
	flag.IntVar(&Processors, "pros", 1, "从缓冲区取出的工作协程数量")
	flag.IntVar(&connTimeout, "cto", 7000, "连接超时时间，单位ms")
	flag.IntVar(&batchCacheCapacity, "bcc", 0, "批量消费缓冲区大小")
	flag.IntVar(&consumerCacheCapacity, "ccc", 0, "单个消费者缓冲区大小")
	flag.StringVar(&topic, "topic", "test-mytest-topic", "")
	flag.StringVar(&group, "group", "", "")
	flag.StringVar(&broker, "broker", "127.0.0.1:9092", "")
	flag.BoolVar(&graceful, "graceful", true, "优雅停止开关")
}

//
func main() {
	flag.Parse()
	log.Info("service start", "broker", broker, "topic", topic, "group", group)
	log.Info("configs",
		"Consumers", Consumers, "Processors", Processors,
		"connTimeout", connTimeout,
		"batchCacheCapacity", batchCacheCapacity,
		"consumerCacheCapacity", consumerCacheCapacity,
		"graceful", graceful,
	)

	consumer := xkafka.NewConsumer(xkafka.ConsumerConfig{
		Version:        "",
		Brokers:        []string{broker},
		Group:          group,
		Topic:          topic,
		CacheCapacity:  consumerCacheCapacity,
		ConnectTimeout: time.Millisecond * time.Duration(connTimeout),
	}, nil)
	log.Info("dial kafka broker success")
	bc := xkafka.NewBatchConsumer(xkafka.BatchConsumerConf{
		CacheCapacity: batchCacheCapacity,
		Consumers:     Consumers,
		Processors:    Processors,
	}, xkafka.WithHandle(func(key string, data []byte) error {
		log.Info("receive msg:", "value", data)
		time.Sleep(time.Millisecond * 500)
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
			if graceful {
				bc.GracefulStop(context.Background())
			} else {
				bc.Stop()
			}
			time.Sleep(time.Second * 2)
			return
		case syscall.SIGHUP:
			// TODO reload
		default:
			return
		}
	}
}
