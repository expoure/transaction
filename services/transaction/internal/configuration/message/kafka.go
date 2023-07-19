package message

import (
	"os"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var (
	producer *kafka.Producer
	consumer *kafka.Consumer
	once     sync.Once
)

func initializeKafkaObjects() {
	config := &kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_URL"),
		"group.id":          "transaction-group",
		"auto.offset.reset": "earliest",
	}

	var err error
	producer, err = kafka.NewProducer(config)
	if err != nil {
		panic(err)
	}

	consumer, err = kafka.NewConsumer(config)
	if err != nil {
		panic(err)
	}
}

func GetKafkaProducer() *kafka.Producer {
	once.Do(initializeKafkaObjects)
	return producer
}

func GetKafkaConsumer() *kafka.Consumer {
	once.Do(initializeKafkaObjects)
	return consumer
}

func CloseKafkaConnections() {
	if producer != nil {
		producer.Close()
	}
	if consumer != nil {
		consumer.Close()
	}
}
