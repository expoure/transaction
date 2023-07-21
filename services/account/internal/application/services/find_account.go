package service

import (
	"errors"

	"github.com/expoure/pismo/account/internal/application/domain"
	"github.com/expoure/pismo/account/internal/configuration/customized_errors"
	"github.com/expoure/pismo/account/internal/configuration/logger"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (ad *accountDomainService) FindAccountByIDServices(
	id uuid.UUID,
) (*domain.AccountDomain, *customized_errors.RestErr) {
	logger.Info("Init findAccountByID services.",
		zap.String("journey", "findAccountById"))

	account, err := ad.repository.FindAccountByID(id)

	if err != nil {
		return handleFindError(err)
	}

	return account, nil
}

func (ad *accountDomainService) FindAccountByDocumentNumberServices(
	documentNumber string,
) (*domain.AccountDomain, *customized_errors.RestErr) {
	logger.Info("Init findAccountByEmail services.",
		zap.String("journey", "findAccountById"))

	account, err := ad.repository.FindAccountByDocumentNumber(documentNumber)
	if err != nil {
		return handleFindError(err)
	}

	return account, nil
}

func handleFindError(err *error) (*domain.AccountDomain, *customized_errors.RestErr) {
	if errors.Is(*err, customized_errors.EntityNotFound) {
		return nil, customized_errors.NewBadRequestError("Account not found")
	}
	return nil, customized_errors.NewInternalServerError("Internal Server Error")
}
