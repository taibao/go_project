package config

import (
	"app_es_service_go/providers"
	"strings"
	"talkcheap.xiaoeknow.com/xiaoetong/eframe/library"
)

var KafkaSyncProducerConfig []*library.KafkaProducerConfig

func InitKafkaSyncConfigs(env *library.Env) {
	KafkaSyncProducerConfig = []*library.KafkaProducerConfig{
		{
			Receiver:      &providers.KafkaShopMsgSyncProducer,
			Name:          "kafka-shop-msg-sync-producer",
			Version:       env.GetString("KAFKA_FOR_SHOP_MSG_VERSION"),
			BrokerAddress: strings.Split(env.GetString("KAFKA_FOR_SHOP_MSG_BROKER"), ","),
			ConsoleDebug:  true,
		},
	}
}
