package main

import (
	"net/http"
)

func main() {
	mux := http.NewServerMux()

	http.ListenAndServe(":8080", mux)
}
