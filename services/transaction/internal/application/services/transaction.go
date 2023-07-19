package service

import (
	"github.com/expoure/pismo/transaction/internal/application/port/input"
	"github.com/expoure/pismo/transaction/internal/application/port/output"
)

func NewTransactionDomainService(
	transactionRepository output.TransactionRepository,
	transactionProducer output.TransactionProducer,
) input.TransactionDomainService {
	return &transactionDomainService{
		transactionRepository,
		transactionProducer,
	}
}

type transactionDomainService struct {
	repository output.TransactionRepository
	producer   output.TransactionProducer
}
