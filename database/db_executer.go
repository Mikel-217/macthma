package database

import (
	"errors"

	"mikel-kunze.com/matchma/logging"
)

// NOTE:
// Only use this for executing DELETE, UPDATE or INSERT commads!!

type result struct {
	lastId   uint
	errorMsg error
}

// This func needs the complete SQL-Statement as a string
// It returns an struct with the last Id and an error to indicate success
func ExecuteSQL(sqlQuery string) *result {

	db := CreateDBConn()

	if db == nil {

		return &result{
			lastId:   0,
			errorMsg: errors.New("DB error"),
		}
	}

	defer db.Close()

	queryResult, err := db.Exec(sqlQuery)

	if err != nil {
		logging.Log(logging.Error, err.Error())
		return &result{
			lastId:   0,
			errorMsg: err,
		}
	}

	Id, _ := queryResult.LastInsertId()

	return &result{
		lastId:   Id,
		errorMsg: nil,
	}
}
