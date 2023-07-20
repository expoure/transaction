package sqlc

import (
	"context"
	"reflect"
	"testing"

	custom_types "github.com/expoure/pismo/account/internal/configuration/database/custom_types"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	type fields struct {
		db DBTX
	}
	type args struct {
		ctx            context.Context
		documentNumber string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Account
		wantErr bool
	}{
		{
			name:    "It create an account",
			fields:  fields{db: TestQueries.db},
			args:    args{ctx: context.Background(), documentNumber: "12345678956"},
			want:    Account{DocumentNumber: "12345678956"},
			wantErr: false,
		},
		{
			name:    "It does not create an account",
			fields:  fields{db: TestQueries.db},
			args:    args{ctx: context.Background(), documentNumber: "12345678956"},
			want:    Account{DocumentNumber: "12345678956"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queries{
				db: tt.fields.db,
			}
			got, err := q.CreateAccount(tt.args.ctx, tt.args.documentNumber)
			if (err != nil) && tt.wantErr {
				require.ErrorContainsf(t, err, "duplicate key", "")
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, got)
				require.Equal(t, tt.want.DocumentNumber, got.DocumentNumber)
				require.NotNil(t, got.CreatedAt)
				require.NotNil(t, got.UpdatedAt)
			}
		})
	}
}

func TestFindAccountBalanceById(t *testing.T) {
	accountUUID, err := findFirstAccountIdHelper()
	require.NoError(t, err)

	type fields struct {
		db DBTX
	}
	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *custom_types.Money
		wantErr bool
	}{
		{
			name:    "It find an account balance",
			fields:  fields{db: TestQueries.db},
			args:    args{ctx: context.TODO(), id: accountUUID},
			want:    &custom_types.Money{Amount: 0, Currency: "BRL"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queries{
				db: tt.fields.db,
			}
			got, err := q.FindAccountBalanceById(tt.args.ctx, tt.args.id)
			if (err != nil) && tt.wantErr {
				require.Error(t, err)
			} else {
				require.Equal(t, tt.want, got)
			}
		})
	}
}

func TestFindAccountByDocumentNumber(t *testing.T) {
	type fields struct {
		db DBTX
	}
	type args struct {
		ctx            context.Context
		documentNumber string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Account
		wantErr bool
	}{
		{
			name:    "It find an account by document number",
			fields:  fields{db: TestQueries.db},
			args:    args{ctx: context.TODO(), documentNumber: "12345678956"},
			want:    Account{DocumentNumber: "12345678956"},
			wantErr: false,
		},
		{
			name:    "It does not find an account by document number",
			fields:  fields{db: TestQueries.db},
			args:    args{ctx: context.TODO(), documentNumber: "12345678955"},
			want:    Account{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queries{
				db: tt.fields.db,
			}
			got, err := q.FindAccountByDocumentNumber(tt.args.ctx, tt.args.documentNumber)
			if (err != nil) && tt.wantErr {
				require.ErrorContains(t, err, "no rows in result set")
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want.DocumentNumber, got.DocumentNumber)
				require.NotNil(t, got.CreatedAt)
				require.NotNil(t, got.UpdatedAt)
				require.IsType(t, tt.want.ID, uuid.UUID{})
			}
		})
	}
}

func TestFindAccountById(t *testing.T) {
	type fields struct {
		db DBTX
	}
	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Account
		wantErr bool
	}{
		{
			name:    "It does not find an account by id",
			fields:  fields{db: TestQueries.db},
			args:    args{ctx: context.TODO(), id: uuid.New()},
			want:    Account{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queries{
				db: tt.fields.db,
			}
			got, err := q.FindAccountById(tt.args.ctx, tt.args.id)
			if (err != nil) && tt.wantErr {
				require.ErrorContains(t, err, "no rows in result set")
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queries.FindAccountById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func findFirstAccountIdHelper() (uuid.UUID, error) {
	row := TestQueries.db.QueryRowContext(context.TODO(), "SELECT id FROM account LIMIT 1")
	var accountUUID uuid.UUID
	err := row.Scan(&accountUUID)
	return accountUUID, err
}
