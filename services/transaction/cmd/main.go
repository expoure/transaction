package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/expoure/pismo/transaction/internal/adapter/input/controller"
	"github.com/expoure/pismo/transaction/internal/adapter/input/controller/routes"
	"github.com/expoure/pismo/transaction/internal/adapter/output/repository"
	service "github.com/expoure/pismo/transaction/internal/application/services"
	"github.com/expoure/pismo/transaction/internal/configuration/database/postgres"
	"github.com/expoure/pismo/transaction/internal/configuration/database/sqlc"
	"github.com/expoure/pismo/transaction/internal/configuration/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Starting transaction microservice")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	databaseCon, err := postgres.NewPostgresConnection(context.Background())
	if err != nil {
		log.Fatalf(
			"Error trying to connect to database, error=%s \n",
			err.Error())
		return
	}
	defer databaseCon.Close()

	transactionController := initDependencies(databaseCon)

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
	databaseConn *sql.DB,
) controller.transactionControllerInterface {
	queries := sqlc.New(databaseConn)
	transactionRepo := repository.NewTransactionRepository(queries, databaseConn)
	transactionService := service.NewTransactionDomainService(transactionRepo)
	return controller.NewtransactionControllerInterface(transactionService)
}
