package service

import (
	"github.com/expoure/pismo/account/internal/application/domain"
	"github.com/expoure/pismo/account/internal/configuration/logger"
	"github.com/expoure/pismo/account/internal/configuration/rest_errors"
	"go.uber.org/zap"
)

func (ad *accountDomainService) CreateAccountServices(
	accountDomain domain.AccountDomain,
) (*domain.AccountDomain, *rest_errors.RestErr) {
	logger.Info("Init createAccount model.",
		zap.String("journey", "createAccount"))

	account, _ := ad.FindAccountByDocumentNumberServices(accountDomain.DocumentNumber)

	if account != nil {
		return nil, rest_errors.NewBadRequestError("This document number is already registered")
	}

	accountCreated, err := ad.repository.CreateAccount(accountDomain)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "createAccount"))
		return nil, err
	}
	// criar evento de account_created

	logger.Info(
		"CreateAccount service executed successfully",
		zap.String("AccountId", accountCreated.ID.String()),
		zap.String("journey", "createAccount"))
	return accountCreated, nil
}
