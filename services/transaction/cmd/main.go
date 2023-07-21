package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/expoure/pismo/transaction/internal/adapter/input/controller"
	"github.com/expoure/pismo/transaction/internal/adapter/input/controller/routes"
	"github.com/expoure/pismo/transaction/internal/adapter/output/producer"
	"github.com/expoure/pismo/transaction/internal/adapter/output/repository"
	service "github.com/expoure/pismo/transaction/internal/application/services"
	"github.com/expoure/pismo/transaction/internal/configuration/database/postgres"
	"github.com/expoure/pismo/transaction/internal/configuration/logger"
	"github.com/expoure/pismo/transaction/internal/configuration/message_broker"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Starting transaction microservice")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	connPool, err := postgres.NewPostgresConnection(context.Background())
	if err != nil {
		log.Fatalf(
			"Error trying to connect to database, error=%s \n",
			err.Error())
		return
	}
	defer connPool.Close()

	producer := message_broker.GetKafkaProducer()
	defer message_broker.CloseKafkaConnections()

	transactionController := initDependencies(connPool, producer)

	router := gin.Default()
	group := router.Group("/v1/transactions")
	routes.InitTransactionRoutes(group, transactionController)

	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if err := router.Run(addr); err != nil {
		log.Fatal(err)
	}

	// aqui eu devo inicializar o kafka (ou talvez no configuration?)
	// e no adapter vai estar apenas o consumer em si
}

func initDependencies(
	connPool *pgxpool.Pool,
	messageProducer *kafka.Producer,
) controller.TransactionControllerInterface {
	transactionRepo := repository.NewTransactionRepository(connPool)
	transactionProducer := producer.NewTransactionProducer(messageProducer)
	transactionService := service.NewTransactionDomainService(transactionRepo, transactionProducer)
	return controller.NewTransactionControllerInterface(transactionService)
}
