package user

import (
	"net/http"

	"github.com/gorilla/websocket"
	"mikel-kunze.com/matchma/logging"
)

type ConnectedUser struct {
	UserId      uint
	Conn        *websocket.Conn
	IsConnected bool
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// TODO: upgrade connection to a websocket
func HandlePlayerJoin(w http.ResponseWriter, r *http.Request) {

	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		logging.Log(Logging.Error, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}

func handleWebsocketConns(ws *websocket.Conn) {}
