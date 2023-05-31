package helpers

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/jinzhu/gorm"
)

func HandleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func HashAndSalt(pass []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	HandleErr(err)

	return string(hashed)
}

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5433 user=postgres dbname=tracebankapp password=postgres sslmode=disable")
	HandleErr(err)
	return db
}
