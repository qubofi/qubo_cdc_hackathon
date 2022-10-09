package router

import (
	"qubo/qubo-backend/application"
	"qubo/qubo-backend/infrastructure/service_impl"
	"qubo/qubo-backend/interfaces/handlers"

	"github.com/gin-gonic/gin"
)

func MerchantRouter(g *gin.RouterGroup) {
	merchantService := service_impl.NewMerchantService()
	merchantApplication := application.NewMerchantApplication(merchantService)
	merchantHandler := handlers.NewMerchantHandler(merchantApplication)

	g.POST("/merchant", merchantHandler.SaveMerchant)
	g.GET("/merchant/:id", merchantHandler.GetMerchantById)
	g.PUT("/merchant/:id", merchantHandler.UpdateMerchant)
}
