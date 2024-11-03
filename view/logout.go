package handler

import (
	"os"
	"restaurant-app/utils"
)

func Logout() {

	sessionFile := "session.json"

	if _, err := os.Stat(sessionFile); os.IsNotExist(err) {
		utils.SendJSONResponse(401, "You are logged out or not logged in.", nil)
		return
	}

	err := os.Remove(sessionFile)
	if err != nil {
		utils.SendJSONResponse(400, err.Error(), nil)
		return
	}
	

	utils.SendJSONResponse(200, "logout success", nil)
}
