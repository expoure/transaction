package service

import (
	"github.com/expoure/pismo/transaction/internal/application/domain"
	"github.com/expoure/pismo/transaction/internal/configuration/logger"
	"github.com/expoure/pismo/transaction/internal/configuration/rest_errors"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (td *transactionDomainService) ListTransactionsService(
	accoountId uuid.UUID,
) (*[]domain.TransactionDomain, *rest_errors.RestErr) {
	logger.Info("Init listTransaction.",
		zap.String("journey", "createTransaction"))

	transactions, err := td.repository.ListTransactionsByAccountID(accoountId)

	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "createTransaction"))
		return nil, err
	}
	// criar evento de transaction_created

	logger.Info(
		"ListTransactions service executed successfully",
		zap.String("journey", "listTransaction"))

	return transactions, nil
}
