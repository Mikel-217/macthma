package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"mikel-kunze.com/matchma/logging"
	"mikel-kunze.com/matchma/matchmaking"
	"mikel-kunze.com/matchma/startup"
	"mikel-kunze.com/matchma/user"
)

func main() {
	// TODO: add server-console

	if startup.CreateTables() {
		logging.Log(logging.Information, "Setup successfull")
		fmt.Println("Success setting things up!")
	}

	var testingEnabled bool = false

	if os.Args[1] == "--testing" {
		testingEnabled = true
		playerCommand := strings.ReplaceAll(os.Args[2], "--player-count=", "")

		// gets the given player count from the startup command
		playerCount, err := strconv.ParseInt(playerCommand, 0, 0)

		if err != nil {
			logging.Log(logging.Error, err.Error())
			// if an occurs set default val to 200
			playerCount = 200
		}

		go startup.AddTesting(int(playerCount))
	}

	// creates a new websocket server and runs it
	ws := matchmaking.CreateNewWSServer(testingEnabled)
	go ws.Run()

	// makes a handler for upgrading and the connection logic
	wsHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ws.HandlePlayerJoin(w, r)
	})

	mux := http.NewServeMux()

	fmt.Println("Server listening on http://localhost:8080/")

	mux.HandleFunc("POST /login", user.HandleUserLogin)
	mux.HandleFunc("POST /register", user.HandleUserRegister)

	mux.Handle("GET /join-match", handleAuthentication(http.HandlerFunc(wsHandler)))
	mux.Handle("POST /match-data", handleAuthentication(http.HandlerFunc(matchmaking.HandleNewMatchData)))

	http.ListenAndServe(":8080", mux)
}
