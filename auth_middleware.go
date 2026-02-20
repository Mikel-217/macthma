package main

import (
	"os"

	"github.com/golang-jwt/jwt/v4"
	dbtoken "mikel-kunze.com/matchma/database/db_token"
	"mikel-kunze.com/matchma/logging"
)

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
