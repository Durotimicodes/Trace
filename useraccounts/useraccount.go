package useraccounts

import (
	"fmt"

	"github.com/durotimicodes/trace-backend/api/database"
	"github.com/durotimicodes/trace-backend/helpers"
	"github.com/durotimicodes/trace-backend/models"
	"github.com/durotimicodes/trace-backend/transactions"
)

func updateAccount(id uint, amount int) models.ResponseAccount {

	account := models.Account{}
	responseAcc := models.ResponseAccount{}

	database.DB.Where("id = ?", id).First(&account)
	account.Balance = uint(amount)
	database.DB.Save(&account)

	responseAcc.ID = account.ID
	responseAcc.Name = account.Name
	responseAcc.Balance = int(account.Balance)

	return responseAcc
}

func getAccount(id uint) *models.Account {
	
	account := &models.Account{}
	if database.DB.Where("id = ?", id).First(&account).RecordNotFound() {
		return nil
	}
	return account
}

func Transaction(userId uint, from uint, to uint, amount int, jwt string) map[string]interface{} {

	userIdString := fmt.Sprintf("%v", userId)
	isValid := helpers.ValidateToken(userIdString, jwt)

	if isValid {
		fromAccount := getAccount(from)
		toAccount := getAccount(to)

		if fromAccount == nil || toAccount == nil {
			return map[string]interface{}{"message": "Account not found"}
		} else if fromAccount.UserID != userId {
			return map[string]interface{}{"message": "Wrong user account"}
		} else if int(fromAccount.Balance) < amount {
			return map[string]interface{}{"message": "Insufficient Account Balance"}
		}

		updatedAccount := updateAccount(from, int(fromAccount.Balance) - amount)
		updateAccount(to, int(toAccount.Balance) + amount)

		transactions.CreateTransaction(from, to, amount)
		var response = map[string]interface{}{"message": "Transaction successfull"}
		response["data"] = updatedAccount

		return response

	} else {
		return map[string]interface{}{"message": "Not valid token"}
	}
}
