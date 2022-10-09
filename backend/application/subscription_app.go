package application

import (
	"qubo/qubo-backend/domain/entity"
	"qubo/qubo-backend/infrastructure/service_impl"
	"time"

	"github.com/go-co-op/gocron"
)

type SubscriptionAppInterface interface {
	SaveSubscription(subscription *entity.Subscription) (*entity.Subscription, error)
	GetSubscriptionById(id uint64) (*entity.Subscription, error)
	UpdateSubscription(subscription *entity.Subscription) (*entity.Subscription, error)
	CancelSubscription(id uint64) error
	GetSubscriptionsByMerchantId(id uint64) ([]*entity.Subscription, error)
}

type SubscriptionApp struct {
	subscriptionServiceImpl *service_impl.SubscriptionServiceImpl
}

func NewSubscriptionApplication(subscriptionServiceImpl *service_impl.SubscriptionServiceImpl) *SubscriptionApp {
	return &SubscriptionApp{
		subscriptionServiceImpl: subscriptionServiceImpl,
	}
}

var _ SubscriptionAppInterface = &SubscriptionApp{}

func (subscriptionApp *SubscriptionApp) SaveSubscription(subscription *entity.Subscription) (*entity.Subscription, error) {
	newSubscription, err := subscriptionApp.subscriptionServiceImpl.SaveSubscription(subscription)

	if err != nil {
		return nil, err
	}
	// Create subscription in smart contract

	// add to scheduler to trigger smart contract
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Month().Do(subscriptionApp.subscriptionServiceImpl.TriggerSmartContract, newSubscription)

	return newSubscription, nil
}

func (subscriptionApp *SubscriptionApp) GetSubscriptionById(id uint64) (*entity.Subscription, error) {
	subscription, err := subscriptionApp.subscriptionServiceImpl.GetSubscriptionById(id)

	if err != nil {
		return nil, err
	}

	return subscription, nil
}

func (subscriptionApp *SubscriptionApp) UpdateSubscription(subscription *entity.Subscription) (*entity.Subscription, error) {
	updatedSubscription, err := subscriptionApp.subscriptionServiceImpl.UpdateSubscription(subscription)

	if err != nil {
		return nil, err
	}

	return updatedSubscription, nil
}

func (subscriptionApp *SubscriptionApp) CancelSubscription(id uint64) error {
	err := subscriptionApp.subscriptionServiceImpl.CancelSubscription(id)

	if err != nil {
		return err
	}

	return nil
}

func (subscriptionApp *SubscriptionApp) GetSubscriptionsByMerchantId(id uint64) ([]*entity.Subscription, error) {
	subscriptions, err := subscriptionApp.subscriptionServiceImpl.GetSubscriptionsByMerchantId(id)

	if err != nil {
		return nil, err
	}

	return subscriptions, nil
}
