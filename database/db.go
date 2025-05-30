package database

import (
	"database/sql"
	"fmt"
)

func CreateTable(db *sql.DB) error {
	_, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS quotes(
    id SERIAL PRIMARY KEY,
    Author VARCHAR(255) NOT NULL,
    Quote VARCHAR(255) NOT NULL);
`)
	if err != nil {
		return fmt.Errorf("error creating quotes: %v", err)
	}
	return nil
}
