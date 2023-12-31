package service

import (
	"github.com/expoure/pismo/transaction/internal/application/constants"
	"github.com/expoure/pismo/transaction/internal/application/domain"
	"github.com/expoure/pismo/transaction/internal/configuration/customized_errors"
	"github.com/expoure/pismo/transaction/internal/configuration/logger"
	"go.uber.org/zap"
)

func (td *transactionDomainService) CreateTransactionServices(
	transactionDomain domain.TransactionDomain,
) (*domain.TransactionDomain, *customized_errors.RestErr) {
	logger.Info("Init createTransaction service.",
		zap.String("journey", "createTransaction"))

	if err := td.validateTransaction(&transactionDomain); err != nil {
		return nil, err
	}

	transaction, err := td.repository.CreateTransaction(transactionDomain)

	if err != nil {
		logger.Error("Error trying to call repository",
			*err,
			zap.String("journey", "createTransaction"))
		return nil, customized_errors.NewInternalServerError(constants.ErrWasNotPossibleToCreateTransaction)

	}
	// criar evento de transaction_created
	td.producer.TransactionCreated(*transaction)

	logger.Info(
		"CreateTransaction service executed successfully",
		zap.String("TransactionId", transaction.ID.String()),
		zap.String("journey", "createTransaction"))

	return transaction, nil
}

func (td *transactionDomainService) checkIfAcountHasEnoughMoney(transactionDomain *domain.TransactionDomain) *customized_errors.RestErr {
	account, err := td.accountClient.GetAccount(transactionDomain.AccountID.String())
	if err != nil {
		return customized_errors.NewInternalServerError(constants.ErrWasNotPossibleToCreateTransaction)
	}
	if account.Balance.SmallUnitAmount < -transactionDomain.Amount.Amount() {
		return customized_errors.NewBadRequestError(constants.ErrNotEnoughMoney)
	}

	return nil
}

func (td *transactionDomainService) validateTransaction(transaction *domain.TransactionDomain) *customized_errors.RestErr {
	switch transaction.OperationTypeID {

	case constants.OperationPayCash, constants.OperationPayInstallments, constants.OperationWithdraw:
		if transaction.Amount.Amount() >= 0 {
			return customized_errors.NewBadRequestError(constants.InvalidOutboundOperation)
		}
		return td.checkIfAcountHasEnoughMoney(transaction)

	case constants.OperationPayment:
		if transaction.Amount.Amount() <= 0 {
			return customized_errors.NewBadRequestError(constants.InvalidInboundOperation)
		}

	default:
		return customized_errors.NewBadRequestError(constants.InvalidOperation)
	}

	return nil
}
