package controller

import (
	"net/http"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/expoure/pismo/transaction/internal/adapter/input/mapper"
	"github.com/expoure/pismo/transaction/internal/adapter/input/model/request"
	"github.com/expoure/pismo/transaction/internal/adapter/input/model/response"
	"github.com/expoure/pismo/transaction/internal/application/domain"
	"github.com/expoure/pismo/transaction/internal/application/port/input"
	"github.com/expoure/pismo/transaction/internal/configuration/logger"
	"github.com/expoure/pismo/transaction/internal/configuration/rest_errors"
	"github.com/expoure/pismo/transaction/internal/configuration/validation"
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewTransactionControllerInterface(
	serviceInterface input.TransactionDomainService,
) TransactionControllerInterface {
	return &transactionControllerImpl{
		service: serviceInterface,
	}
}

type TransactionControllerInterface interface {
	ListTransaction(c *gin.Context)
	CreateTransaction(c *gin.Context)
}

type transactionControllerImpl struct {
	service input.TransactionDomainService
}

func (ac *transactionControllerImpl) CreateTransaction(c *gin.Context) {
	logger.Info("Init Transaction controller",
		zap.String("journey", "createTransaction"),
	)
	var transactionRequest request.TransactionRequest

	if err := c.ShouldBindJSON(&transactionRequest); err != nil {
		logger.Error("Error trying to validate transaction info", err,
			zap.String("journey", "createTransaction"))
		errRest := validation.ValidateTransactionError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	accountId, _ := uuid.Parse(transactionRequest.AccountID)
	transactionDomain := domain.TransactionDomain{
		AccountID:       accountId,
		OperationTypeID: transactionRequest.OperationTypeID,
		EventDate:       time.Now(),
		Amount:          money.NewFromFloat(transactionRequest.Amount, money.BRL),
	}

	domainResult, err := ac.service.CreateTransactionServices(transactionDomain)
	if err != nil {
		logger.Error(
			"Error trying to call CreateTransaction service",
			err,
			zap.String("journey", "createTransaction"))
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, mapper.MapDomainToResponse(
		domainResult,
	))
}

func (uc *transactionControllerImpl) ListTransaction(c *gin.Context) {
	logger.Info("Init ListTransaction controller",
		zap.String("journey", "ListTransaction"),
	)

	accountId := c.Query("accountId")
	accountUuid, uuidErr := uuid.Parse(accountId)
	if uuidErr != nil {
		logger.Error("Error trying to validate id",
			uuidErr,
			zap.String("journey", "ListTransaction"),
		)
		errorMessage := rest_errors.NewBadRequestError(
			"id is not a valid uuid",
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	transactionsDomain, err := uc.service.ListTransactionsService(accountUuid)
	if err != nil {
		logger.Error(
			"Error trying to call ListTransactionsService services",
			err,
			zap.String("journey", "ListTransaction"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("ListTransactions controller executed successfully",
		zap.String("journey", "ListTransactions"),
	)

	transactionResponse := []response.TransactionResponse{}

	for _, transaction := range *transactionsDomain {
		transactionResponse = append(transactionResponse, mapper.MapDomainToResponse(&transaction))
	}

	c.JSON(http.StatusOK, transactionResponse)
}
