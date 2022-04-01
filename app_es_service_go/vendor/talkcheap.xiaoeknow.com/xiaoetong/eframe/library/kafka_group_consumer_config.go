package library

import "github.com/Shopify/sarama"

type KafkaGroupConsumerConfig struct {
	Receiver      **KafkaGroupConsumer
	Name          string
	GroupName     string
	Version       string
	Topics        []string
	BrokerAddress []string
	ConsoleDebug  bool
	ExtraConfig   *sarama.Config
}
