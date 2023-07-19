package controller

import (
	"net/http"

	"github.com/expoure/pismo/transaction/internal/adapter/input/mapper"
	"github.com/expoure/pismo/transaction/internal/adapter/input/model/request"
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
	serviceInterface input.AccountDomainService,
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
	service input.AccountDomainService
}

func (ac *transactionControllerImpl) CreateTransaction(c *gin.Context) {
	logger.Info("Init Account controller",
		zap.String("journey", "createAccount"),
	)
	var accountRequest request.AccountRequest

	if err := c.ShouldBindJSON(&accountRequest); err != nil {
		logger.Error("Error trying to validate account info", err,
			zap.String("journey", "createAccount"))
		errRest := validation.ValidateAccountError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	accountDomain := domain.TransactionDomain{
		DocumentNumber: accountRequest.DocumentNumber,
	}

	domainResult, err := ac.service.CreateAccountServices(accountDomain)
	if err != nil {
		logger.Error(
			"Error trying to call CreateAccount service",
			err,
			zap.String("journey", "createAccount"))
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, mapper.MapDomainToResponse(
		domainResult,
	))
}

func (uc *transactionControllerImpl) ListTransaction(c *gin.Context) {
	logger.Info("Init findAccountByID controller",
		zap.String("journey", "findAccountByID"),
	)

	accountId := c.Param("id")

	accountUuid, uuidErr := uuid.Parse(accountId)
	if uuidErr != nil {
		logger.Error("Error trying to validate id",
			uuidErr,
			zap.String("journey", "findAccountByID"),
		)
		errorMessage := rest_errors.NewBadRequestError(
			"id is not a valid uuid",
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	accountDomain, err := uc.service.FindAccountByIDServices(accountUuid)
	if err != nil {
		logger.Error("Error trying to call findAccountByID services",
			err,
			zap.String("journey", "findAccountByID"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindAccountByID controller executed successfully",
		zap.String("journey", "findAccountByID"),
	)
	c.JSON(http.StatusOK, mapper.MapDomainToResponse(
		accountDomain,
	))
}
