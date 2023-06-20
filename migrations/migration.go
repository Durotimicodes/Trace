package migrations

import (
	"fmt"

	"github.com/durotimicodes/trace-backend/helpers"
	"github.com/durotimicodes/trace-backend/models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)




func createAccounts() {
	db := helpers.ConnectDB()

	//dommy data
	users := []models.User{
		{Username: "Oluwadurotimi", Email: "edmondfagbuyi@gmail.com"},
		{Username: "Ebunoluwa", Email: "omotarfagbuyi@gmail.com"},
	}

	for i:=0 ; i < len(users); i++ {
		generatePassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := models.User{
			Username: users[i].Username,
			Email: users[i].Email,
			Password: generatePassword,
		}
		db.Create(&user)

		account := models.Account{
			Type: "Savings Account",
			Name: string(users[i].Username + "'s" + " account"),
			Balance: uint(10000 * int(i+1)),
			UserID: user.ID,
		}
		db.Create(&account)
	}

	defer db.Close()//close the db connection
}


func Migrate() error {

	User := &models.User{}
	Account := &models.Account{}
	db:= helpers.ConnectDB()

	err := db.AutoMigrate(&User, &Account)
	if err != nil {
		return fmt.Errorf("error migrating models: %v", err)
	}
	defer db.Close()
	createAccounts()

	return nil

}

func MigrateTranscations() {

	Transactions := &models.Transaction{}

	db := helpers.ConnectDB()
	db.AutoMigrate(&Transactions)
	defer db.Close()

}




