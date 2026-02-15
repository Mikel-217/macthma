package user

import (
	"encoding/base64"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"mikel-kunze.com/matchma/database"
	dbuser "mikel-kunze.com/matchma/database/db_user"
	"mikel-kunze.com/matchma/logging"
	matchmastructs "mikel-kunze.com/matchma/matchma_structs"
)

type claims struct {
	UserId uint
	jwt.RegisteredClaims
}

// Handels the login for user and sends a new accesstoken
func HandleUserLogin(w http.ResponseWriter, r *http.Request) {

	authenticationHeader := r.Header.Get("Authorization")

	if authenticationHeader == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	auth := strings.ReplaceAll(authenticationHeader, "Basic ", "")

	encoded, err := base64.StdEncoding.DecodeString(auth)

	if err != nil {
		logging.Log(logging.Error, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// its stored like this: username:PW
	userCredentials := strings.Split(string(encoded), ":")

	// has the pw
	hashedPW, err := bcrypt.GenerateFromPassword([]byte(userCredentials[1]), 14)

	if err != nil {
		logging.Log(logging.Error, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := dbuser.GetUserByName(userCredentials[0])

	// checks if the user has the right credentials
	if user != nil && user.UserName != userCredentials[0] && user.UserPW != string(hashedPW) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	} else if user == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// generates a token
	token := generateToken(user)

	if token == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(*token))
}

// Generates a JWT
func generateToken(user *matchmastructs.UserStruct) *string {

	claims := claims{
		UserId: user.UserId,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, 3)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT-Secret")))

	if err != nil {
		logging.Log(logging.Error, err.Error())
		return nil
	}

	query := "INSERT INTO AccessTokens VALUES(DEFAULT, ?, ?);"
	queryArgs := []string{tokenString, time.Now().AddDate(0, 0, 3).Format("2006-01-02 15:04:05")}

	result := database.ExecuteSQL(query, queryArgs)

	if result.ErrorMsg != nil {
		logging.Log(logging.Error, result.ErrorMsg.Error())
		return nil
	}

	return &tokenString
}
