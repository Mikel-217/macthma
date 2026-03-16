package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	matchmastructs "mikel-kunze.com/matchma-client/matchma_structs"
)

func HandleRegistration() {

	if AccesToken != "" {
		HandleUserInput()
	}

	var newUser matchmastructs.UserStruct

retry:
	fmt.Println("======= User registration ======= \n Please enter your username:")

	if _, err := fmt.Scan(&newUser.UserName); err != nil {
		fmt.Println("Please answer correctly")
		goto retry
	}

	newUser.UserMail = newUser.UserName + "@boobles.cloud"

	fmt.Println("Please enter your password:")

	if _, err := fmt.Scan(&newUser.UserPW); err != nil {
		fmt.Println("Please answer correctly")
		goto retry
	}

	json, err := json.Marshal(&newUser)

	if err != nil {
		fmt.Println("failed to create json", err)
		HandleUserInput()
	}

	response, err := http.Post(UrlRegister, "application/json", bytes.NewBuffer(json))

	if err != nil || response.StatusCode >= 400 {
		fmt.Println("Failed request", err)
		HandleUserInput()
	}

	fmt.Println("Success with registration")
	HandleUserInput()
}
