package data

import (
	"database/sql"
	"log"
)

func InsertUser(username, email, passwordHash string, key []byte) error {
	// SQL statement to insert a new user
	insertUserSQL := `
		INSERT INTO users (username, email, password_hash, key)
		VALUES (?, ?, ?, ?);
	`

	_ = CreateDatabaseIfNotExists("")
	// Open a connection to the SQLite database
	db, err := sql.Open("sqlite", GetDBPath())
	if err != nil {
		log.Printf("Error opening database: %v", err)
		return err
	}
	defer db.Close()

	// Execute the SQL statement to insert the user
	_, err = db.Exec(insertUserSQL, username, email, passwordHash, key)
	if err != nil {
		return err
	}

	return nil
}
