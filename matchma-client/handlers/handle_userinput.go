package handlers

import (
	"fmt"
)

var AccesToken string

// handles the user input -> gets called from all other handlers
func HandleUserInput() {
	for {
	retry:
		fmt.Println("====== Select ====== \n [1] Register \n [2] Login \n [3] Connect to Lobby")

		var response int
		if _, err := fmt.Scan(&response); err != nil {
			fmt.Println("Please answer correctly")
			goto retry
		}

		// selects what the user wants
		switch response {
		case 1:
			HandleRegistration()
		case 2:
			AccesToken = HandleLogin()
		case 3:
			HandleRegistration()
		default:
			fmt.Println("Please answer correctly")
			goto retry
		}
	}

}
