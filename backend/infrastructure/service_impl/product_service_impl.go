package service_impl

import (
	"database/sql"
	"qubo/qubo-backend/domain/entity"
	"qubo/qubo-backend/domain/service"
	"qubo/qubo-backend/infrastructure/database"
	"qubo/qubo-backend/infrastructure/database/dal"
)

type ProductServiceImpl struct {
	db *sql.DB
}

func NewProductService() *ProductServiceImpl {
	return &ProductServiceImpl{
		db: database.Db,
	}
}

var _ service.ProductService = &ProductServiceImpl{}

func (productRepo *ProductServiceImpl) SaveProduct(product *entity.Product) (*entity.Product, error) {
	newProduct, err := dal.InsertIntoProducts(product)

	if err != nil {
		return nil, err
	}

	return newProduct, nil
}

func (productRepo *ProductServiceImpl) GetProductById(id uint64) (*entity.Product, error) {
	product, err := dal.SelectFromProductsById(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (productRepo *ProductServiceImpl) UpdateProduct(product *entity.Product) (*entity.Product, error) {
	updatedProduct, err := dal.UpdateProduct(product)

	if err != nil {
		return nil, err
	}

	return updatedProduct, nil
}

func (productRepo *ProductServiceImpl) DeleteProduct(id uint64) error {
	err := dal.DeleteProduct(id)

	if err != nil {
		return err
	}

	return nil
}

func (productRepo *ProductServiceImpl) GetProductsByMerchantId(id uint64) ([]*entity.Product, error) {
	products, err := dal.SelectFromProductsByMerchantId(id)

	if err != nil {
		return nil, err
	}

	return products, nil
}
