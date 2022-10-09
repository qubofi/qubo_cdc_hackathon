package handlers

import (
	"net/http"
	"qubo/qubo-backend/application"
	"qubo/qubo-backend/domain/entity"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MerchantHandler struct {
	merchantApplicationInterface application.MerchantAppInterface
}

func NewMerchantHandler(merchantApplicationInterface application.MerchantAppInterface) *MerchantHandler {
	return &MerchantHandler{
		merchantApplicationInterface: merchantApplicationInterface,
	}
}

func (merchantHandler *MerchantHandler) SaveMerchant(c *gin.Context) {
	var merchant *entity.Merchant
	c.BindJSON(&merchant)
	merchant, err := merchantHandler.merchantApplicationInterface.SaveMerchant(merchant)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, merchant)
}

func (merchantHandler *MerchantHandler) GetMerchantById(c *gin.Context) {
	id, exist := c.Params.Get("id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	merchantId, _ := strconv.ParseUint(id, 10, 64)
	merchant, err := merchantHandler.merchantApplicationInterface.GetMerchantById(merchantId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, merchant)
}

func (merchantHandler *MerchantHandler) UpdateMerchant(c *gin.Context) {
	var merchant *entity.Merchant
	c.BindJSON(&merchant)
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	merchant.ID = id
	merchant, err = merchantHandler.merchantApplicationInterface.UpdateMerchant(merchant)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, merchant)
}
