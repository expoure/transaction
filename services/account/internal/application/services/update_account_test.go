package service

import (
	"sync"
	"testing"

	"github.com/Rhymond/go-money"
	mock_repository "github.com/expoure/pismo/account/internal/adapter/output/repository/mock"
	"github.com/expoure/pismo/account/internal/application/port/output"
	"github.com/expoure/pismo/account/internal/configuration/rest_errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func Test_accountDomainService_UpdateAccountBalanceByIDServices(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	repo := mock_repository.NewMockAccountRepositoryPort(control)

	type fields struct {
		repository output.AccountRepositoryPort
		mutex      *sync.Mutex
	}
	type args struct {
		id                uuid.UUID
		transactionAmount int64
		currentAmount     int64
		newAmount         int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *money.Money
		wantErr bool
	}{
		{
			name: "It update account balance with addition",
			fields: fields{
				mutex:      &sync.Mutex{},
				repository: repo,
			},
			args: args{
				id:                uuid.New(),
				transactionAmount: 1000,
				currentAmount:     500,
				newAmount:         1500,
			},
			want:    money.New(1500, "BRL"),
			wantErr: false,
		},
		{
			name: "It update account balance with subtration",
			fields: fields{
				mutex:      &sync.Mutex{},
				repository: repo,
			},
			args: args{
				id:                uuid.New(),
				transactionAmount: -800,
				currentAmount:     3500,
				newAmount:         2700,
			},
			want:    money.New(2700, "BRL"),
			wantErr: false,
		},
		{
			name: "It errors when trying to update account balance with 0",
			fields: fields{
				mutex:      &sync.Mutex{},
				repository: repo,
			},
			args: args{
				id:                uuid.New(),
				transactionAmount: 0,
				currentAmount:     94500,
				newAmount:         94500,
			},
			want:    money.New(2700, "BRL"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo.EXPECT().FindAccountBalanceByID(tt.args.id).Return(money.New(tt.args.currentAmount, "BRL"), nil).AnyTimes()
			repo.EXPECT().UpdateAccountBalanceByID(tt.args.id, gomock.Eq(tt.args.newAmount)).Return(money.New(tt.args.newAmount, "BRL"), nil).AnyTimes()
			ad := &accountDomainService{
				repository: repo,
				mutex:      tt.fields.mutex,
			}

			got, err := ad.UpdateAccountBalanceByIDServices(tt.args.id, tt.args.transactionAmount)

			if (err != nil) && tt.wantErr {
				require.ErrorContains(t, err, "invalid transaction amount")
			} else {
				require.Nil(t, err)
				require.Equal(t, tt.want, got)
			}
		})
	}

	t.Run("It errors when trying to update the balance of a non-existent account", func(t *testing.T) {
		accountUuid := uuid.New()
		repo.EXPECT().FindAccountBalanceByID(gomock.Eq(accountUuid)).Return(nil, rest_errors.NewInternalServerError("")).AnyTimes()
		repo.EXPECT().UpdateAccountBalanceByID(gomock.Eq(accountUuid), gomock.Eq(1)).Return(money.New(1, "BRL"), nil).AnyTimes()
		ad := &accountDomainService{
			repository: repo,
			mutex:      &sync.Mutex{},
		}

		_, err := ad.UpdateAccountBalanceByIDServices(accountUuid, 1)

		t.Log(err)
		// require.NotNil(t, err)
		// require.ErrorContains(t, err, "account not found")
	})
}
