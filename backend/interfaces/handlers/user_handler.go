package handlers

import (
	"net/http"
	"qubo/qubo-backend/application"
	"qubo/qubo-backend/domain/entity"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userApplicationInterface application.UserAppInterface
}

func NewUserHandler(userApplicationInterface application.UserAppInterface) *UserHandler {
	return &UserHandler{
		userApplicationInterface: userApplicationInterface,
	}
}

func (userHandler *UserHandler) SaveUser(c *gin.Context) {
	var user *entity.User
	c.BindJSON(&user)
	user, err := userHandler.userApplicationInterface.SaveUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user.Desensitized())
}

func (userHandler *UserHandler) GetUser(c *gin.Context) {
	id, exist := c.Get("userId")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	user, err := userHandler.userApplicationInterface.GetUserById(id.(uint64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user.Desensitized())
}
