package router

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	api := r.Group("/api")
	UserRouter(api)
	AuthRouter(api)
	MerchantRouter(api)
	ProductRouter(api)
	TransactionRouter(api)
	SubscriptionRouter(api)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}
