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

	defer db.Close()

	var user matchmastructs.UserStruct

	if err := db.QueryRow("SELECT * FROM Users WHERE UserMail = ?;", mail).Scan(&user.UserId, &user.UserName, &user.UserPW, &user.UserMail); err != nil {
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

	defer db.Close()

	var user matchmastructs.UserStruct

	if err := db.QueryRow("SELECT * FROM Users WHERE UserName = ?;", userName).Scan(&user.UserId, &user.UserName, &user.UserPW, &user.UserMail); err != nil {
		logging.Log(logging.Error, err.Error())
		return nil
	}

	return &user
}

// NOTE: This func is only for testing!!
func GetAllUserInfoForTesting() []matchmastructs.UserInformation {

	db := database.CreateDBConn()

	if db == nil {
		return nil
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM UserMatches;")

	if err != nil {
		logging.Log(logging.Error, err.Error())
		return nil
	}

	var matchData = make([]matchmastructs.UserInformation, 1000)

	for rows.Next() {

		userMatchData := matchmastructs.UserMatches{}

		if err := rows.Scan(&userMatchData.MatchId, &userMatchData.UserPlace, &userMatchData.UserKills, &userMatchData.UserPlayTime, &userMatchData.MatchDate, &userMatchData.UserId); err != nil {
			logging.Log(logging.Error, err.Error())
			continue
		}

		for i := range matchData {

			if matchData[i].UserId != userMatchData.UserId {
				matchData = append(matchData, matchmastructs.UserInformation{TotalKills: userMatchData.UserKills, TotalPlayTime: userMatchData.UserPlayTime, TotalWins: uint(IsUserWin(userMatchData.UserPlace)), UserId: userMatchData.UserId})
			} else {
				matchData[i].TotalKills += userMatchData.UserKills
				matchData[i].TotalPlayTime += userMatchData.UserPlayTime
				matchData[i].TotalWins += uint(IsUserWin(userMatchData.UserPlace))
			}
		}
	}

	return matchData
}

// TODO: implement this func for prod
func GetAllUserInfo() {}
