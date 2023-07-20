package service

import (
	"sync"
	"testing"

	"github.com/Rhymond/go-money"
	mock_repository "github.com/expoure/pismo/account/internal/adapter/output/repository/mock"
	"github.com/expoure/pismo/account/internal/application/port/output"
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
		name   string
		fields fields
		args   args
		want   int64
	}{
		{
			name: "It should update account balance by id",
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
			want: 1500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo.EXPECT().FindAccountBalanceByID(tt.args.id).Return(money.New(tt.args.currentAmount, "BRL"), nil).Times(1)
			repo.EXPECT().UpdateAccountBalanceByID(tt.args.id, gomock.Eq(tt.args.newAmount)).Return(money.New(1500, "BRL"), nil).Times(1)
			ad := &accountDomainService{
				repository: repo,
				mutex:      tt.fields.mutex,
			}

			got, err := ad.UpdateAccountBalanceByIDServices(tt.args.id, tt.args.transactionAmount)
			require.Nil(t, err)
			require.Equal(t, tt.want, got.Amount())
		})
	}
}
