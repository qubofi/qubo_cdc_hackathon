package service

import (
	"qubo/qubo-backend/domain/entity"
)

type ProductService interface {
	SaveProduct(product *entity.Product) (*entity.Product, error)
	GetProductById(id uint64) (*entity.Product, error)
	UpdateProduct(product *entity.Product) (*entity.Product, error)
	DeleteProduct(id uint64) error
	GetProductsByMerchantId(id uint64) ([]*entity.Product, error)
}
