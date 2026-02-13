package main

import (
	"fmt"
	"net/http"
	"os"

	"mikel-kunze.com/matchma/logging"
	"mikel-kunze.com/matchma/matchmaking"
	"mikel-kunze.com/matchma/user"
)

// Middleware
func handleAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RemoteAddr)
		logging.Log(logging.Information, r.RemoteAddr)

		// TODO: check for authentication

		next.ServeHTTP(w, r)
	})

}

func main() {

	if os.Args[1] == "--testing" {
		fmt.Println("Testing enabled")
		// TODO implement startup Service which creates 100 Players
		// Create a channel where it returns done or not
	}

	mux := http.NewServeMux()

	fmt.Println("Server listening on http://localhost:8080/")

	mux.HandleFunc("POST /login", user.HandleUserLogin)
	mux.HandleFunc("POST /register", user.HandleUserRegister)

	mux.Handle("/join-match", handleAuthentication(http.HandlerFunc(matchmaking.HandlePlayerJoin)))

	http.ListenAndServe(":8080", mux)

}
