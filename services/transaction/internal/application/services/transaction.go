package service

import (
	"github.com/expoure/pismo/transaction/internal/application/port/input"
	"github.com/expoure/pismo/transaction/internal/application/port/output"
)

func NewTransactionDomainService(
	transactionRepository output.TransactionPort,
) input.TransactionDomainService {
	return &transactionDomainService{
		transactionRepository,
	}
}

type transactionDomainService struct {
	repository output.TransactionPort
}
