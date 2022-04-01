package library

import "github.com/Shopify/sarama"

type KafkaProducerConfig struct {
	Receiver      **KafkaSyncProducer
	Name          string
	Version       string
	BrokerAddress []string
	ConsoleDebug  bool
	ExtraConfig   *sarama.Config
}
