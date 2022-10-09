package handlers

import (
	"net/http"
	"qubo/qubo-backend/application"
	"qubo/qubo-backend/domain/entity"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productApplicationInterface application.ProductAppInterface
}

func NewProductHandler(productApplicationInterface application.ProductAppInterface) *ProductHandler {
	return &ProductHandler{
		productApplicationInterface: productApplicationInterface,
	}
}

func (productHandler *ProductHandler) SaveProduct(c *gin.Context) {
	var product *entity.Product
	c.BindJSON(&product)
	product, err := productHandler.productApplicationInterface.SaveProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (productHandler *ProductHandler) GetProductById(c *gin.Context) {
	id, exist := c.Params.Get("id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	productId, _ := strconv.ParseUint(id, 10, 64)
	product, err := productHandler.productApplicationInterface.GetProductById(productId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (productHandler *ProductHandler) UpdateProduct(c *gin.Context) {
	var product *entity.Product
	c.BindJSON(&product)
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	product.ID = id
	product, err = productHandler.productApplicationInterface.UpdateProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (productHandler *ProductHandler) DeleteProduct(c *gin.Context) {
	id, exist := c.Params.Get("id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	productId, _ := strconv.ParseUint(id, 10, 64)
	err := productHandler.productApplicationInterface.DeleteProduct(productId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

func (productHandler *ProductHandler) GetProductsByMerchantId(c *gin.Context) {
	id, exist := c.Params.Get("id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	merchantId, _ := strconv.ParseUint(id, 10, 64)
	products, err := productHandler.productApplicationInterface.GetProductsByMerchantId(merchantId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}
