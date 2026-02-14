package database

import (
	"database/sql"
	"os"

	"mikel-kunze.com/matchma/logging"
)

// Creates a new database connection
func CreateDBConn() *sql.DB {

	db, err := sql.Open("mysql", os.Getenv(""))

	if err != nil {
		logging.Log(logging.Error, err.Error())
		return nil
	}

	return db
}
