package service

import (
	"github.com/Rhymond/go-money"
	"github.com/expoure/pismo/account/internal/configuration/logger"
	"github.com/expoure/pismo/account/internal/configuration/rest_errors"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (ad *accountDomainService) UpdateAccountBalanceByIDServices(
	id uuid.UUID,
	transactionAmount int64,
) (*money.Money, *rest_errors.RestErr) {
	logger.Info("Init updateAccountBalance.",
		zap.String("journey", "updateAccountBalance"))

	ad.mutex.Lock()
	defer ad.mutex.Unlock()

	result, err := ad.repository.FindAccountBalanceByID(id)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "updateAccountBalance"))
		return nil, err
	}

	newBalance, _ := result.Add(money.New(transactionAmount, "BRL"))

	balance, err := ad.repository.UpdateAccountBalanceByID(id, newBalance.Amount())

	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "updateAccountBalance"))
		return nil, err
	}
	// criar evento de balance_updated

	logger.Info(
		"UpdateAccountBalanceByIDServices service executed successfully",
		zap.String("journey", "updateAccountBalance"))

	return balance, nil
}
