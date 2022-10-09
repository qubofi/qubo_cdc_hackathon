package handlers

import (
	"net/http"
	"qubo/qubo-backend/application"
	"qubo/qubo-backend/domain/entity"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	transactionApplicationInterface application.TransactionAppInterface
}

func NewTransactionHandler(transactionApplicationInterface application.TransactionAppInterface) *TransactionHandler {
	return &TransactionHandler{
		transactionApplicationInterface: transactionApplicationInterface,
	}
}

func (transactionHandler *TransactionHandler) SaveTransaction(c *gin.Context) {
	var transaction *entity.Transaction
	c.BindJSON(&transaction)
	transaction, err := transactionHandler.transactionApplicationInterface.SaveTransaction(transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transaction)
}

func (transactionHandler *TransactionHandler) GetTransactionById(c *gin.Context) {
	id, exist := c.Params.Get("id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	transactionId, _ := strconv.ParseUint(id, 10, 64)
	transaction, err := transactionHandler.transactionApplicationInterface.GetTransactionById(transactionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transaction)
}

func (transactionHandler *TransactionHandler) GetTransactionsBySubscriptionId(c *gin.Context) {
	id, exist := c.Params.Get("id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	subscriptionId, _ := strconv.ParseUint(id, 10, 64)
	transactions, err := transactionHandler.transactionApplicationInterface.GetTransactionsBySubscriptionId(subscriptionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transactions)
}
