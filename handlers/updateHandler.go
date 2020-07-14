package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
)

// UpdateHandler - Handles update route
func UpdateHandler(w http.ResponseWriter, req *http.Request) {
	var userInfo userInfo
	var field string

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("Error reading request body"))

		return
	}

	json.Unmarshal(reqBody, &userInfo)

	if userInfo.UserName != "" {
		field = "username"
	} else if userInfo.Email != "" {
		field = "email"
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request"))

		return
	}

	userInfo.escapeHTML()

	if field == "username" {
		if !checkUsername(userInfo.UserName) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid username"))

			return
		}
	} else if field == "email" {
		if !checkEmail(userInfo.Email) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid email address"))

			return
		}
	}

	statusCode, err := userInfo.updateInDb(field, mux.Vars(req)["id"])
	if err != nil {
		w.WriteHeader(statusCode)
		w.Write([]byte(err.Error()))

		return
	}

	userFromDb, err := getUser(mux.Vars(req)["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Details updated but unable fetch user info"))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(userFromDb)
}

func (userInfo *userInfo) updateInDb(field string, id string) (int, error) {
	var err error

	if field == "username" {
		_, err = Db.Exec("UPDATE users SET username=$1 WHERE id=$2;", userInfo.UserName, id)
	} else if field == "email" {
		_, err = Db.Exec("UPDATE users SET email=$1 WHERE id=$2;", userInfo.Email, id)
	}

	if err != nil {
		if errDb, ok := err.(*pq.Error); ok {
			if errDb.Code.Name() == "unique_violation" {
				if strings.Contains(errDb.Message, "users_username_key") {
					return http.StatusConflict, errors.New("Username exists already")
				} else if strings.Contains(errDb.Message, "users_email_key") {
					return http.StatusConflict, errors.New("Email exists already")
				} else {
					return http.StatusInternalServerError, errors.New("Database error")
				}
			} else {
				return http.StatusInternalServerError, errors.New("Database error")
			}
		}

		return http.StatusInternalServerError, errors.New("Database error")
	}

	return http.StatusOK, nil
}

func getUser(id string) (userInfo, error) {
	var userFromDb userInfo

	row := Db.QueryRow("SELECT username,email FROM users WHERE id=$1;", id)
	switch err := row.Scan(&userFromDb.UserName, &userFromDb.Email); err {
	case sql.ErrNoRows:
		return userFromDb, errors.New("User not found")
	case nil:
		return userFromDb, nil
	default:
		return userFromDb, errors.New("Database error")
	}
}
