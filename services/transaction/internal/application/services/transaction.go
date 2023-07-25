package service

import (
	"github.com/expoure/pismo/transaction/internal/application/port/input"
	"github.com/expoure/pismo/transaction/internal/application/port/output"
)

func NewTransactionDomainService(
	transactionRepository output.TransactionRepository,
	transactionProducer output.TransactionProducer,
	accountClient output.AccountHttpClient,
) input.TransactionDomainService {
	return &transactionDomainService{
		transactionRepository,
		transactionProducer,
		accountClient,
	}
}

type transactionDomainService struct {
	repository    output.TransactionRepository
	producer      output.TransactionProducer
	accountClient output.AccountHttpClient
}
