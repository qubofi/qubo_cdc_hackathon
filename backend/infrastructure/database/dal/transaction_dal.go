package dal

import (
	"qubo/qubo-backend/domain/entity"
	"qubo/qubo-backend/infrastructure/database"
)

func InsertIntoTransactions(transaction *entity.Transaction) (*entity.Transaction, error) {
	err := database.Db.QueryRow(`INSERT INTO "transactions"(subscription_id, transaction_hash) VALUES($1, $2) RETURNING id`,
		transaction.SubscriptionId,
		transaction.TransactionHash).Scan(&transaction.ID)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func SelectFromTransactionsById(id uint64) (*entity.Transaction, error) {
	transaction := &entity.Transaction{}
	err := database.Db.QueryRow(`SELECT id, subscription_id, transaction_hash FROM "transactions" WHERE id = $1`, id).
		Scan(&transaction.ID, &transaction.SubscriptionId, &transaction.TransactionHash)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func SelectFromTransactionsBySubscriptionId(subscriptionId uint64) ([]entity.Transaction, error) {
	rows, err := database.Db.Query(`SELECT id, subscription_id, transaction_hash FROM "transactions" WHERE subscription_id = $1`, subscriptionId)
	if err != nil {
		return nil, err
	}

	transactions := make([]entity.Transaction, 0)
	for rows.Next() {
		transaction := &entity.Transaction{}
		err := rows.Scan(&transaction.ID, &transaction.SubscriptionId, &transaction.TransactionHash)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, *transaction)
	}

	return transactions, nil
}
