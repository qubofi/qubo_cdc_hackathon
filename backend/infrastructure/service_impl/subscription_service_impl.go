package service_impl

import (
	"database/sql"
	"qubo/qubo-backend/domain/entity"
	"qubo/qubo-backend/domain/service"
	"qubo/qubo-backend/infrastructure/database"
	"qubo/qubo-backend/infrastructure/database/dal"
)

type SubscriptionServiceImpl struct {
	db *sql.DB
}

func NewSubscriptionService() *SubscriptionServiceImpl {
	return &SubscriptionServiceImpl{
		db: database.Db,
	}
}

var _ service.SubscriptionService = &SubscriptionServiceImpl{}

func (subscriptionRepo *SubscriptionServiceImpl) SaveSubscription(subscription *entity.Subscription) (*entity.Subscription, error) {
	newSubscription, err := dal.InsertIntoSubscriptions(subscription)

	if err != nil {
		return nil, err
	}

	return newSubscription, nil
}

func (subscriptionRepo *SubscriptionServiceImpl) GetSubscriptionById(id uint64) (*entity.Subscription, error) {
	subscription, err := dal.SelectFromSubscriptionsById(id)

	if err != nil {
		return nil, err
	}

	return subscription, nil
}

func (subscriptionRepo *SubscriptionServiceImpl) UpdateSubscription(subscription *entity.Subscription) (*entity.Subscription, error) {
	updatedSubscription, err := dal.UpdateSubscription(subscription)

	if err != nil {
		return nil, err
	}

	return updatedSubscription, nil
}

func (subscriptionRepo *SubscriptionServiceImpl) CancelSubscription(id uint64) error {
	err := dal.CancelSubscription(id)

	if err != nil {
		return err
	}

	return nil
}

func (subscriptionRepo *SubscriptionServiceImpl) GetSubscriptionsByMerchantId(id uint64) ([]*entity.Subscription, error) {
	subscriptions, err := dal.SelectFromSubscriptionsByMerchantId(id)

	if err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func (subscriptionRepo *SubscriptionServiceImpl) TriggerSmartContract(subscription *entity.Subscription) error {
	// Trigger smart contract method executePayment
	return nil
}
