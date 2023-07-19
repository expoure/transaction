package repository

import (
	"context"
	"database/sql"

	"github.com/expoure/pismo/transaction/internal/adapter/output/mapper"
	"github.com/expoure/pismo/transaction/internal/application/domain"
	"github.com/expoure/pismo/transaction/internal/application/port/output"
	"github.com/expoure/pismo/transaction/internal/configuration/database/custom_types"
	"github.com/expoure/pismo/transaction/internal/configuration/database/sqlc"
	"github.com/expoure/pismo/transaction/internal/configuration/logger"
	"github.com/expoure/pismo/transaction/internal/configuration/rest_errors"
	"github.com/google/uuid"

	"go.uber.org/zap"
)

func NewTransactionRepository(
	queries *sqlc.Queries,
	databaseCon *sql.DB,
) output.TransactionPort {
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
	accountDomain domain.TransactionDomain,
) (*domain.TransactionDomain, *rest_errors.RestErr) {
	logger.Info("Init CreateTransaction repository",
		zap.String("journey", "createTransaction"))

	// usar driver pg
	result, err := tr.queries.CreateTransaction(
		context.Background(),
		sqlc.CreateTransactionParams{
			AccountID:       accountDomain.ID,
			OperationTypeID: sql.NullInt32{Int32: accountDomain.OperationTypeID, Valid: true},
			EventDate:       accountDomain.EventDate,
			Amount: &custom_types.Money{
				Amount:   accountDomain.Amount.Amount(),
				Currency: accountDomain.Amount.Currency().Code,
			},
		},
	)
	if err != nil {
		logger.Error("Error trying to create transaction",
			err,
			zap.String("journey", "createTransaction"))
		return nil, rest_errors.NewInternalServerError(err.Error())
	}

	return mapper.MapEntityToDomain(result), nil

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
}
