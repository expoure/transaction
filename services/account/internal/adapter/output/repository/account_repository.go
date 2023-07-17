package repository

import (
	"context"

	"github.com/expoure/pismo/account/internal/adapter/output/mapper"
	"github.com/expoure/pismo/account/internal/application/domain"
	"github.com/expoure/pismo/account/internal/application/port/output"
	"github.com/expoure/pismo/account/internal/configuration/database/sqlc"
	"github.com/expoure/pismo/account/internal/configuration/logger"
	"github.com/expoure/pismo/account/internal/configuration/rest_errors"
	"github.com/google/uuid"

	"go.uber.org/zap"
)

func NewAccountRepository(
	queries *sqlc.Queries,
) output.AccountPort {
	return &accountRepositoryImpl{
		queries,
	}
}

type accountRepositoryImpl struct {
	queries *sqlc.Queries
}

func (ar *accountRepositoryImpl) CreateAccount(
	accountDomain domain.AccountDomain,
) (*domain.AccountDomain, *rest_errors.RestErr) {
	logger.Info("Init createAccount repository",
		zap.String("journey", "createAccount"))

	result, err := ar.queries.CreateAccount(context.Background(), accountDomain.DocumentNumber)
	if err != nil {
		logger.Error("Error trying to create account",
			err,
			zap.String("journey", "createAccount"))
		return nil, rest_errors.NewInternalServerError(err.Error())
	}

	return mapper.MapEntityToDomain(result), nil
}

func (ar *accountRepositoryImpl) FindAccountByDocumentNumber(
	documentNumber string,
) (*domain.AccountDomain, *rest_errors.RestErr) {
	logger.Info("Init FindAccountByDocumentNumber repository",
		zap.String("journey", "FindAccountByDocumentNumber"))

	result, err := ar.queries.FindAccountByDocumentNumber(
		context.Background(),
		documentNumber,
	)

	if err != nil {
		logger.Error("Error trying to FindAccountByDocumentNumber",
			err,
			zap.String("journey", "FindAccountByDocumentNumber"))
		return nil, rest_errors.NewInternalServerError(err.Error())
	}

	return mapper.MapEntityToDomain(result), nil
}

func (ar *accountRepositoryImpl) FindAccountByID(
	id uuid.UUID,
) (*domain.AccountDomain, *rest_errors.RestErr) {
	logger.Info("Init FindAccountByID repository",
		zap.String("journey", "FindAccountByID"))

	result, err := ar.queries.FindAccountById(
		context.Background(),
		id,
	)

	if err != nil {
		logger.Error("Error trying to FindAccountByID",
			err,
			zap.String("journey", "FindAccountByID"))
		return nil, rest_errors.NewInternalServerError(err.Error())
	}

	return mapper.MapEntityToDomain(result), nil
}
