package dal

import (
	"qubo/qubo-backend/domain/entity"
	"qubo/qubo-backend/infrastructure/database"
)

func InsertIntoMerchants(merchant *entity.Merchant) (*entity.Merchant, error) {
	err := database.Db.QueryRow(`INSERT INTO "merchants"(merchant_name, wallet_address, email) VALUES($1, $2, $3) RETURNING id`,
		merchant.MerchantName,
		merchant.Email,
		merchant.WalletAddress).Scan(&merchant.ID)

	if err != nil {
		return nil, err
	}

	return merchant, nil
}

func SelectFromMerchantsById(id uint64) (*entity.Merchant, error) {
	merchant := &entity.Merchant{}
	err := database.Db.QueryRow(`SELECT id, merchant_name, wallet_address, email FROM "merchants" WHERE id = $1`, id).
		Scan(&merchant.ID, &merchant.MerchantName, &merchant.WalletAddress, &merchant.Email)

	if err != nil {
		return nil, err
	}

	return merchant, nil
}

func SelectFromMerchantsByEmail(email string) (*entity.Merchant, error) {
	merchant := &entity.Merchant{}
	err := database.Db.QueryRow(`SELECT id, merchant_name, wallet_address, email FROM "merchants" WHERE email = $1`, email).
		Scan(&merchant.ID, &merchant.MerchantName, &merchant.WalletAddress, &merchant.Email)

	if err != nil {
		return nil, err
	}

	return merchant, nil
}

func UpdateMerchant(merchant *entity.Merchant) (*entity.Merchant, error) {
	_, err := database.Db.Exec(`UPDATE "merchants" SET merchant_name = $1, wallet_address = $2, email = $3 WHERE id = $4`,
		merchant.MerchantName,
		merchant.WalletAddress,
		merchant.Email,
		merchant.ID)

	if err != nil {
		return nil, err
	}

	return merchant, nil
}
