package handlers

import (
	"fmt"
	"net/http"
	"time"

	"go-api-transaction/models"
	"go-api-transaction/storage"
	"go-api-transaction/utils"

	"github.com/gin-gonic/gin"
)

type TransactionRequest struct {
	CardNumber string  `json:"cardNumber"`
	Pin        string  `json:"pin"`
	Type       string  `json:"type"`
	Amount     float64 `json:"amount"`
}

func ProcessTransaction(c *gin.Context) {
	var req TransactionRequest

	
	if err := c.ShouldBindJSON(&req); err != nil {
		respondWithError(c, req, "FAILED", "", "Invalid request format")
		return
	}

	
	card, exists := storage.Cards[req.CardNumber]
	if !exists || card.Status != "ACTIVE" {
		respondWithError(c, req, "FAILED", "05", "Invalid card")
		return
	}

	
	if utils.HashPin(req.Pin) != card.PinHash {
		respondWithError(c, req, "FAILED", "06", "Invalid PIN")
		return
	}

	
	if req.Amount <= 0 {
		respondWithError(c, req, "FAILED", "", "Invalid amount")
		return
	}

	
	switch req.Type {

	case "withdraw":
		if card.Balance < req.Amount {
			respondWithError(c, req, "FAILED", "99", "Insufficient balance")
			return
		}
		card.Balance -= req.Amount

	case "topup":
		card.Balance += req.Amount

	default:
		respondWithError(c, req, "FAILED", "", "Invalid transaction type")
		return
	}

	
	storage.Cards[req.CardNumber] = card

	
	logTransaction(req.CardNumber, req.Type, req.Amount, "SUCCESS")

	
	c.JSON(http.StatusOK, gin.H{
		"status":   "SUCCESS",
		"respCode": "00",
		"balance":  card.Balance,
	})
}


func respondWithError(c *gin.Context, req TransactionRequest, status, code, message string) {
	logTransaction(req.CardNumber, req.Type, req.Amount, status)

	response := gin.H{
		"status":  status,
		"message": message,
	}

	if code != "" {
		response["respCode"] = code
	}

	c.JSON(http.StatusBadRequest, response)
}


func logTransaction(cardNumber, txnType string, amount float64, status string) {
	txn := models.Transaction{
		TransactionID: fmt.Sprintf("%d", time.Now().UnixNano()),
		CardNumber:    cardNumber,
		Type:          txnType,
		Amount:        amount,
		Status:        status,
		Timestamp:     time.Now().Format(time.RFC3339),
	}

	storage.Transactions = append(storage.Transactions, txn)
}
