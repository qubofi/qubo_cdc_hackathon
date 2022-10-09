package router

import (
	"qubo/qubo-backend/application"
	"qubo/qubo-backend/infrastructure/service_impl"
	"qubo/qubo-backend/interfaces/handlers"
	"qubo/qubo-backend/interfaces/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRouter(g *gin.RouterGroup) {
	userService := service_impl.NewUserService()
	userApplication := application.NewUserApplication(userService)
	authHandler := handlers.NewAuthHandler(userApplication)

	g.POST("/login", authHandler.Login)
	g.POST("/logout", middleware.AuthMiddleware(), authHandler.Logout)
	g.POST("/refresh", authHandler.Refresh)
}
