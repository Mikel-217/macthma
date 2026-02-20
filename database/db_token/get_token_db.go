package dbtoken

import (
	"mikel-kunze.com/matchma/database"
	"mikel-kunze.com/matchma/logging"
)

// If token isnt there it returns false
func IsTokenThere(searchToken string) bool {

	db := database.CreateDBConn()

	if db == nil {
		return false
	}

	result := db.QueryRow("SELECT * FROM AccessTokens WHERE TokenVal = ?;", searchToken)

	if result.Err() != nil {
		logging.Log(logging.Error, result.Err().Error())
		return false
	}

	if result == nil {
		return false
	}

	return true
}
