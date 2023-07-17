package service

import (
	"github.com/expoure/pismo/account/internal/application/port/input"
	"github.com/expoure/pismo/account/internal/application/port/output"
)

func NewAccountDomainService(
	accountRepository output.AccountPort,
) input.AccountDomainService {
	return &accountDomainService{
		accountRepository,
	}
}

type accountDomainService struct {
	repository output.AccountPort
}
