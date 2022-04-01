package config

import (
	"app_es_service_go/providers"
	"strings"
	"talkcheap.xiaoeknow.com/xiaoetong/eframe/library"
)

var KafkaGroupConsumerConfig []*library.KafkaGroupConsumerConfig

func InitKafkaGroupConsumerConfigs(env *library.Env) {
	KafkaGroupConsumerConfig = []*library.KafkaGroupConsumerConfig{
		{
			Receiver:      &providers.KafkaShopMsgGroupConsumer,
			Name:          "kafka-shop-msg-group-consumer",
			GroupName:     env.GetString("KAFKA_FOR_SHOP_MSG_GROUP_NAME"),
			Version:       env.GetString("KAFKA_FOR_SHOP_MSG_VERSION"),
			BrokerAddress: strings.Split(env.GetString("KAFKA_FOR_SHOP_MSG_BROKER"), ","),
			Topics:        strings.Split(env.GetString("KAFKA_FOR_SHOP_MSG_CONSUMER_TOPICS"), ","),
			ConsoleDebug:  false,
		},
	}
}
