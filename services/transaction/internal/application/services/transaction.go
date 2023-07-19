package service

import (
	"github.com/expoure/pismo/transaction/internal/application/port/input"
	"github.com/expoure/pismo/transaction/internal/application/port/output"
)

func NewTransactionDomainService(
	transactionRepository output.TransactionPort,
) input.AccountDomainService {
	return &transactionDomainService{
		transactionRepository,
	}
}

type transactionDomainService struct {
	repository output.TransactionPort
}
