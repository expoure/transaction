package service

import (
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/Rhymond/go-money"
	mock_repository "github.com/expoure/pismo/account/internal/adapter/output/repository/mock"
	"github.com/expoure/pismo/account/internal/application/domain"
	"github.com/expoure/pismo/account/internal/application/port/output"
	"github.com/expoure/pismo/account/internal/configuration/rest_errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func Test_accountDomainService_FindAccountByIDServices(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()
	repo := mock_repository.NewMockAccountRepositoryPort(control)

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

		ad := &accountDomainService{
			repository: repo,
			mutex:      &sync.Mutex{},
		}

		got, err := ad.FindAccountByIDServices(accountId)
		require.Nil(t, err)
		require.Equal(t, want, got)
	})

	t.Run("It errors when account not found", func(t *testing.T) {
		want := &domain.AccountDomain{}
		repo.EXPECT().FindAccountByID(gomock.Any()).Return(want, rest_errors.NewBadRequestError("account not found"))

		ad := &accountDomainService{
			repository: repo,
			mutex:      &sync.Mutex{},
		}

		_, err := ad.FindAccountByIDServices(uuid.New())
		require.NotNil(t, err)
		// criar validation de repo
	})
}

func Test_accountDomainService_FindAccountByDocumentNumberServices(t *testing.T) {
	type fields struct {
		repository output.AccountRepositoryPort
		mutex      *sync.Mutex
	}
	type args struct {
		documentNumber string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *domain.AccountDomain
		want1  *rest_errors.RestErr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ad := &accountDomainService{
				repository: tt.fields.repository,
				mutex:      tt.fields.mutex,
			}
			got, got1 := ad.FindAccountByDocumentNumberServices(tt.args.documentNumber)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("accountDomainService.FindAccountByDocumentNumberServices() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("accountDomainService.FindAccountByDocumentNumberServices() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
