package service

import (
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/Rhymond/go-money"
	mock_repository "github.com/expoure/pismo/account/internal/adapter/output/repository/mock"
	"github.com/expoure/pismo/account/internal/application/domain"
	"github.com/expoure/pismo/account/internal/configuration/customized_errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func Test_accountDomainService_CreateAccountServices(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()
	repo := mock_repository.NewMockAccountRepositoryPort(control)
	ad := &accountDomainService{
		repository: repo,
		mutex:      &sync.Mutex{},
	}

	accountToCreate := domain.AccountDomain{DocumentNumber: "15195811067"}
	accountCreated := domain.AccountDomain{
		ID:             uuid.New(),
		DocumentNumber: "15195811067",
		Balance:        money.New(0000, "BRL"),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	t.Run("It sucessfully create account", func(t *testing.T) {
		repo.EXPECT().
			CreateAccount(gomock.Eq(accountToCreate)).
			Return(&accountCreated, nil).Times(1)

		account, err := ad.CreateAccountServices(accountToCreate)
		require.Nil(t, err)
		require.Equal(t, accountCreated, *account)
	})

	t.Run("It errors unknow in repository when create account", func(t *testing.T) {
		repoError := errors.New("unknow error")
		repo.EXPECT().
			CreateAccount(gomock.Eq(accountToCreate)).
			Return(nil, &repoError).Times(1)

		_, err := ad.CreateAccountServices(accountToCreate)
		require.NotNil(t, err)
		require.ErrorContains(t, err, "Was not possible to create account")
	})

	t.Run("It errors when account already exists", func(t *testing.T) {
		errTo := &customized_errors.DuplicateKey
		repo.EXPECT().
			CreateAccount(gomock.Eq(accountToCreate)).
			Return(nil, errTo).Times(1)

		_, err := ad.CreateAccountServices(accountToCreate)

		require.NotNil(t, err)
		require.IsType(t, &customized_errors.RestErr{}, err)
		require.ErrorContains(t, err, "This document number is already registered")
	})

	t.Run("It errors if documentNumber is invalid", func(t *testing.T) {
		_, err := ad.CreateAccountServices(domain.AccountDomain{DocumentNumber: "123456789055"})

		require.NotNil(t, err)
		require.IsType(t, &customized_errors.RestErr{}, err)
		require.Equal(t, err.Message, "Invalid document number")
	})

	t.Run("It errors if documentNumber is empty", func(t *testing.T) {
		_, err := ad.CreateAccountServices(domain.AccountDomain{DocumentNumber: ""})

		require.NotNil(t, err)
		require.IsType(t, &customized_errors.RestErr{}, err)
		require.Equal(t, err.Message, "Document number is required")
	})
}
