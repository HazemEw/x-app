package database

import (
	"fmt"
)

func CreateCityTable() {

	query := `
	CREATE TABLE IF NOT EXISTS cities (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL UNIQUE
	);
	`

	err := DB.Exec(query).Error
	if err != nil {
		fmt.Println("Failed to create `cities` table:", err)
	} else {
		fmt.Println("cities` table created successfully!")
	}
}
