package service

import (
	"qubo/qubo-backend/domain/entity"
)

type SubscriptionService interface {
	SaveSubscription(subscription *entity.Subscription) (*entity.Subscription, error)
	GetSubscriptionById(id uint64) (*entity.Subscription, error)
	UpdateSubscription(subscription *entity.Subscription) (*entity.Subscription, error)
	CancelSubscription(id uint64) error
	GetSubscriptionsByMerchantId(id uint64) ([]*entity.Subscription, error)
}
