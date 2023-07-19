package service

import (
	"github.com/expoure/pismo/transaction/internal/application/domain"
	"github.com/expoure/pismo/transaction/internal/configuration/logger"
	"github.com/expoure/pismo/transaction/internal/configuration/rest_errors"
	"go.uber.org/zap"
)

func (td *transactionDomainService) CreateTransactionServices(
	transactionDomain domain.TransactionDomain,
) (*domain.TransactionDomain, *rest_errors.RestErr) {
	logger.Info("Init createTransaction model.",
		zap.String("journey", "createTransaction"))

	transaction, err := td.repository.CreateTransaction(transactionDomain)

	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "createTransaction"))
		return nil, err
	}
	// criar evento de transaction_created

	logger.Info(
		"CreateTransaction service executed successfully",
		zap.String("TransactionId", transaction.ID.String()),
		zap.String("journey", "createTransaction"))
	return transaction, nil
}
