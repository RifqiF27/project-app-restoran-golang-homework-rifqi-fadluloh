package utils

import (
	"encoding/json"
	"io"
	"os"
)

func DecodeJSONFile(filename string, out interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		SendJSONResponse(500, err.Error(), nil)
		return nil
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(out)
	if err != nil && err != io.EOF {
		SendJSONResponse(500, err.Error(), nil)
		return nil
	}

	return nil
}

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
		// SendJSONResponse(500, err.Error(), nil)
		return nil, nil
	}

	var session map[string]interface{}
	err = json.Unmarshal(sessionData, &session)
	if err != nil {
		SendJSONResponse(500, err.Error(), nil)
		return nil, nil
	}
	return session, nil
}

func ReadBodyJSON(bodyFile string, data interface{}) error {
	file, err := os.Open(bodyFile)
	if err != nil {
		SendJSONResponse(500, err.Error(), nil)
		return err
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(data); err != nil {
		SendJSONResponse(500, err.Error(), nil)
		return err
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
