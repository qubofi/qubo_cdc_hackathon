package application

import (
	"qubo/qubo-backend/domain/entity"
	"qubo/qubo-backend/infrastructure/service_impl"
)

type MerchantAppInterface interface {
	SaveMerchant(merchant *entity.Merchant) (*entity.Merchant, error)
	GetMerchantById(id uint64) (*entity.Merchant, error)
	UpdateMerchant(merchant *entity.Merchant) (*entity.Merchant, error)
}

type merchantApp struct {
	merchantServiceImpl *service_impl.MerchantServiceImpl
}

var _ MerchantAppInterface = &merchantApp{}

func NewMerchantApplication(merchantServiceImpl *service_impl.MerchantServiceImpl) *merchantApp {
	return &merchantApp{
		merchantServiceImpl: merchantServiceImpl,
	}
}

func (merchantApp *merchantApp) SaveMerchant(merchant *entity.Merchant) (*entity.Merchant, error) {
	return merchantApp.merchantServiceImpl.SaveMerchant(merchant)
}

func (merchantApp *merchantApp) GetMerchantById(id uint64) (*entity.Merchant, error) {
	return merchantApp.merchantServiceImpl.GetMerchantById(id)
}

func (merchantApp *merchantApp) UpdateMerchant(merchant *entity.Merchant) (*entity.Merchant, error) {
	return merchantApp.merchantServiceImpl.UpdateMerchant(merchant)
}
