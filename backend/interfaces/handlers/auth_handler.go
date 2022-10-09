package handlers

import (
	"fmt"
	"net/http"
	"qubo/qubo-backend/application"
	"qubo/qubo-backend/domain/entity"
	"qubo/qubo-backend/infrastructure/security"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userApplicationInterface application.UserAppInterface
}

func NewAuthHandler(userApplicationInterface application.UserAppInterface) *AuthHandler {
	return &AuthHandler{
		userApplicationInterface: userApplicationInterface,
	}
}

func (authHandler *AuthHandler) Login(c *gin.Context) {
	// user entity that will only contain email and password
	var u *entity.User

	// Bind the request body to the user struct
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the user by email
	user, emailError := authHandler.userApplicationInterface.GetUserByEmail(u.Email)
	if emailError != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email is incorrect"})
		return
	}

	// Check the password
	passwordError := user.CheckPasswordHash(u.Password)
	if passwordError != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Password is incorrect: %s", passwordError.Error())})
		return
	}

	// Create the token
	tokenDetail, tokenErr := security.CreateToken(user.ID)
	if tokenErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": tokenErr.Error()})
		return
	}

	// Save the token in the database
	authError := security.CreateAuth(user.ID, tokenDetail)
	if authError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": authError.Error()})
		return
	}

	// Return the token
	tokens := map[string]string{
		"access_token":  tokenDetail.AccessToken,
		"refresh_token": tokenDetail.RefreshToken,
	}
	c.JSON(http.StatusOK, gin.H{"tokens": tokens})
}

func (authHandler *AuthHandler) Logout(c *gin.Context) {
	// Get the metadata from the request
	metadata, err := security.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		c.Abort()
		return
	}

	// Delete the access token
	deletedUuid, deleteErr := security.DeleteAuth(metadata.AccessUuid)
	if deleteErr != nil || deletedUuid == 0 {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

func (authHandler *AuthHandler) Refresh(c *gin.Context) {
	tokenMap := map[string]string{}
	if err := c.ShouldBindJSON(&tokenMap); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	refreshToken := tokenMap["refresh_token"]

	tokenDetail, err := security.GenerateNewTokenFromRefreshToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	tokens := map[string]string{
		"access_token":  tokenDetail.AccessToken,
		"refresh_token": tokenDetail.RefreshToken,
	}

	c.JSON(http.StatusOK, gin.H{"tokens": tokens})

}
