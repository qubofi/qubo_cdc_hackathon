package handlers

import (
	"net/http"
	"qubo/qubo-backend/application"
	"qubo/qubo-backend/domain/entity"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SubscriptionHandler struct {
	subscriptionApplicationInterface application.SubscriptionAppInterface
}

func NewSubscriptionHandler(subscriptionApplicationInterface application.SubscriptionAppInterface) *SubscriptionHandler {
	return &SubscriptionHandler{
		subscriptionApplicationInterface: subscriptionApplicationInterface,
	}
}

func (subscriptionHandler *SubscriptionHandler) SaveSubscription(c *gin.Context) {
	var subscription *entity.Subscription
	c.BindJSON(&subscription)
	subscription, err := subscriptionHandler.subscriptionApplicationInterface.SaveSubscription(subscription)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subscription)
}

func (subscriptionHandler *SubscriptionHandler) GetSubscriptionById(c *gin.Context) {
	id, exist := c.Params.Get("id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	subscriptionId, _ := strconv.ParseUint(id, 10, 64)
	subscription, err := subscriptionHandler.subscriptionApplicationInterface.GetSubscriptionById(subscriptionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subscription)
}

func (subscriptionHandler *SubscriptionHandler) UpdateSubscription(c *gin.Context) {
	var subscription *entity.Subscription
	c.BindJSON(&subscription)
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	subscription.ID = id
	subscription, err = subscriptionHandler.subscriptionApplicationInterface.UpdateSubscription(subscription)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subscription)
}

func (subscriptionHandler *SubscriptionHandler) CancelSubscription(c *gin.Context) {
	id, exist := c.Params.Get("id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	subscriptionId, _ := strconv.ParseUint(id, 10, 64)
	err := subscriptionHandler.subscriptionApplicationInterface.CancelSubscription(subscriptionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Subscription cancelled successfully"})
}

func (subscriptionHandler *SubscriptionHandler) GetSubscriptionsByMerchantId(c *gin.Context) {
	id, exist := c.Params.Get("id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	merchantId, _ := strconv.ParseUint(id, 10, 64)
	subscriptions, err := subscriptionHandler.subscriptionApplicationInterface.GetSubscriptionsByMerchantId(merchantId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subscriptions)
}
