package dal

import (
	"qubo/qubo-backend/domain/entity"
	"qubo/qubo-backend/infrastructure/database"
)

func InsertIntoProducts(product *entity.Product) (*entity.Product, error) {
	err := database.Db.QueryRow(`INSERT INTO "products"(merchant_id, product_name, cost, currency, interval, interval_unit) VALUES($1, $2, $3, $4, $5, $6) RETURNING id`,
		product.MerchantId,
		product.ProductName,
		product.Cost,
		product.Currency,
		product.Interval,
		product.IntervalUnit).Scan(&product.ID)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func SelectFromProductsById(id uint64) (*entity.Product, error) {
	product := &entity.Product{}
	err := database.Db.QueryRow(`SELECT id, merchant_id, product_name, cost, currency, interval, interval_unit FROM "products" WHERE id = $1`, id).
		Scan(&product.ID, &product.MerchantId, &product.ProductName, &product.Cost, &product.Currency, &product.Interval, &product.IntervalUnit)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func UpdateProduct(product *entity.Product) (*entity.Product, error) {
	_, err := database.Db.Exec(`UPDATE "products" SET merchant_id = $1, product_name = $2, cost = $3, currency = $4, interval = $5, interval_unit = $6 WHERE id = $7`,
		product.MerchantId,
		product.ProductName,
		product.Cost,
		product.Currency,
		product.Interval,
		product.IntervalUnit,
		product.ID)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func DeleteProduct(id uint64) error {
	_, err := database.Db.Exec(`DELETE FROM "products" WHERE id = $1`, id)

	if err != nil {
		return err
	}

	return nil
}

func SelectFromProductsByMerchantId(id uint64) ([]*entity.Product, error) {
	rows, err := database.Db.Query(`SELECT id, merchant_id, product_name, cost, currency, interval, interval_unit FROM "products" WHERE merchant_id = $1`, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*entity.Product

	for rows.Next() {
		product := &entity.Product{}
		err := rows.Scan(&product.ID, &product.MerchantId, &product.ProductName, &product.Cost, &product.Currency, &product.Interval, &product.IntervalUnit)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}
