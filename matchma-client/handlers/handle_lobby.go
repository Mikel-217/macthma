package handlers

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

func HandleLobby() {

	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/join-match"}

	header := make(http.Header)
	header.Add("Authorization", "Baerer "+AccesToken)

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), header)

	if err != nil {
		log.Println(err.Error())
		HandleUserInput()
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	done := make(chan struct{})

	go func() {
		defer conn.Close()
		for {
			_, msg, err := conn.ReadMessage()

			if err != nil {
				log.Fatal("Read-error", err)
				break
			}

			fmt.Println("msg from server: ", string(msg))
		}
	}()

	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("User requested closing")

			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))

			if err != nil {
				log.Fatal(err)
			}
			return
		}
	}
}
