package user

import (
	"encoding/json"
	"io"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"mikel-kunze.com/matchma/database"
	dbuser "mikel-kunze.com/matchma/database/db_user"
	"mikel-kunze.com/matchma/logging"
	matchmastructs "mikel-kunze.com/matchma/matchma_structs"
)

// handels the registration of an user
func HandleUserRegister(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		logging.Log(logging.Error, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var userST matchmastructs.UserStruct

	if err := json.Unmarshal(body, &userST); err != nil {
		logging.Log(logging.Error, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// checks if the user is already registered
	dbUser := dbuser.GetUserByMail(userST.UserMail)

	if dbUser.UserMail == userST.UserMail {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// hashes the pw from the given user
	hashedPW, err := bcrypt.GenerateFromPassword([]byte(userST.UserPW), 14)

	if err != nil {
		logging.Log(logging.Error, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userST.UserPW = string(hashedPW)

	query := "INSERT INTO Users VALUES(DEFAULT, ?, ?, ?);"
	queryArgs := []string{userST.UserName, userST.UserPW, userST.UserMail}

	result := database.ExecuteSQL(query, queryArgs)

	// checks for an error
	if result.ErrorMsg != nil {
		logging.Log(logging.Error, result.ErrorMsg.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// everything is ok :)
	w.WriteHeader(http.StatusOK)
}
