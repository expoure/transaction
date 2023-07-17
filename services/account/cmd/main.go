package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/expoure/pismo/account/internal/adapter/input/controller"
	"github.com/expoure/pismo/account/internal/adapter/input/controller/routes"
	"github.com/expoure/pismo/account/internal/adapter/output/repository"
	service "github.com/expoure/pismo/account/internal/application/services"
	"github.com/expoure/pismo/account/internal/configuration/database/postgres"
	"github.com/expoure/pismo/account/internal/configuration/database/sqlc"
	"github.com/expoure/pismo/account/internal/configuration/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Starting account microservice")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	database, err := postgres.NewPostgresConnection(context.Background())
	if err != nil {
		log.Fatalf(
			"Error trying to connect to database, error=%s \n",
			err.Error())
		return
	}
	defer database.Close()

	queries := sqlc.New(database)

	accountController := initDependencies(queries)

	router := gin.Default()
	group := router.Group("/v1/accounts")
	routes.InitAccountRoutes(group, accountController)

	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if err := router.Run(addr); err != nil {
		log.Fatal(err)
	}

	// aqui eu devo inicializar o kafka (ou talvez no configuration?)
	// e no adapter vai estar apenas o consumer em si
}

func initDependencies(
	queries *sqlc.Queries,
) controller.AccountControllerInterface {
	accountRepo := repository.NewAccountRepository(queries)
	accountService := service.NewAccountDomainService(accountRepo)
	return controller.NewAccountControllerInterface(accountService)
}
