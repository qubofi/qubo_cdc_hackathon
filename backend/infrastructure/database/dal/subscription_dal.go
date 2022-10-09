package dal

import (
	"qubo/qubo-backend/domain/entity"
	"qubo/qubo-backend/infrastructure/database"
	"time"
)

func InsertIntoSubscriptions(subscription *entity.Subscription) (*entity.Subscription, error) {
	started_date_time, _ := time.Parse("2006-01-02T15:04:05.000Z", subscription.StartedDate)
	last_paid_date_time, _ := time.Parse("2006-01-02T15:04:05.000Z", subscription.LastPaidDate)
	started_date_time_str := started_date_time.Format("2006-01-02 15:04:05")
	last_paid_date_time_str := last_paid_date_time.Format("2006-01-02 15:04:05")

	println(started_date_time_str, last_paid_date_time_str)
	err := database.Db.QueryRow(`INSERT INTO "subscriptions"(product_id, customer_wallet_address, started_date, last_paid_date, cancelled_at) VALUES($1, $2, $3, $4, $5) RETURNING id`,
		subscription.ProductId,
		subscription.CustomerWalletAddress,
		started_date_time_str,
		last_paid_date_time_str,
		"").Scan(&subscription.ID)

	if err != nil {
		return nil, err
	}

	return subscription, nil
}

func SelectFromSubscriptionsById(id uint64) (*entity.Subscription, error) {
	subscription := &entity.Subscription{}
	err := database.Db.QueryRow(`SELECT id, product_id, customer_wallet_address, started_date, last_paid_date, cancelled_at FROM "subscriptions" WHERE id = $1`, id).
		Scan(&subscription.ID, &subscription.ProductId, &subscription.CustomerWalletAddress, &subscription.StartedDate, &subscription.LastPaidDate, &subscription.CancelledAt)

	if err != nil {
		return nil, err
	}

	return subscription, nil
}

func UpdateSubscription(subscription *entity.Subscription) (*entity.Subscription, error) {
	started_date_time, _ := time.Parse("2006-01-02T15:04:05.000Z", subscription.StartedDate)
	last_paid_date_time, _ := time.Parse("2006-01-02T15:04:05.000Z", subscription.LastPaidDate)
	started_date_time_str := started_date_time.Format("2006-01-02 15:04:05")
	last_paid_date_time_str := last_paid_date_time.Format("2006-01-02 15:04:05")
	_, err := database.Db.Exec(`UPDATE "subscriptions" SET product_id = $1, customer_wallet_address = $2, started_date = $3, last_paid_date = $4, cancelled_at = $5 WHERE id = $6`,
		subscription.ProductId,
		subscription.CustomerWalletAddress,
		started_date_time_str,
		last_paid_date_time_str,
		subscription.CancelledAt,
		subscription.ID)

	if err != nil {
		return nil, err
	}

	return subscription, nil
}

func CancelSubscription(id uint64) error {
	_, err := database.Db.Exec(`UPDATE "subscriptions" SET cancelled_at = NOW() WHERE id = $1`, id)

	if err != nil {
		return err
	}

	return nil
}

func SelectFromSubscriptionsByMerchantId(merchantId uint64) ([]*entity.Subscription, error) {
	subscriptions := []*entity.Subscription{}
	rows, err := database.Db.Query(`SELECT id, product_id, customer_wallet_address, started_date, last_paid_date, cancelled_at FROM "subscriptions" WHERE product_id IN (SELECT id FROM "products" WHERE merchant_id = $1)`, merchantId)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		subscription := &entity.Subscription{}
		err := rows.Scan(&subscription.ID, &subscription.ProductId, &subscription.CustomerWalletAddress, &subscription.StartedDate, &subscription.LastPaidDate, &subscription.CancelledAt)

		if err != nil {
			return nil, err
		}

		subscriptions = append(subscriptions, subscription)
	}

	return subscriptions, nil
}
