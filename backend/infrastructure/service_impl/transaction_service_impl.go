package service_impl

import (
	"database/sql"
	"qubo/qubo-backend/domain/entity"
	"qubo/qubo-backend/domain/service"
	"qubo/qubo-backend/infrastructure/database"
	"qubo/qubo-backend/infrastructure/database/dal"
)

type TransactionServiceImpl struct {
	db *sql.DB
}

func NewTransactionService() *TransactionServiceImpl {
	return &TransactionServiceImpl{
		db: database.Db,
	}
}

var _ service.TransactionService = &TransactionServiceImpl{}

func (transactionRepo *TransactionServiceImpl) SaveTransaction(transaction *entity.Transaction) (*entity.Transaction, error) {
	newTransaction, err := dal.InsertIntoTransactions(transaction)

	if err != nil {
		return nil, err
	}

	return newTransaction, nil
}

func (transactionRepo *TransactionServiceImpl) GetTransactionById(id uint64) (*entity.Transaction, error) {
	transaction, err := dal.SelectFromTransactionsById(id)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (transactionRepo *TransactionServiceImpl) GetTransactionsBySubscriptionId(subscriptionId uint64) ([]entity.Transaction, error) {
	transactions, err := dal.SelectFromTransactionsBySubscriptionId(subscriptionId)

	if err != nil {
		return nil, err
	}

	return transactions, nil
}
