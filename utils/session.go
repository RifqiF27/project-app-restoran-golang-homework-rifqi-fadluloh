package utils

func SessionAdmin() (int, bool) {
	session, err := ReadSession()
	if err != nil {
		SendJSONResponse(401, "Unauthorized", nil)
		return 0, false
	}
	role, ok := session["Role"].(string)
	userID, okID := session["ID"].(float64)
	if !ok || role != "Admin" || !okID {
		SendJSONResponse(403, "Forbidden", nil)
		return 0, false
	}

	return int(userID), role == "Admin"
}

func Session() (int, bool) {
	session, err := ReadSession()
	if err != nil {
		SendJSONResponse(401, "Unauthorized", nil)
		return 0, false
	}
	userID, okID := session["ID"].(float64)
	if  !okID {
		SendJSONResponse(403, "Forbidden", nil)
		return 0, false
	}

	return int(userID), true
}