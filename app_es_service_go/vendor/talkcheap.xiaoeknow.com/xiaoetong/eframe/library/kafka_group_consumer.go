package library

import (
	"context"
	"errors"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"sync"
)

type KafkaGroupConsumer struct {
	consumerGroup    sarama.ConsumerGroup
	messageHandle    MessageHandleFun
	consumeErrHandle ConsumeErrHandleFunc
	topics           []string
	lock             sync.Mutex
}

func NewKafkaGroupConsumer(config *KafkaGroupConsumerConfig) (kafkaGroupConsumer *KafkaGroupConsumer, err error) {
	if config.ExtraConfig == nil {
		config.ExtraConfig = sarama.NewConfig()
	}

	if config.Name != "" {
		config.ExtraConfig.ClientID = config.Name
	}

	if config.Version == "" {
		config.ExtraConfig.Version = sarama.V2_6_0_0
	}

	if config.Topics == nil || len(config.Topics) == 0 {
		err = errors.New("请指定需要消费的topic 以,号分隔")
		return nil, err
	}

	config.ExtraConfig.Version, err = sarama.ParseKafkaVersion(config.Version)
	if err != nil {
		err = fmt.Errorf("[%s] version string is err: %w", config.Name, err)
		return
	}

	if config.ConsoleDebug == true {
		sarama.Logger = log.New(os.Stdout, "["+config.Name+"]", log.LstdFlags)
	}

	consumerGroup, err := sarama.NewConsumerGroup(config.BrokerAddress, config.GroupName, config.ExtraConfig)
	if err != nil {
		return nil, err
	}

	return &KafkaGroupConsumer{
		consumerGroup: consumerGroup,
		topics:        config.Topics,
	}, nil
}

func NewGroupConsumerHandler(handlerFun MessageHandleFun) *GroupConsumerHandler {
	return &GroupConsumerHandler{
		handleMessage: handlerFun,
	}
}

func (c *KafkaGroupConsumer) Close() error {
	return c.consumerGroup.Close()
}

func (c *KafkaGroupConsumer) SetMessageHandleFunc(f MessageHandleFun) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.messageHandle = f
}

func (c *KafkaGroupConsumer) SetConsumeErrHandleFunc(f ConsumeErrHandleFunc) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.consumeErrHandle = f
}

func (c *KafkaGroupConsumer) StartConsume(ctx context.Context) (err error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.consumeErrHandle == nil {
		err = errors.New("请指定 ConsumeErrHandleFunc 消费异常处理逻辑")
		return
	}

	if c.messageHandle == nil {
		err = errors.New("请指定 MessageHandleFun 消息消费逻辑")
		return
	}

	handler := NewGroupConsumerHandler(c.messageHandle)
	go func() {
		for {
			if err := c.consumerGroup.Consume(ctx, c.topics, handler); err != nil {
				c.consumeErrHandle(err)
				return
			}
			if ctx.Err() != nil {
				return
			}
		}
	}()
	return nil
}

type ConsumeErrHandleFunc func(error)

type MessageHandleFun func(message *sarama.ConsumerMessage)

type GroupConsumerHandler struct {
	handleMessage func(message *sarama.ConsumerMessage)
}

func (h *GroupConsumerHandler) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (h *GroupConsumerHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (h *GroupConsumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case msg, ok := <-claim.Messages():
			if !ok {
				return nil
			}
			h.handleMessage(msg)
			session.MarkMessage(msg, "")
		case <-session.Context().Done():
			return nil
		}
	}
}
