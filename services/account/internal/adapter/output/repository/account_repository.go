package repository

import (
	"context"
	"database/sql"

	"github.com/Rhymond/go-money"
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
	databaseCon *sql.DB,
) output.AccountPort {
	return &accountRepositoryImpl{
		queries,
		databaseCon,
	}
}

type accountRepositoryImpl struct {
	queries     *sqlc.Queries
	databaseCon *sql.DB
}

func (ar *accountRepositoryImpl) CreateAccount(
	accountDomain domain.AccountDomain,
) (*domain.AccountDomain, *rest_errors.RestErr) {
	logger.Info("Init createAccount repository",
		zap.String("journey", "createAccount"))

	result, err := ar.queries.CreateAccount(
		context.Background(),
		accountDomain.DocumentNumber,
	)

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

func (ar *accountRepositoryImpl) UpdateAccountBalanceByID(
	id uuid.UUID,
	transactionAmount int64,
) *rest_errors.RestErr {
	logger.Info("Init UpdateAccountBalance repository",
		zap.String("journey", "UpdateAccountBalance"))

	// tenho que colocar um mutex aqui
	_, err := ar.databaseCon.Exec("UPDATE account SET balance = ($1, $2), updated_at = NOW() WHERE id = $3", transactionAmount, "BRL", id)

	if err != nil {
		logger.Error("Error trying to FindAccountByID",
			err,
			zap.String("journey", "FindAccountByID"))
		return rest_errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (ar *accountRepositoryImpl) FindAccountBalanceByID(id uuid.UUID) (*money.Money, *rest_errors.RestErr) {
	logger.Info("Init FindAccountBalanceByID repository",
		zap.String("journey", "FindAccountBalanceByID"))

	result, err := ar.queries.FindAccountBalanceById(
		context.Background(),
		id,
	)

	if err != nil {
		logger.Error("Error trying to FindAccountBalanceByID",
			err,
			zap.String("journey", "FindAccountBalanceByID"))
		return nil, rest_errors.NewInternalServerError(err.Error())
	}

	return money.New(result.Amount, result.Currency), nil
}
