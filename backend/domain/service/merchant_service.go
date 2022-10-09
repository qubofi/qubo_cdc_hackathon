package service

import "qubo/qubo-backend/domain/entity"

type MerchantService interface {
	SaveMerchant(merchant *entity.Merchant) (*entity.Merchant, error)
	GetMerchantById(id uint64) (*entity.Merchant, error)
	UpdateMerchant(merchant *entity.Merchant) (*entity.Merchant, error)
}
