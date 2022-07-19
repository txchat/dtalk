package kafka

import (
	"fmt"
	"strconv"

	cluster "github.com/bsm/sarama-cluster"
	"github.com/txchat/dtalk/service/offline-push/config"
)

func NewKafkaConsumers(appId string, cfg *config.MQSubClient, groupIdx int) map[string]*Consumer {
	store := make(map[string]*Consumer)
	num := int(cfg.Number)
	for i := 0; i < num; i++ {
		store[strconv.Itoa(i)] = &Consumer{Consumer: newKafkaSub(appId, groupIdx, cfg.Brokers)}
	}
	return store
}

func newKafkaSub(appId string, groupIdx int, brokers []string) *cluster.Consumer {
	c := cluster.NewConfig()
	c.Consumer.Return.Errors = true
	c.Group.Return.Notifications = true

	topic := fmt.Sprintf("offpush-%s-topic", appId)
	group := fmt.Sprintf("offpush-%s-group", appId)
	if groupIdx > 0 {
		group = fmt.Sprintf("%s-%d", group, groupIdx)
	}
	consumer, err := cluster.NewConsumer(brokers, group, []string{topic}, c)
	if err != nil {
		panic(err)
	}
	return consumer
}
