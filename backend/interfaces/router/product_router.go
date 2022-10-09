package router

import (
	"qubo/qubo-backend/application"
	"qubo/qubo-backend/infrastructure/service_impl"
	"qubo/qubo-backend/interfaces/handlers"

	"github.com/gin-gonic/gin"
)

func ProductRouter(g *gin.RouterGroup) {
	productService := service_impl.NewProductService()
	productApplication := application.NewProductApplication(productService)
	productHandler := handlers.NewProductHandler(productApplication)

	g.POST("/product", productHandler.SaveProduct)
	g.GET("/product/:id", productHandler.GetProductById)
	g.PUT("/product/:id", productHandler.UpdateProduct)
	g.DELETE("/product/:id", productHandler.DeleteProduct)
	g.GET("/product/merchant/:id", productHandler.GetProductsByMerchantId)
}
