package job

import (
	"context"
	"app_es_service_go/providers"
	"fmt"
	"github.com/Shopify/sarama"
	"os"
)

func StartShopMsgConsume(ctx context.Context) {
	providers.KafkaShopMsgGroupConsumer.SetMessageHandleFunc(handleShopInfoMsg)
	providers.KafkaShopMsgGroupConsumer.SetConsumeErrHandleFunc(handleConsumeError)
	err := providers.KafkaShopMsgGroupConsumer.StartConsume(ctx)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stdout, "%s", err.Error())
	}
}

func handleShopInfoMsg(msg *sarama.ConsumerMessage) {
	fmt.Println("consume msg:", msg.Value)
}

func handleConsumeError(err error) {
	fmt.Println("consume err:", err)
}
