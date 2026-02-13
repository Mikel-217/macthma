package user

import (
	"encoding/json"
	"io"
	"net/http"

	"mikel-kunze.com/matchma/database"
	"mikel-kunze.com/matchma/logging"
	matchmastructs "mikel-kunze.com/matchma/matchma_structs"
)

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
	// TODO: Also check if the user email already exists!!

	query := "INSERT INTO Users VALUES(DEFAULT, ?, ?, ?)"
	queryArgs := []string{userST.UserName, userST.UserPW, userST.UserMail}

	database.ExecuteSQL(query, queryArgs)

	// TODO: error check!

	w.WriteHeader(http.StatusOK)
	return
}
