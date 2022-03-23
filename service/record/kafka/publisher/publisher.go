package publisher

import (
	kafka "gopkg.in/Shopify/sarama.v1"
)

func NewKafkaPub(brokers []string) kafka.SyncProducer {
	kc := kafka.NewConfig()
	kc.Producer.RequiredAcks = kafka.WaitForAll // Wait for all in-sync replicas to ack the message
	kc.Producer.Retry.Max = 10                  // Retry up to 10 times to produce the message
	kc.Producer.Return.Successes = true
	pub, err := kafka.NewSyncProducer(brokers, kc)
	if err != nil {
		panic(err)
	}
	return pub
}
