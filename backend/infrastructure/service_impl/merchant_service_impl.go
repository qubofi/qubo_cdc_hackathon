package service_impl

import (
	"database/sql"
	"errors"
	"qubo/qubo-backend/domain/entity"
	"qubo/qubo-backend/domain/service"
	"qubo/qubo-backend/infrastructure/database"
	"qubo/qubo-backend/infrastructure/database/dal"
)

type MerchantServiceImpl struct {
	db *sql.DB
}

func NewMerchantService() *MerchantServiceImpl {
	return &MerchantServiceImpl{
		db: database.Db,
	}
}

var _ service.MerchantService = &MerchantServiceImpl{}

func (merchantRepo *MerchantServiceImpl) SaveMerchant(merchant *entity.Merchant) (*entity.Merchant, error) {
	existingMerchant, _ := dal.SelectFromMerchantsByEmail(merchant.Email)
	if existingMerchant != nil {
		return nil, errors.New("email already exists")
	}

	newMerchant, err := dal.InsertIntoMerchants(merchant)

	if err != nil {
		return nil, err
	}

	return newMerchant, nil
}

func (merchantRepo *MerchantServiceImpl) GetMerchantById(id uint64) (*entity.Merchant, error) {
	merchant, err := dal.SelectFromMerchantsById(id)

	if err != nil {
		return nil, err
	}

	return merchant, nil
}

func (merchantRepo *MerchantServiceImpl) UpdateMerchant(merchant *entity.Merchant) (*entity.Merchant, error) {
	updatedMerchant, err := dal.UpdateMerchant(merchant)

	if err != nil {
		return nil, err
	}

	return updatedMerchant, nil
}
