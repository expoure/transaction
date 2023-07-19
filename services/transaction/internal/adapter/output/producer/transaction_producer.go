package producer

import (
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/expoure/pismo/transaction/internal/application/domain"
	"github.com/expoure/pismo/transaction/internal/application/port/output"
)

var TRANSACTION_CREATED_TOPIC string = "transaction_created"

func NewTransactionProducer(
	kafka *kafka.Producer,
) output.TransactionProducer {
	return &transactionProducerImpl{
		kafka,
	}
}

type transactionProducerImpl struct {
	kafka *kafka.Producer
}

func (tp *transactionProducerImpl) TransactionCreated(transactionDomain domain.TransactionDomain) {
	transactionJson, _ := json.Marshal(transactionDomain)

	tp.kafka.Produce(
		&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &TRANSACTION_CREATED_TOPIC,
				Partition: kafka.PartitionAny,
			},
			Value: []byte(transactionJson),
		},
		nil,
	)
}
