package migrations

import (
	"fmt"

	"github.com/durotimicodes/trace-backend/api/database"
	"github.com/durotimicodes/trace-backend/helpers"
	"github.com/durotimicodes/trace-backend/models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func createAccounts() {

	users := []models.User{}

	for i := 0; i < len(users); i++ {
		generatePassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := models.User{
			Username: users[i].Username,
			Email:    users[i].Email,
			Password: generatePassword,
		}
		database.DB.Create(&user)

		account := models.Account{
			Type:    "Savings Account",
			Name:    string(users[i].Username + "'s" + " account"),
			Balance: uint(10000 * int(i+1)),
			UserID:  user.ID,
		}
		database.DB.Create(&account)
	}

	defer database.DB.Close() //close the db connection
}

func Migrate() error {

	User := &models.User{}
	Account := &models.Account{}
	

	err := database.DB.AutoMigrate(&User, &Account)
	if err != nil {
		return fmt.Errorf("error migrating models: %v", err)
	}
	defer database.DB.Close()
	createAccounts()

	return nil

}

func MigrateTranscations() error {

	Transactions := &models.Transaction{}

	err := database.DB.AutoMigrate(&Transactions)
	if err != nil {
		return fmt.Errorf("error migrating models: %v", err)
	}
	defer database.DB.Close()

	return nil

}
