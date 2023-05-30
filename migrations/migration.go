package migrations

import (
	"github.com/durotimicodes/trace-backend/helpers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialect/postgres"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=tracebankapp password=postgres sslmode=disabled")
	helpers.HandleErr(err)
	return db
}




