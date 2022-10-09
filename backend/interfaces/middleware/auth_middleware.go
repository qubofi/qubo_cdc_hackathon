package middleware

import (
	"net/http"
	"qubo/qubo-backend/infrastructure/security"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		metadata, err := security.ExtractTokenMetadata(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "unauthorized")
			c.Abort()
			return
		}

		userId, err := security.FetchAuth(metadata)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "unauthorized")
			c.Abort()
			return
		}
		c.Set("userId", userId)
		c.Next()
	}
}
