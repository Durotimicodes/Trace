package useraccounts

import (
	"github.com/durotimicodes/trace-backend/helpers"
	"github.com/durotimicodes/trace-backend/models"
	"fmt"
)

func updateAccount(id uint, amount int) models.ResponseAccount {
	db := helpers.ConnectDB()

	account := models.Account{}
	responseAcc := models.ResponseAccount{}

	db.Where("id = ?", id).First(&account)
	account.Balance = uint(amount)
	db.Save(&account)

	responseAcc.ID = account.ID
	responseAcc.Name = account.Name
	responseAcc.Balance = int(account.Balance)

	defer db.Close()

	return responseAcc
}

func getAccount(id uint) *models.Account {
	db := helpers.ConnectDB()
	account := &models.Account{}
	if db.Where("id = ?", id).First(&account).RecordNotFound() {
		return nil
	}

	defer db.Close()

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
			return map[string]interface{}{"message": "Account not found"}
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
