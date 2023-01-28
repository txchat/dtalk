package kafka

import (
	"errors"
	"strconv"
	"time"

	"github.com/Shopify/sarama"
	log "github.com/inconshreveable/log15"
)

type ProducerConfig struct {
	Version string   `json:",optional"`
	Brokers []string `json:",optional"`

	realVersion sarama.KafkaVersion `json:"-,optional"`
}

type Producer struct {
	brokers []string
	conn    sarama.SyncProducer
}

func ensureProducerConfig(cfg *ProducerConfig) {
	//version
	version, err := sarama.ParseKafkaVersion(cfg.Version)
	if err != nil {
		cfg.realVersion = sarama.V2_0_0_0
	}
	cfg.realVersion = version

	if len(cfg.Brokers) == 0 {
		panic(errors.New("brokers is empty"))
	}
}

func NewProducer(cfg ProducerConfig) *Producer {
	ensureProducerConfig(&cfg)
	kc := sarama.NewConfig()
	kc.Version = cfg.realVersion //当前kafka的版本
	//kc.Producer.Compression = sarama.CompressionSnappy //将数据进行压缩传输，提高数据传输的效率
	kc.Producer.RequiredAcks = sarama.WaitForAll //等待所有同步中的副本都确认消息
	kc.Producer.Retry.Max = 10                   //发送消息重试的次数
	kc.Producer.Return.Successes = true

	pub, err := sarama.NewSyncProducer(cfg.Brokers, kc)
	if err != nil {
		panic(err)
	}
	log.Info("producer dial kafka server", "brokers", cfg.Brokers)
	return &Producer{
		brokers: cfg.Brokers,
		conn:    pub,
	}
}

func (p *Producer) Publish(topic, k string, v []byte) (int32, int64, error) {
	if k == "" {
		k = strconv.FormatInt(time.Now().UnixNano(), 10)
	}
	log.Debug("push params", "topic", topic, "brokers", p.brokers)
	m := &sarama.ProducerMessage{
		Key:       sarama.StringEncoder(k),
		Topic:     topic,
		Value:     sarama.ByteEncoder(v),
		Timestamp: time.Now(),
	}
	partition, offset, err := p.conn.SendMessage(m)
	return partition, offset, err
}

func (p *Producer) Close() error {
	return p.conn.Close()
}
