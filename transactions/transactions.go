package transactions

import (
	"github.com/durotimicodes/trace-backend/api/database"
	"github.com/durotimicodes/trace-backend/helpers"
	"github.com/durotimicodes/trace-backend/models"
)

func CreateTransaction(From uint, To uint, Amount int) {

	transaction := &models.Transaction{
		From:   From,
		To:     To,
		Amount: Amount,
	}

	database.DB.Create(&transaction)

}

func GetTransactionsByAccount(id uint) []models.ResponseTransaction {

	transaction := []models.ResponseTransaction{}

	database.DB.Table("transactions").Select("id, transactions.from, transactions.to, amount").Where(models.Transaction{
		From: id}).Or(models.Transaction{To: id}).Scan(&transaction)

	return transaction
}

func GetMyTransactions(id string, jwt string) map[string]interface{} {

	isValid := helpers.ValidateToken(id, jwt)

	if isValid {
		accounts := []models.ResponseAccount{}

		database.DB.Table("accounts").Select("id, name, balance").Where("user_id = ?", id).Scan(&accounts)

		transactions := []models.ResponseTransaction{}

		for i := 0; i < len(accounts); i++ {
			accTransactions := GetTransactionsByAccount(accounts[i].ID)
			transactions = append(transactions, accTransactions...)
		}

		var response = map[string]interface{}{"message": "all is fine"}
		response["data"] = transactions
		return response

	} else {
		return map[string]interface{}{"message": "Not valid token"}
	}

}
