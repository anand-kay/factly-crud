package handlers

import (
	"database/sql"
	"html"
	"regexp"
)

// Db - Database instance used throughout the app
var Db *sql.DB

type userInfo struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
}

func checkUsername(username string) bool {
	re := regexp.MustCompile("^[A-Za-z]+$")

	if !re.MatchString(username) {
		return false
	}

	return true
}

func checkEmail(email string) bool {
	re := regexp.MustCompile("^[a-z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?(?:\\.[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?)*$")

	if !re.MatchString(email) {
		return false
	}

	return true
}

func (userInfo *userInfo) escapeHTML() {
	userInfo.UserName = html.EscapeString(userInfo.UserName)
	userInfo.Email = html.EscapeString(userInfo.Email)
}
