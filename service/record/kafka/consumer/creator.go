package consumer

import (
	"strconv"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

func NewKafkaConsumers(topic, group string, brokers []string, number uint32) map[string]*Consumer {
	store := make(map[string]*Consumer)
	num := int(number)
	for i := 0; i < num; i++ {
		store[strconv.Itoa(i)] = &Consumer{Consumer: newKafkaSub(topic, group, brokers)}
	}
	return store
}

func newKafkaSub(topic, group string, brokers []string) *cluster.Consumer {
	c := cluster.NewConfig()
	c.Consumer.Return.Errors = true
	c.Group.Return.Notifications = true
	c.Version = sarama.V0_11_0_2

	consumer, err := cluster.NewConsumer(brokers, group, []string{topic}, c)
	if err != nil {
		panic(err)
	}
	return consumer
}
