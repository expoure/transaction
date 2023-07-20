package service

import (
	"sync"

	"github.com/expoure/pismo/account/internal/application/port/input"
	"github.com/expoure/pismo/account/internal/application/port/output"
)

func NewAccountDomainService(
	accountRepository output.AccountRepositoryPort,
) input.AccountDomainService {
	return &accountDomainService{
		accountRepository,
		&sync.Mutex{},
	}
}

type accountDomainService struct {
	repository output.AccountRepositoryPort
	mutex      *sync.Mutex
}
