package user

import (
	"net/http"

	"github.com/gorilla/websocket"
	"mikel-kunze.com/matchma/logging"
)

// TODO: restructure!!

type WSServer struct {
	clients          map[*Client]bool
	broadcast        chan []byte
	registerClient   chan *Client
	unregisterClient chan *Client
}

type Client struct {
	UserId uint
	Conn   *websocket.Conn
	Send   chan []byte
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// TODO
func CreateNewWSServer() *WSServer {
	return &WSServer{}
}

func HandlePlayerJoin(w http.ResponseWriter, r *http.Request) {

	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		logging.Log(logging.Error, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	go handleWebsocketConns(ws)
}

func handleWebsocketConns(ws *websocket.Conn) {}
