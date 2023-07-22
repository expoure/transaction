package service

import (
	"github.com/expoure/pismo/transaction/internal/application/constants"
	"github.com/expoure/pismo/transaction/internal/application/domain"
	"github.com/expoure/pismo/transaction/internal/configuration/customized_errors"
	"github.com/expoure/pismo/transaction/internal/configuration/logger"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (td *transactionDomainService) ListTransactionsService(
	accoountId uuid.UUID,
) (*[]domain.TransactionDomain, *customized_errors.RestErr) {
	logger.Info("Init listTransaction",
		zap.String("journey", "listTransaction"))

	transactions, err := td.repository.ListTransactionsByAccountID(accoountId)

	if err != nil {
		logger.Error("Error trying to call repository",
			*err,
			zap.String("journey", "listTransaction"))
		return nil, customized_errors.NewInternalServerError(constants.ErrWasNotPossibleToListTransactions)
	}

	// criar evento de transaction_created

	logger.Info(
		"ListTransactions service executed successfully",
		zap.String("journey", "listTransaction"))

	return transactions, nil
}
