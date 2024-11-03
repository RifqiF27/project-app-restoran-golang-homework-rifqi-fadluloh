package utils

import (
	"encoding/json"
	"fmt"
	"restaurant-app/model"
)

func SendJSONResponse(statusCode int, message string, data interface{}) {
	response := model.Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}

	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}
