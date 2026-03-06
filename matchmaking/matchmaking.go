package matchmaking

import (
	dbuser "mikel-kunze.com/matchma/database/db_user"
	matchmastructs "mikel-kunze.com/matchma/matchma_structs"
)

type Lobby struct {
	Cliens      []uint // The ids of all users that get to the lobby
	LobbyStatus int
}

// Gets called by the Run() func if player count == 10
// if testing mode -> this func will be called if a new player has joined the websocket server
func (ws *WSServer) CreateLobbys() []*Lobby {

	if ws.IsTesting {
		sortingPlayer(dbuser.GetAllUserInfoForTesting())
	} else {
		// TODO: is not implemented yet
		dbuser.GetAllUserInfo()
	}

	return nil
}

// returns the lobbys for the given skill lvls
func sortingPlayer(userInfos []matchmastructs.UserInformation) map[uint][]uint {

	return nil
}
