package consumer

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/expoure/pismo/account/internal/adapter/input/mapper"
	"github.com/expoure/pismo/account/internal/application/port/input"
)

func NewConsumerInterface(
	kafka *kafka.Consumer,
	service input.AccountDomainService,
) Consumer {
	return &consumerImpl{
		kafka:   kafka,
		service: service,
	}
}

type Consumer interface {
	Consume(topics []string)
}

type consumerImpl struct {
	kafka   *kafka.Consumer
	service input.AccountDomainService
}

func (c *consumerImpl) Consume(topics []string) {
	c.kafka.SubscribeTopics(topics, nil)
	for {
		msg, err := c.kafka.ReadMessage(-1)
		if err == nil {

			var transaction = mapper.MapMessageToTransaction(msg.Value)
			amount, _ := transaction.Amount.Int64()
			c.service.UpdateAccountBalanceByIDServices(
				transaction.AccountID,
				amount,
			)
		} else {
			// logger.Error(
			// 	"Error trying to consume message",
			// 	zap.Error(err),
			// )
		}
	}
}
