package migrations

import (
	"github.com/durotimicodes/trace-backend/helpers"
	"github.com/durotimicodes/trace-backend/models"
	
	_ "github.com/jinzhu/gorm/dialects/postgres"
)




func createAccount() {
	db := helpers.ConnectDB()

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

func Migrate() {

	db:= helpers.ConnectDB()
	db.AutoMigrate(&models.User{}, &models.Account{})
	defer db.Close()
	createAccount()

	

}





