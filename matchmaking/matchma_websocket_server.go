package matchmaking

import (
	"net/http"

	"github.com/gorilla/websocket"
	"mikel-kunze.com/matchma/logging"
)

type WSServer struct {
	Clients          map[*Client]bool
	Broadcast        chan []byte
	registerClient   chan *Client
	unregisterClient chan *Client
	IsTesting        bool
}

type Client struct {
	UserId uint
	Conn   *websocket.Conn
	Send   chan []byte
}

// Creates a new Websocket server
func CreateNewWSServer(testing bool) *WSServer {
	return &WSServer{
		Clients:          make(map[*Client]bool, 200),
		Broadcast:        make(chan []byte),
		registerClient:   make(chan *Client),
		unregisterClient: make(chan *Client),
		IsTesting:        testing,
	}
}

// Runs the given Server
func (ws *WSServer) Run() {
	for {
		select {
		case newClient := <-ws.registerClient:
			ws.Clients[newClient] = true
		case unregisterClient := <-ws.unregisterClient:
			if _, ok := ws.Clients[unregisterClient]; ok {
				delete(ws.Clients, unregisterClient)
				close(unregisterClient.Send)
			}
		case message := <-ws.Broadcast:
			for client := range ws.Clients {
				select {
				case client.Send <- message:
				// if the client has an error we close the chan and delete it from the map
				default:
					close(client.Send)
					delete(ws.Clients, client)
				}
			}
		}
	}
}

// Settings for upgrading a connection
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// Handels upgrading a connection
func (ws *WSServer) HandlePlayerJoin(w http.ResponseWriter, r *http.Request) {

	newConn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		logging.Log(logging.Error, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// TODO: get the client id!!

	client := Client{
		UserId: 0,
		Conn:   newConn,
		Send:   make(chan []byte),
	}

	ws.registerClient <- &client

	// make a new go routine for the client
	go client.handleConnection(ws)
}

func (client *Client) handleConnection(ws *WSServer) {

}
