package matchmaking

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"mikel-kunze.com/matchma/database"
	"mikel-kunze.com/matchma/logging"
	matchmastructs "mikel-kunze.com/matchma/matchma_structs"
)

// If the player has played a round, the new data gets posted here
func HandleNewMatchData(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)

	if err != nil {
		logging.Log(logging.Error, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var match matchmastructs.UserMatches

	if err := json.Unmarshal(body, &match); err != nil {
		logging.Log(logging.Error, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	query := "INSERT INTO UserMatches Values(DEFAULT, ?, ?, ?, ?, ?)"
	queryArgs := []string{strconv.FormatUint(uint64(match.UserPlace), 0), strconv.FormatUint(uint64(match.UserKills), 0), strconv.Itoa(match.UserPlayTime), match.MatchDate.Format("2006-01-02 15:04:05"), strconv.FormatUint(uint64(match.UserId), 0)}

	result := database.ExecuteSQL(query, queryArgs)

	if result.ErrorMsg != nil {
		logging.Log(logging.Error, result.ErrorMsg.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
