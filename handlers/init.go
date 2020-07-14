package handlers

import "database/sql"

// Db - Database instance used throughout the app
var Db *sql.DB

type userInfo struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
}
