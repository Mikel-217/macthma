package matchmaking

import (
	"sort"

	dbuser "mikel-kunze.com/matchma/database/db_user"
	matchmastructs "mikel-kunze.com/matchma/matchma_structs"
)

type Lobby struct {
	Cliens      []uint // The ids of all users that get to the lobby
	LobbyStatus int
}

// Gets called by the Run() func if player count == 20
// if testing mode -> this func will be called if a new player has joined the websocket server
func (ws *WSServer) CreateLobbys() {

	ws.Broadcast <- []byte("Creating lobbys...")

	if ws.IsTesting {
		// TODO
		sortingPlayer(dbuser.GetAllUserInfoForTesting())
		ws.Broadcast <- []byte("Creating lobbys done")
	} else {
		// TODO: is not implemented yet
		ws.Broadcast <- []byte("not implemented yet...")
		dbuser.GetAllUserInfo()
	}
}

// returns the lobbys for the given skill lvls
// first uint is the skill lvl and the other are the given user Id
func sortingPlayer(userInfos []matchmastructs.UserInformation) map[uint][]uint {

	// Sorts the given players for theyre skill
	sort.Slice(userInfos, func(i, j int) bool {
		return userInfos[i].GetSkillScore() > userInfos[j].GetSkillScore()
	})

	lobbys := make(map[uint][]uint, 100)

	keys := make([]uint, 0, len(lobbys))
	for k := range lobbys {
		keys = append(keys, k)
	}

	for j := range lobbys {
		var usersForLobby []uint
		var count int
		for i := range userInfos {
			count++
			usersForLobby = append(usersForLobby, userInfos[i].UserId)

			if count == 10 {
				lobbys[j] = usersForLobby
				count = 0
			}
		}
	}

	return lobbys
}
