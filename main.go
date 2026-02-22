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

// Middleware
func handleAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logging.Log(logging.Information, r.RemoteAddr)

		token := r.Header.Get("Autorization")

		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if !authenticate(token) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// TODO: add server-console

	if startup.CreateTables() {
		logging.Log(logging.Information, "Setup successfull")
		fmt.Println("Success setting fings up!")
	}

	if os.Args[1] == "--testing" {

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

	mux := http.NewServeMux()

	fmt.Println("Server listening on http://localhost:8080/")

	mux.HandleFunc("POST /login", user.HandleUserLogin)
	mux.HandleFunc("POST /register", user.HandleUserRegister)

	mux.Handle("/join-match", handleAuthentication(http.HandlerFunc(user.HandlePlayerJoin)))
	mux.Handle("POST /match-data", handleAuthentication(http.HandlerFunc(matchmaking.HandleNewMatchData)))

	http.ListenAndServe(":8080", mux)
}
