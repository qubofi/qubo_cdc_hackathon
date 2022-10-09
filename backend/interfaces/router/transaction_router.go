package router

import (
	"qubo/qubo-backend/application"
	"qubo/qubo-backend/infrastructure/service_impl"
	"qubo/qubo-backend/interfaces/handlers"

	"github.com/gin-gonic/gin"
)

func TransactionRouter(g *gin.RouterGroup) {
	transactionService := service_impl.NewTransactionService()
	transactionApplication := application.NewTransactionApplication(transactionService)
	transactionHandler := handlers.NewTransactionHandler(transactionApplication)

	g.POST("/transaction", transactionHandler.SaveTransaction)
	g.GET("/transaction/:id", transactionHandler.GetTransactionById)
	g.GET("/transaction/subscription/:id", transactionHandler.GetTransactionsBySubscriptionId)
}
