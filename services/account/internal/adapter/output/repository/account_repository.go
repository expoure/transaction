package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Rhymond/go-money"
	"github.com/expoure/pismo/account/internal/adapter/output/mapper"
	"github.com/expoure/pismo/account/internal/application/domain"
	"github.com/expoure/pismo/account/internal/application/port/output"
	"github.com/expoure/pismo/account/internal/configuration/customized_errors"
	"github.com/expoure/pismo/account/internal/configuration/database/custom_types"
	"github.com/expoure/pismo/account/internal/configuration/database/sqlc"
	"github.com/expoure/pismo/account/internal/configuration/logger"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"

	"go.uber.org/zap"
)

const UPDATE_BALANCE_RETURNING_SQL = `
	UPDATE account 
	SET balance = ($1, $2), updated_at = NOW() 
	WHERE id = $3 
	RETURNING balance;
`

func NewAccountRepository(
	connPool *pgxpool.Pool,
) output.AccountRepositoryPort {
	return &accountRepositoryImpl{
		queries:  sqlc.New(connPool),
		connPool: connPool,
	}
}

type accountRepositoryImpl struct {
	queries  *sqlc.Queries
	connPool *pgxpool.Pool
}

func (ar *accountRepositoryImpl) CreateAccount(
	accountDomain domain.AccountDomain,
) (*domain.AccountDomain, *error) {
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
		var e *pgconn.PgError
		if errors.As(err, &e) && e.Code == "23505" {
			return nil, &customized_errors.DuplicateKey
		}
		return nil, &err
	}

	return mapper.MapEntityToDomain(result), nil
}

func (ar *accountRepositoryImpl) FindAccountByDocumentNumber(
	documentNumber string,
) (*domain.AccountDomain, *error) {
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
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, &customized_errors.EntityNotFound
		}
		return nil, &err
	}

	return mapper.MapEntityToDomain(result), nil
}

func (ar *accountRepositoryImpl) FindAccountByID(
	id uuid.UUID,
) (*domain.AccountDomain, *error) {
	logger.Info("Init FindAccountByID repository",
		zap.String("journey", "FindAccountByID"))

	result, err := ar.queries.FindAccountById(
		context.Background(),
		id,
	)

	fmt.Println("==============find", result)

	if err != nil {
		logger.Error("Error trying to FindAccountByID",
			err,
			zap.String("journey", "FindAccountByID"))

		if err.Error() == pgx.ErrNoRows.Error() {
			return nil, &customized_errors.EntityNotFound
		}
		return nil, &err
	}

	return mapper.MapEntityToDomain(result), nil
}

func (ar *accountRepositoryImpl) UpdateAccountBalanceByID(
	id uuid.UUID,
	transactionAmount int64,
) (*money.Money, *error) {
	logger.Info("Init UpdateAccountBalance repository",
		zap.String("journey", "UpdateAccountBalance"))

	row := ar.connPool.QueryRow(
		context.Background(),
		UPDATE_BALANCE_RETURNING_SQL,
		transactionAmount,
		"BRL",
		id,
	)

	var balance custom_types.Money

	err := row.Scan(
		&balance,
	)

	if err != nil {
		logger.Error("Error trying to UpdateAccountBalance",
			err,
			zap.String("journey", "UpdateAccountBalance"))
		return nil, &err
	}

	return money.New(balance.Amount, balance.Currency), nil
}

func (ar *accountRepositoryImpl) FindAccountBalanceByID(id uuid.UUID) (*money.Money, *error) {
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
		return nil, &err
	}

	return money.New(result.Amount, result.Currency), nil
}
