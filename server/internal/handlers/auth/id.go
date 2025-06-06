package auth

import (
	"database/sql"
	"errors"
	"net/http"

	"execute/internal"
)

// getUserID reads the session cookie, validates it, looks up the username
// then looks up that user’s ID in the database
func GetUserID(r *http.Request) (int, error) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			return 0, errors.New("no session token found")
		}
		return 0, err
	}
	token := cookie.Value

	username, exists := GetSessionUsername(token)
	if !exists {
		return 0, errors.New("invalid or expired session token")
	}

	var userID int
	err = internal.DB.QueryRow(
		"SELECT id FROM users WHERE username = $1",
		username,
	).Scan(&userID)
	if err == sql.ErrNoRows {
		return 0, errors.New("user not found")
	}
	return userID, err
}
