package handlers

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
)

func HandleLogin() string {

	var userName, pw string

	fmt.Println("Please enter your username:")

	if _, err := fmt.Scan(&userName); err != nil {
		fmt.Println("Please enter valid credentials")
		HandleLogin()
	}

	fmt.Println("Please enter your password")

	if _, err := fmt.Scan(&pw); err != nil {
		fmt.Println("Please enter valid credentials")
		HandleLogin()
	}

	basicAuth := base64.StdEncoding.EncodeToString([]byte(userName + ":" + pw))

	// build the request
	req, _ := http.NewRequest("POST", UrlLogin, nil)
	headerval := "Basic " + basicAuth
	req.Header.Add("Authorization", headerval)

	// send the request
	response, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err.Error())
		HandleUserInput()
	}

	if response.StatusCode != 200 {
		fmt.Println("Failed to login status:", response.StatusCode, response.Status)
	}

	bodyData, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err.Error())
		HandleUserInput()
	}

	return string(bodyData)
}
