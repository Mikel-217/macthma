package dbuser

import (
	"mikel-kunze.com/matchma/database"
	"mikel-kunze.com/matchma/logging"
	matchmastructs "mikel-kunze.com/matchma/matchma_structs"
)

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
