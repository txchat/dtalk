package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	cluster "github.com/bsm/sarama-cluster"
	"github.com/inconshreveable/log15"
	"github.com/txchat/dtalk/service/offline-push/service/kafka"
	logic "github.com/txchat/im/api/logic/grpc"
)

var log = log15.New("test main", "model", "kafka consume")

func main() {
	store := newKafkaConsumers()
	p := &process{}
	for i, c := range store {
		log.Debug(fmt.Sprintf("accept %v", i))
		go c.Listen(p)
	}

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

func newKafkaSub() *cluster.Consumer {
	c := cluster.NewConfig()
	c.Consumer.Return.Errors = true
	c.Group.Return.Notifications = true

	topic := fmt.Sprintf("goim-%s-topic", "dtalk")
	group := fmt.Sprintf("goim-%s-group", "dtalk")
	consumer, err := cluster.NewConsumer([]string{"127.0.0.1:9092"}, group, []string{topic}, c)
	if err != nil {
		panic(err)
	}
	return consumer
}

func newKafkaConsumers() map[string]*kafka.Consumer {
	store := make(map[string]*kafka.Consumer)
	num := 1
	for i := 0; i < num; i++ {
		store[strconv.Itoa(i)] = &kafka.Consumer{Consumer: newKafkaSub()}
	}
	return store
}

type process struct {
}

func (p *process) Deal(m *logic.BizMsg) error {
	fmt.Println(m)
	return nil
}
