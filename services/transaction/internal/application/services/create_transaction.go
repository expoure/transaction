package service

import (
	"github.com/expoure/pismo/transaction/internal/application/domain"
	"github.com/expoure/pismo/transaction/internal/configuration/customized_errors"
	"github.com/expoure/pismo/transaction/internal/configuration/logger"
	"go.uber.org/zap"
)

func (td *transactionDomainService) CreateTransactionServices(
	transactionDomain domain.TransactionDomain,
) (*domain.TransactionDomain, *customized_errors.RestErr) {
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
	td.producer.TransactionCreated(*transaction)

	logger.Info(
		"CreateTransaction service executed successfully",
		zap.String("TransactionId", transaction.ID.String()),
		zap.String("journey", "createTransaction"))
	return transaction, nil
}
