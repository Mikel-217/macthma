package dbuser

import (
	"mikel-kunze.com/matchma/database"
	"mikel-kunze.com/matchma/logging"
	matchmastructs "mikel-kunze.com/matchma/matchma_structs"
)

// Gets a user by the given mail
func GetUserByMail(mail string) *matchmastructs.UserStruct {

	db := database.CreateDBConn()

	if db == nil {
		return nil
	}

	var user matchmastructs.UserStruct

	if err := db.QueryRow("SELECT * FROM Users WHERE UserMail = ?;", mail).Scan(&user); err != nil {
		logging.Log(logging.Error, err.Error())
		return nil
	}

	return &user
}

// Gets a user by the given username
func GetUserByName(userName string) *matchmastructs.UserStruct {

	db := database.CreateDBConn()

	if db == nil {
		return nil
	}

	var user matchmastructs.UserStruct

	if err := db.QueryRow("SELECT * FROM Users WHERE UserName = ?;", userName).Scan(&user); err != nil {
		logging.Log(logging.Error, err.Error())
		return nil
	}

	return &user
}
