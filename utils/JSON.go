package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func WriteJSONFile(filename string, data interface{}) error {
	sessionJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		SendJSONResponse(500, err.Error(), nil)
		return nil
	}

	err = os.WriteFile(filename, sessionJSON, 0644)
	if err != nil {
		SendJSONResponse(500, err.Error(), nil)
		return nil
	}

	return nil
}

func ReadSession() (map[string]interface{}, error) {
	sessionData, err := os.ReadFile("session.json")
	if err != nil {
		return nil, fmt.Errorf("failed to read session: %w", err)
	}

	var session map[string]interface{}
	err = json.Unmarshal(sessionData, &session)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal session data: %w", err)
	}
	return session, nil
}

func ReadBodyJSON(bodyFile string, data interface{}) error {
	file, err := os.Open(bodyFile)
	if err != nil {
		return fmt.Errorf("failed to read body data: %w", err)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(data); err != nil {
		return fmt.Errorf("failed to decode body data: %w", err)
	}
	return nil
}

func ReadLoggedIn(sessionFile string) bool {
	if _, err := os.Stat(sessionFile); err == nil {
		SendJSONResponse(403, "user already logged in", nil)
		return true
	}
	return false
}
