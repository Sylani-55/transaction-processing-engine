package handlers

import (
	"net/http"

	"go-api-transaction/models"
	"go-api-transaction/storage"

	"github.com/gin-gonic/gin"
)

func GetBalance(c *gin.Context) {
	cardNumber := c.Param("cardNumber")

	card, exists := storage.Cards[cardNumber]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Card not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"balance": card.Balance,
	})
}

func GetTransactions(c *gin.Context) {
	cardNumber := c.Param("cardNumber")

	var result []models.Transaction

	for _, t := range storage.Transactions {
		if t.CardNumber == cardNumber {
			result = append(result, t)
		}
	}

	c.JSON(http.StatusOK, result)
}
