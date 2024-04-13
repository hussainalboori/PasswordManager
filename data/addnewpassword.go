package data

import (
    "database/sql"
    //"errors"
)

// InsertPassword inserts a new password into the database
func InsertPassword(userID int, website, username, password string) error {
    // SQL statement to insert a new password
    insertPasswordSQL := `
        INSERT INTO passwords (user_id, website, username, password)
        VALUES (?, ?, ?, ?);
    `

    // Open a connection to the SQLite3 database
    db, err := sql.Open("sqlite3", "data.db")
    if err != nil {
        return err
    }
    defer db.Close()

    // Execute the SQL statement to insert the password
    _, err = db.Exec(insertPasswordSQL, userID, website, username, password)
    if err != nil {
        return err
    }

    return nil
}
