package storage

import "go-api-transaction/models"

// In-memory storage
var Cards = map[string]*models.Card{}
var Transactions = []models.Transaction{}
