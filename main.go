package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServerMux()

	fmt.Println("Server listening on http://localhost:8080/")

	http.ListenAndServe(":8080", mux)
}
