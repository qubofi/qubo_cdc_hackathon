package application

import (
	"qubo/qubo-backend/domain/entity"
	"qubo/qubo-backend/infrastructure/service_impl"
)

type TransactionAppInterface interface {
	SaveTransaction(transaction *entity.Transaction) (*entity.Transaction, error)
	GetTransactionById(id uint64) (*entity.Transaction, error)
	GetTransactionsBySubscriptionId(subscriptionId uint64) ([]entity.Transaction, error)
}

type transactionApp struct {
	transactionServiceImpl *service_impl.TransactionServiceImpl
}

var _ TransactionAppInterface = &transactionApp{}

func NewTransactionApplication(transactionServiceImpl *service_impl.TransactionServiceImpl) *transactionApp {
	return &transactionApp{
		transactionServiceImpl: transactionServiceImpl,
	}
}

func (transactionApp *transactionApp) SaveTransaction(transaction *entity.Transaction) (*entity.Transaction, error) {
	return transactionApp.transactionServiceImpl.SaveTransaction(transaction)
}

func (transactionApp *transactionApp) GetTransactionById(id uint64) (*entity.Transaction, error) {
	return transactionApp.transactionServiceImpl.GetTransactionById(id)
}

func (transactionApp *transactionApp) GetTransactionsBySubscriptionId(subscriptionId uint64) ([]entity.Transaction, error) {
	return transactionApp.transactionServiceImpl.GetTransactionsBySubscriptionId(subscriptionId)
}
