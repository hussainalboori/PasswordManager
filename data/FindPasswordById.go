package data

import (
	"database/sql"
	"log"
)

// DeletePasswordByID deletes password data from the database by password ID.
func DeletePasswordByID(passwordID int) error {
	// SQL statement to delete password data by ID
	deletePasswordSQL := `
        DELETE FROM passwords WHERE id = ?;
    `

	// Open a connection to the SQLite3 database
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		log.Printf("Error opening database: %v", err)
		return err
	}
	defer db.Close()

	// Execute the SQL statement to delete password data by ID
	_, err = db.Exec(deletePasswordSQL, passwordID)
	if err != nil {
		log.Printf("Error deleting password data: %v", err)
		return err
	}

	return nil
}
