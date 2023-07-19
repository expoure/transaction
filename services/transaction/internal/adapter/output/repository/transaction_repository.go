package repository

import (
	"context"
	"database/sql"

	"github.com/expoure/pismo/transaction/internal/adapter/output/mapper"
	"github.com/expoure/pismo/transaction/internal/adapter/output/model/entity"
	"github.com/expoure/pismo/transaction/internal/application/domain"
	"github.com/expoure/pismo/transaction/internal/application/port/output"
	"github.com/expoure/pismo/transaction/internal/configuration/database/sqlc"
	"github.com/expoure/pismo/transaction/internal/configuration/logger"
	"github.com/expoure/pismo/transaction/internal/configuration/rest_errors"
	"github.com/google/uuid"

	"go.uber.org/zap"
)

const CREATE_TRANSACTION_RETURNING_SQL = `
	INSERT INTO transaction (
		account_id,
		operation_type_id,
		event_date,
		amount
	) 
	VALUES ($1, $2, $3, ($4, $5)) 
	RETURNING id, account_id, operation_type_id, event_date, amount
`

func NewTransactionRepository(
	queries *sqlc.Queries,
	databaseCon *sql.DB,
) output.TransactionRepository {
	return &transactionRepositoryImpl{
		queries,
		databaseCon,
	}
}

type transactionRepositoryImpl struct {
	queries     *sqlc.Queries
	databaseCon *sql.DB
}

func (tr *transactionRepositoryImpl) CreateTransaction(
	transactionDomain domain.TransactionDomain,
) (*domain.TransactionDomain, *rest_errors.RestErr) {
	logger.Info("Init CreateTransaction repository",
		zap.String("journey", "createTransaction"))

	row := tr.databaseCon.QueryRowContext(
		context.Background(),
		CREATE_TRANSACTION_RETURNING_SQL,
		transactionDomain.AccountID,
		transactionDomain.OperationTypeID,
		transactionDomain.EventDate,
		transactionDomain.Amount.Amount(),
		transactionDomain.Amount.Currency().Code,
	)

	var transaction entity.TransactionEntity

	err := row.Scan(
		&transaction.ID,
		&transaction.AccountID,
		&transaction.OperationTypeID,
		&transaction.EventDate,
		&transaction.Amount,
	)

	if err != nil {
		logger.Error("Error trying to create transaction",
			err,
			zap.String("journey", "createTransaction"))
		return nil, rest_errors.NewInternalServerError(err.Error())
	}

	return mapper.MapEntityToDomain(transaction), nil

}

func (tr *transactionRepositoryImpl) ListTransactionsByAccountID(
	id uuid.UUID,
) (*[]domain.TransactionDomain, *rest_errors.RestErr) {
	logger.Info("Init ListTransactionsByAccountID repository",
		zap.String("journey", "ListTransactionsByAccountID"))

	result, err := tr.queries.ListTransactions(
		context.Background(),
		id,
	)

	if err != nil {
		logger.Error("Error trying to list transactions",
			err,
			zap.String("journey", "ListTransactionsByAccountID"))
		return nil, rest_errors.NewInternalServerError(err.Error())
	}

	transactionsDomain := []domain.TransactionDomain{}

	for _, transaction := range result {
		transactionsDomain = append(transactionsDomain, *mapper.MapEntityToDomain(transaction))
	}

	return &transactionsDomain, nil
}
