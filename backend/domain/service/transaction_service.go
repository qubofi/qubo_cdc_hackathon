package service

import (
	"qubo/qubo-backend/domain/entity"
)

type TransactionService interface {
	SaveTransaction(transaction *entity.Transaction) (*entity.Transaction, error)
	GetTransactionById(id uint64) (*entity.Transaction, error)
	GetTransactionsBySubscriptionId(subscriptionId uint64) ([]entity.Transaction, error)
}
