package startup

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"mikel-kunze.com/matchma/database"
	"mikel-kunze.com/matchma/logging"
)

type dbsetup struct {
	TableName string `json:"table-name"`
	Command   string `json:"sql-command"`
}

// Creates all tables and returns true if successfull
func CreateTables() bool {

	db := database.CreateDBConn()

	if db == nil {
		logging.Log(logging.Error, "Database connection failed")
		panic("db Cannot start. See logs")
	}

	currPath, _ := os.Getwd()
	path := path.Join(currPath, "/startup/tables.json")

	content, err := os.ReadFile(path)

	if err != nil {
		logging.Log(logging.Error, err.Error())
		panic("Cannot start. See logs")
	}

	var commands []dbsetup

	if err := json.Unmarshal(content, &commands); err != nil {
		logging.Log(logging.Error, err.Error())
		panic("Cannot start. See logs")
	}

	for i, _ := range commands {
		result := database.ExecuteSQL(commands[i].Command, nil)

		if result.ErrorMsg != nil {
			logging.Log(logging.Error, result.ErrorMsg.Error())
			fmt.Println(result.ErrorMsg)
		}
	}

	return true
}
