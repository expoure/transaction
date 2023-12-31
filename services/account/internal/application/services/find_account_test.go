package service

import (
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/Rhymond/go-money"
	mock_repository "github.com/expoure/pismo/account/internal/adapter/output/repository/mock"
	"github.com/expoure/pismo/account/internal/application/constants"
	"github.com/expoure/pismo/account/internal/application/domain"
	"github.com/expoure/pismo/account/internal/configuration/customized_errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func Test_accountDomainService_FindAccountByIDServices(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()
	repo := mock_repository.NewMockAccountRepositoryPort(control)

	ad := &accountDomainService{
		repository: repo,
		mutex:      &sync.Mutex{},
	}

	t.Run("It finds account by id", func(t *testing.T) {
		accountId := uuid.New()
		want := &domain.AccountDomain{
			ID:             accountId,
			DocumentNumber: "123456789",
			Balance:        money.New(1000, "BRL"),
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			DeletedAt:      nil,
		}

		repo.EXPECT().
			FindAccountByID(gomock.Eq(accountId)).
			Return(want, nil).Times(1)

		got, err := ad.FindAccountByIDServices(accountId)
		require.Nil(t, err)
		require.Equal(t, want, got)
	})

	t.Run("It errors when account not found", func(t *testing.T) {
		repo.EXPECT().FindAccountByID(gomock.Any()).Return(nil, &customized_errors.EntityNotFound)

		_, err := ad.FindAccountByIDServices(uuid.New())
		require.NotNil(t, err)
		require.IsType(t, &customized_errors.RestErr{}, err)
		require.ErrorContains(t, err, constants.ErrAccountNotFound)
	})

	t.Run("It errors unknown when trying to find account by id", func(t *testing.T) {
		unknowError := errors.New("unknow error")
		repo.EXPECT().FindAccountByID(gomock.Any()).Return(nil, &unknowError)

		_, err := ad.FindAccountByIDServices(uuid.New())
		require.NotNil(t, err)
		require.IsType(t, &customized_errors.RestErr{}, err)
	})
}

func Test_accountDomainService_FindAccountByDocumentNumberServices(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()
	repo := mock_repository.NewMockAccountRepositoryPort(control)

	ad := &accountDomainService{
		repository: repo,
		mutex:      &sync.Mutex{},
	}

	t.Run("It finds account by DocumentNumber", func(t *testing.T) {
		documentNumber := "09876543210"
		want := &domain.AccountDomain{
			ID:             uuid.New(),
			DocumentNumber: documentNumber,
			Balance:        money.New(1000, "BRL"),
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			DeletedAt:      nil,
		}

		repo.EXPECT().
			FindAccountByDocumentNumber(gomock.Eq(documentNumber)).
			Return(want, nil).Times(1)

		got, err := ad.FindAccountByDocumentNumberServices(documentNumber)
		require.Nil(t, err)
		require.Equal(t, want, got)
	})

	t.Run("It errors when account not found by document number", func(t *testing.T) {
		repo.EXPECT().FindAccountByDocumentNumber(gomock.Any()).Return(nil, &customized_errors.EntityNotFound)

		_, err := ad.FindAccountByDocumentNumberServices("09876543210")
		require.NotNil(t, err)
		require.IsType(t, &customized_errors.RestErr{}, err)
		require.Equal(t, constants.ErrAccountNotFound, err.Error())
	})

	t.Run("It errors unknown when trying to find account by document number", func(t *testing.T) {
		unknownError := errors.New("unknow error")
		repo.EXPECT().FindAccountByDocumentNumber(gomock.Any()).Return(nil, &unknownError)

		_, err := ad.FindAccountByDocumentNumberServices("09876543210")
		require.NotNil(t, err)
		require.IsType(t, &customized_errors.RestErr{}, err)
		require.Equal(t, constants.ErrInternalServerError, err.Error())
	})
}
