package main

import (
	"go-api-transaction/handlers"
	"go-api-transaction/models"
	"go-api-transaction/storage"
	"go-api-transaction/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	
	storage.Cards["4123456789012345"] = &models.Card{
		CardNumber: "4123456789012345",
		CardHolder: "John Doe",
		PinHash:    utils.HashPin("1234"),
		Balance:    1000,
		Status:     "ACTIVE",
	}

	router := gin.Default()

	router.POST("/api/transaction", handlers.ProcessTransaction)
	router.GET("/api/card/balance/:cardNumber", handlers.GetBalance)
	router.GET("/api/card/transactions/:cardNumber", handlers.GetTransactions)

	router.Run(":8080")
}
