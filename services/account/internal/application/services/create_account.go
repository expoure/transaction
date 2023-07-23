package service

import (
	"regexp"

	"github.com/expoure/pismo/account/internal/application/constants"
	"github.com/expoure/pismo/account/internal/application/domain"
	"github.com/expoure/pismo/account/internal/configuration/customized_errors"
	"github.com/expoure/pismo/account/internal/configuration/logger"
	"go.uber.org/zap"
)

func (ad *accountDomainService) CreateAccountServices(
	accountDomain domain.AccountDomain,
) (*domain.AccountDomain, *customized_errors.RestErr) {
	logger.Info("Init createAccount model.",
		zap.String("journey", "createAccount"))

	if err := validateDocumentNumber(accountDomain.DocumentNumber); err != nil {
		return nil, err
	}

	accountCreated, err := ad.repository.CreateAccount(accountDomain)
	if err != nil {
		logger.Error("Error trying to call repository",
			*err,
			zap.String("journey", "createAccount"))

		if err == &customized_errors.DuplicateKey {
			return nil, customized_errors.NewBadRequestError(constants.ErrDocumentNumberAlreadyRegistered)
		}

		return nil, customized_errors.NewInternalServerError(constants.ErrWasNotPossibleToCreateAccount)
	}
	// criar evento de account_created

	logger.Info(
		"CreateAccount service executed successfully",
		zap.String("AccountId", accountCreated.ID.String()),
		zap.String("journey", "createAccount"))

	return accountCreated, nil
}

func validateDocumentNumber(documentNumber string) *customized_errors.RestErr {
	if documentNumber == "" {
		return customized_errors.NewBadRequestError(constants.ErrDocumentNumberIsRequire)
	}
	re := regexp.MustCompile(`^([0-9]{3}[\.]?[0-9]{3}[\.]?[0-9]{3}[-]?[0-9]{2})$`)
	if !re.MatchString(documentNumber) {
		return customized_errors.NewBadRequestError(constants.ErrInvalidDocumentNumber)
	}

	return nil
}
