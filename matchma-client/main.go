package main

import (
	"fmt"
	"net/http"
	"os"

	"mikel-kunze.com/matchma-client/handlers"
)

func main() {
	fmt.Println("===== Matchma client ======")

	if !preCheck() {
		fmt.Println("Cannot connect to server. Program shut downs...")
		os.Exit(0)
	}
	handlers.HandleUserInput()
}

// just checks if the server is there or not :)
func preCheck() bool {
	if _, err := http.Get(handlers.Url); err != nil {
		return false
	}
	return true
}
