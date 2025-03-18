package database

import (
	"log"
)

func RunMigrations() {
	CreateCityTable()
	CreateUsersTable()
	AddIsoToCity()
	log.Println("Migrations completed successfully")

}
