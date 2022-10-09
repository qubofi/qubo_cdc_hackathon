package router

import (
	"qubo/qubo-backend/application"
	"qubo/qubo-backend/infrastructure/service_impl"
	"qubo/qubo-backend/interfaces/handlers"

	"github.com/gin-gonic/gin"
)

func SubscriptionRouter(g *gin.RouterGroup) {
	subscriptionService := service_impl.NewSubscriptionService()
	subscriptionApplication := application.NewSubscriptionApplication(subscriptionService)
	subscriptionHandler := handlers.NewSubscriptionHandler(subscriptionApplication)

	g.POST("/subscription", subscriptionHandler.SaveSubscription)
	g.GET("/subscription/:id", subscriptionHandler.GetSubscriptionById)
	g.PUT("/subscription/:id", subscriptionHandler.UpdateSubscription)
	g.PATCH("/subscription/:id/cancel/", subscriptionHandler.CancelSubscription)
	g.GET("/subscription/merchant/:id", subscriptionHandler.GetSubscriptionsByMerchantId)
}
