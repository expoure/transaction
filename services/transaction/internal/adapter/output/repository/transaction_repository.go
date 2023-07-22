package repository

import (
	"context"

	"github.com/expoure/pismo/transaction/internal/adapter/output/mapper"
	"github.com/expoure/pismo/transaction/internal/adapter/output/model/entity"
	"github.com/expoure/pismo/transaction/internal/application/domain"
	"github.com/expoure/pismo/transaction/internal/application/port/output"
	"github.com/expoure/pismo/transaction/internal/configuration/customized_errors"
	"github.com/expoure/pismo/transaction/internal/configuration/database/sqlc"
	"github.com/expoure/pismo/transaction/internal/configuration/logger"
	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"

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
	connPool *pgxpool.Pool,
) output.TransactionRepository {
	return &transactionRepositoryImpl{
		queries:  sqlc.New(connPool),
		connPool: connPool,
	}
}

type transactionRepositoryImpl struct {
	queries  *sqlc.Queries
	connPool *pgxpool.Pool
}

func (tr *transactionRepositoryImpl) CreateTransaction(
	transactionDomain domain.TransactionDomain,
) (*domain.TransactionDomain, *error) {
	logger.Info("Init CreateTransaction repository",
		zap.String("journey", "createTransaction"))

	row := tr.connPool.QueryRow(
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
		return nil, &err
	}

	return mapper.MapEntityToDomain(transaction), nil

}

func (tr *transactionRepositoryImpl) ListTransactionsByAccountID(
	id uuid.UUID,
) (*[]domain.TransactionDomain, *error) {
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
		if err.Error() == pgx.ErrNoRows.Error() {
			return nil, &customized_errors.EntityNotFound
		}
		return nil, &err
	}

	transactionsDomain := []domain.TransactionDomain{}

	for _, transaction := range result {
		transactionsDomain = append(transactionsDomain, *mapper.MapEntityToDomain(transaction))
	}

	return &transactionsDomain, nil
}
