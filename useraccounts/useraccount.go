package useraccounts

import (
	"github.com/durotimicodes/trace-backend/helpers"
	"github.com/durotimicodes/trace-backend/models"
)

func updateAccount(id uint, amount int) {
	db := helpers.ConnectDB()

	db.Model(&models.Account{}).Where("id = ?", id).Update("balance", amount)
	defer db.Close()
}