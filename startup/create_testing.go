package startup

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"time"

	"github.com/0x6flab/namegenerator"
	"mikel-kunze.com/matchma/database"
)

func AddTesting(userCount int) {

	for range userCount {
		// for generating random names
		nameGen := namegenerator.NewGenerator()
		name := nameGen.Generate()

		// creates a tmp email
		email := name + "@rick-astley.cloud"

		query := "INSERT INTO Users VALUES(DEFAULT, ?, ?, ?);"
		queryArgs := []string{name, "12345", email}

		result := database.ExecuteSQL(query, queryArgs)

		go createRandomPlayerData(result.LastId)
	}

	fmt.Println("Success setting up test setup")
}

// Create random matchdata for the given player
func createRandomPlayerData(userId uint) {

	// first we decide how many matches the player had
	playerTotalMatches := rand.IntN(300)

	for range playerTotalMatches {

		playerPlace := rand.Uint64N(100)
		// getting the max amount of players the given player has killed
		maxKills := 100 - playerPlace

		var playerKills uint64 = 0
		if maxKills != 0 {
			playerKills = rand.Uint64N(maxKills)
		}

		// random player time
		playerTime := rand.IntN(int(playerPlace) + 100)

		query := "INSERT INTO UserMatches VALUES(DEFAULT, ?, ?, ?, ?, ?);"
		queryArgs := []string{strconv.FormatUint(playerPlace, 10), strconv.FormatUint(playerKills, 10), strconv.Itoa(playerTime), time.Now().Format("2006-01-02 15:04:05"), strconv.FormatUint(uint64(userId), 10)}

		database.ExecuteSQL(query, queryArgs)
	}
}
