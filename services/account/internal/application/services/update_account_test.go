package service

import (
	"errors"
	"sync"
	"testing"

	"github.com/Rhymond/go-money"
	mock_repository "github.com/expoure/pismo/account/internal/adapter/output/repository/mock"
	"github.com/expoure/pismo/account/internal/configuration/customized_errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func Test_accountDomainService_UpdateAccountBalanceByIDServices(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	repo := mock_repository.NewMockAccountRepositoryPort(control)

	ad := &accountDomainService{
		repository: repo,
		mutex:      &sync.Mutex{},
	}

	type args struct {
		id                uuid.UUID
		transactionAmount int64
		currentAmount     int64
		newAmount         int64
	}
	tests := []struct {
		name    string
		args    args
		want    *money.Money
		wantErr bool
	}{
		{
			name: "It update account balance with addition",
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

			got, err := ad.UpdateAccountBalanceByIDServices(tt.args.id, tt.args.transactionAmount)

			if (err != nil) && tt.wantErr {
				require.ErrorContains(t, err, "Is not possible to update balance with 0")
			} else {
				require.Nil(t, err)
				require.Equal(t, tt.want, got)
			}
		})
	}

	t.Run("It errors when trying to update the balance of a non-existent account", func(t *testing.T) {
		accountUuid := uuid.New()
		repo.EXPECT().FindAccountBalanceByID(gomock.Eq(accountUuid)).Return(nil, &customized_errors.EntityNotFound).AnyTimes()
		repo.EXPECT().UpdateAccountBalanceByID(gomock.Eq(accountUuid), gomock.Eq(1)).Return(money.New(1, "BRL"), nil).AnyTimes()

		_, err := ad.UpdateAccountBalanceByIDServices(accountUuid, 1)

		require.NotNil(t, err)
		require.IsType(t, &customized_errors.RestErr{}, err)
		require.Equal(t, "Account not found", err.Error())
	})

	t.Run("It errors unknown when trying to update the balance of an account - findAccountBalanceByID", func(t *testing.T) {
		unknownError := errors.New("unknow error")
		repo.EXPECT().FindAccountBalanceByID(gomock.Any()).Return(nil, &unknownError)

		_, err := ad.UpdateAccountBalanceByIDServices(uuid.New(), 1)
		require.NotNil(t, err)
	})

	t.Run("It errors unknown when trying to update the balance of an account - updateAccountBalanceByID", func(t *testing.T) {
		unknownError := errors.New("unknow error")
		repo.EXPECT().FindAccountBalanceByID(gomock.Any()).Return(money.New(1, "BRL"), nil)
		repo.EXPECT().UpdateAccountBalanceByID(gomock.Any(), gomock.Any()).Return(nil, &unknownError)

		_, err := ad.UpdateAccountBalanceByIDServices(uuid.New(), 1)
		require.NotNil(t, err)
	})
}
