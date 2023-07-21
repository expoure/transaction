package service

import (
	"github.com/Rhymond/go-money"
	"github.com/expoure/pismo/account/internal/configuration/customized_errors"
	"github.com/expoure/pismo/account/internal/configuration/logger"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (ad *accountDomainService) UpdateAccountBalanceByIDServices(
	id uuid.UUID,
	transactionAmount int64,
) (*money.Money, *customized_errors.RestErr) {
	if transactionAmount == 0 {
		return nil, customized_errors.NewBadRequestError("Is not possible to update balance with 0")
	}

	logger.Info("Init updateAccountBalance.",
		zap.String("journey", "updateAccountBalance"))

	ad.mutex.Lock()
	defer ad.mutex.Unlock()

	result, err := ad.repository.FindAccountBalanceByID(id)
	if err != nil {
		_, errHandled := handleFindError(err)
		return nil, errHandled
	}

	newBalance, _ := result.Add(money.New(transactionAmount, "BRL"))

	balance, err := ad.repository.UpdateAccountBalanceByID(id, newBalance.Amount())

	if err != nil {
		logger.Error("Error trying to call repository",
			*err,
			zap.String("journey", "updateAccountBalance"))
		return nil, customized_errors.NewInternalServerError("")
	}
	// criar evento de balance_updated

	logger.Info(
		"UpdateAccountBalanceByIDServices service executed successfully",
		zap.String("journey", "updateAccountBalance"))

	return balance, nil
}
