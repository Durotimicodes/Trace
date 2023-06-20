package transactions

import (
	"github.com/durotimicodes/trace-backend/helpers"
	"github.com/durotimicodes/trace-backend/models"
)

func CreateTransaction(From uint, To uint, Amount int) {
	db := helpers.ConnectDB()

	transaction := &models.Transaction{
		From:   From,
		To:     To,
		Amount: Amount,
	}

	db.Create(&transaction)

	defer db.Close()

}
