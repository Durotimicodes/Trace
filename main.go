package main

import (
	"github.com/durotimicodes/trace-backend/api"
	"github.com/durotimicodes/trace-backend/api/database"
)

func main() {
	
	database.InitDatabase()
	api.StartApi()

}
