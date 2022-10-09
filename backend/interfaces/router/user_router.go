package router

import (
	"qubo/qubo-backend/application"
	"qubo/qubo-backend/infrastructure/service_impl"
	"qubo/qubo-backend/interfaces/handlers"
	"qubo/qubo-backend/interfaces/middleware"

	"github.com/gin-gonic/gin"
)

func UserRouter(g *gin.RouterGroup) {
	userService := service_impl.NewUserService()
	userApplication := application.NewUserApplication(userService)
	userHandler := handlers.NewUserHandler(userApplication)

	g.POST("/users", userHandler.SaveUser)
	g.GET("/users", middleware.AuthMiddleware(), userHandler.GetUser)
}
