package database

import (
	"fmt"
)

func AddIsoToCity() {

	query := `
	ALTER TABLE cities
	ADD COLUMN iso VARCHAR(2)  UNIQUE;
	`

	err := DB.Exec(query).Error
	if err != nil {
		fmt.Println("Failed to add `iso` column in `cities` table:", err)
	} else {
		fmt.Println("`iso` column added successfully in `cities` table!")
	}
}
