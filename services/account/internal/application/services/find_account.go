package service

import (
	"fmt"

	"github.com/expoure/pismo/account/internal/application/domain"
	"github.com/expoure/pismo/account/internal/configuration/logger"
	"github.com/expoure/pismo/account/internal/configuration/rest_errors"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (ad *accountDomainService) FindAccountByIDServices(
	id uuid.UUID,
) (*domain.AccountDomain, *rest_errors.RestErr) {
	logger.Info("Init findAccountByID services.",
		zap.String("journey", "findAccountById"))

	fmt.Println("======", id)
	return ad.repository.FindAccountByID(id)
}

func (ad *accountDomainService) FindAccountByDocumentNumberServices(
	documentNumber string,
) (*domain.AccountDomain, *rest_errors.RestErr) {
	logger.Info("Init findAccountByEmail services.",
		zap.String("journey", "findAccountById"))

	return ad.repository.FindAccountByDocumentNumber(documentNumber)
}
