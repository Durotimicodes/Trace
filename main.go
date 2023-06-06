package main

import (
	"github.com/durotimicodes/trace-backend/api"
	"github.com/durotimicodes/trace-backend/migrations"
)

func main() {

	migrations.Migrate()
	api.StartApi()

}
