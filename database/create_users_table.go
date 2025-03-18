package database

import (
	"fmt"
)

func CreateUsersTable() {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		city_id INT REFERENCES cities(id) ON UPDATE CASCADE  ON DELETE  SET NULL
	);
	`
	err := DB.Exec(query).Error
	if err != nil {
		fmt.Println("Failed to create `users` table:", err)
	} else {
		fmt.Println("`users` table created successfully!")
	}
}
