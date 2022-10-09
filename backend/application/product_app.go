package application

import (
	"qubo/qubo-backend/domain/entity"
	"qubo/qubo-backend/infrastructure/service_impl"
)

type ProductAppInterface interface {
	SaveProduct(product *entity.Product) (*entity.Product, error)
	GetProductById(id uint64) (*entity.Product, error)
	UpdateProduct(product *entity.Product) (*entity.Product, error)
	DeleteProduct(id uint64) error
	GetProductsByMerchantId(id uint64) ([]*entity.Product, error)
}

type productApp struct {
	productServiceImpl *service_impl.ProductServiceImpl
}

var _ ProductAppInterface = &productApp{}

func NewProductApplication(productServiceImpl *service_impl.ProductServiceImpl) *productApp {
	return &productApp{
		productServiceImpl: productServiceImpl,
	}
}

func (productApp *productApp) SaveProduct(product *entity.Product) (*entity.Product, error) {
	return productApp.productServiceImpl.SaveProduct(product)
}

func (productApp *productApp) GetProductById(id uint64) (*entity.Product, error) {
	return productApp.productServiceImpl.GetProductById(id)
}

func (productApp *productApp) UpdateProduct(product *entity.Product) (*entity.Product, error) {
	return productApp.productServiceImpl.UpdateProduct(product)
}

func (productApp *productApp) DeleteProduct(id uint64) error {
	return productApp.productServiceImpl.DeleteProduct(id)
}

func (productApp *productApp) GetProductsByMerchantId(id uint64) ([]*entity.Product, error) {
	return productApp.productServiceImpl.GetProductsByMerchantId(id)
}
