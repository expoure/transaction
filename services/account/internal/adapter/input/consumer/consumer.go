package consumer

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/expoure/pismo/account/internal/adapter/input/mapper"
	"github.com/expoure/pismo/account/internal/application/port/input"
	"github.com/expoure/pismo/account/internal/configuration/logger"
	"go.uber.org/zap"
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
	Consume()
	Subscribe(topics []string)
}

type consumerImpl struct {
	kafka   *kafka.Consumer
	service input.AccountDomainService
}

func (c *consumerImpl) Subscribe(topics []string) {
	c.kafka.SubscribeTopics(topics, nil)
}

func (c *consumerImpl) Consume() {
	for {
		msg, err := c.kafka.ReadMessage(-1)
		if err == nil {
			logger.Info("Init Account consumer",
				zap.String("journey", "Consume"),
			)

			switch *msg.TopicPartition.Topic {
			case TRANSACTION_CREATED_TOPIC:

				var transaction = mapper.MapMessageToTransaction(msg.Value)
				amount, _ := transaction.Amount.Int64()

				c.service.UpdateAccountBalanceByIDServices(
					transaction.AccountID,
					amount,
				)
			}

		} else {
			logger.Error(
				"Error trying to consume message",
				err,
				zap.String("journey", "Consume"),
			)
		}
	}
}
