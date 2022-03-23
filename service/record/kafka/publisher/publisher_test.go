package publisher

import (
	"github.com/google/uuid"
	"testing"

	"gopkg.in/Shopify/sarama.v1"
)

func TestNewKafkaPub(t *testing.T) {
	pub := NewKafkaPub([]string{"127.0.0.1:9092"})
	appTopic := "test-topic"

	data := "1"
	m := &sarama.ProducerMessage{
		Key:   sarama.StringEncoder(uuid.New().String()),
		Topic: appTopic,
		Value: sarama.ByteEncoder(data),
	}
	if _, _, err := pub.SendMessage(m); err != nil {
		t.Error("SendMessage", err)
	}
	t.Log("success")
}
