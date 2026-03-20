package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	dbtoken "mikel-kunze.com/matchma/database/db_token"
	"mikel-kunze.com/matchma/logging"
)

// Middleware to authenticate clients
func handleAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logging.Log(logging.Information, r.RemoteAddr)

		token := r.Header.Get("Authorization")

		if token == "" {
			fmt.Println("Token nil")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		var tokenWithoutBaerer string

		if strings.Contains(token, "bearer") {
			tokenWithoutBaerer = strings.ReplaceAll(token, "bearer ", "")
		} else {
			tokenWithoutBaerer = strings.ReplaceAll(token, "Bearer ", "")
		}

		if !authenticate(tokenWithoutBaerer) {
			fmt.Println("cannot authenticate")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// checks if an client is authenticated
func authenticate(token string) bool {

	claims := jwt.MapClaims{}

	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT-Secret")), nil
	})

	if err != nil {
		logging.Log(logging.Error, err.Error())
		return false
	}

	if !parsedToken.Valid && !dbtoken.IsTokenThere(token) {
		return false
	}

	return true
}
