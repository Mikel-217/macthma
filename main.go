package main

import (
	"fmt"
	"net/http"

	"mikel-kunze.com/matchma/user"
)

func main() {
	mux := http.NewServerMux()

	fmt.Println("Server listening on http://localhost:8080/")

	mux.HandleFunc("/login", user.UserLoginHandler)

	http.ListenAndServe(":8080", mux)
}
