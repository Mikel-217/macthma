package matchmaking

import (
	"net/http"
	"strconv"

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

			if ws.IsTesting || len(ws.Clients) == 20 {
				ws.CreateLobbys()
			}

		case unregisterClient := <-ws.unregisterClient:
			if _, ok := ws.Clients[unregisterClient]; ok {
				delete(ws.Clients, unregisterClient)
				close(unregisterClient.Send)
			}
		// FIXME: why no messages?
		case message := <-ws.Broadcast:
			logging.Log(logging.Information, "Trying to broadcast..")
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

	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	urlQuery := r.URL.Query()
	userIdstring := urlQuery.Get("user")

	if userIdstring == "" {
		logging.Log(logging.Error, "UserId nil")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userId, err := strconv.Atoi(userIdstring)

	if err != nil {
		logging.Log(logging.Error, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newConn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		logging.Log(logging.Error, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	client := Client{
		UserId: uint(userId),
		Conn:   newConn,
		Send:   make(chan []byte),
	}

	ws.registerClient <- &client

	// make a new go routine for the client
	go client.handleConnection(ws)
}

// Handels the connection from the client
func (client *Client) handleConnection(ws *WSServer) {
	defer func() {
		ws.unregisterClient <- client
		client.Conn.Close()
	}()

	for {
		message, ok := <-client.Send
		if !ok {
			client.Conn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}

		err := client.Conn.WriteMessage(websocket.TextMessage, message)

		if err != nil {
			logging.Log(logging.Error, "Write error: "+err.Error())
			return
		}
	}
}
