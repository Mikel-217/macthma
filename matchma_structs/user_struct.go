package matchmastructs

import "time"

type UserStruct struct {
	UserId   uint
	UserName string
	UserMail string
	UserPW   string // always hash this!!
}

// just for the code -> getting all data from a user
type UserInformation struct {
	TotalKills    uint
	TotalPlayTime time.Duration
	TotalWins     uint
	User          UserStruct
}

// for storring user info from a specific match
type UserMatches struct {
	MatchId      uint
	UserPlace    uint
	UserKills    uint
	UserPlayTime time.Duration
	MatchDate    time.Time
	UserId       uint // foreign key to Users Table
}
