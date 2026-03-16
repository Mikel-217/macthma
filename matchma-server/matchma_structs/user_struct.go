package matchmastructs

import (
	"time"
)

type UserStruct struct {
	UserId   uint
	UserName string
	UserPW   string // always hash this!!
	UserMail string
}

// just for the code -> getting all data from a user
type UserInformation struct {
	TotalKills    uint
	TotalPlayTime int
	TotalWins     uint
	UserId        uint
}

// for storring user info from a specific match
type UserMatches struct {
	MatchId      uint
	UserPlace    uint
	UserKills    uint
	UserPlayTime int // represents the minutes how long the player was in the given match
	MatchDate    time.Time
	UserId       uint // foreign key to Users Table
}

func (u *UserInformation) GetSkillScore() float64 {
	return float64(u.TotalWins*10) + float64(u.TotalKills)
}
