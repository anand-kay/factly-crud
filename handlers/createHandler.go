package handlers

import (
	"encoding/json"
	"errors"
	"html"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/lib/pq"
)

// CreateHandler - Handles create route
func CreateHandler(w http.ResponseWriter, req *http.Request) {
	var userInfo userInfo

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("Error reading request body"))

		return
	}

	json.Unmarshal(reqBody, &userInfo)

	userInfo.escapeHTML()

	isValid, errMsg := userInfo.checkData()
	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errMsg))

		return
	}

	statusCode, errDb := userInfo.insertToDb()
	if errDb != nil {
		w.WriteHeader(statusCode)
		w.Write([]byte(errDb.Error()))

		return
	}

	w.WriteHeader(statusCode)
	w.Write([]byte("User created successfully"))
}

func (userInfo *userInfo) insertToDb() (int, error) {
	_, err := Db.Query("INSERT INTO users(username,email) VALUES ($1,$2)", userInfo.UserName, userInfo.Email)
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

	return http.StatusCreated, nil
}

func (userInfo *userInfo) checkData() (bool, string) {
	if !checkUsername(userInfo.UserName) {
		return false, "Invalid username"
	}

	if !checkEmail(userInfo.Email) {
		return false, "Invalid email address"
	}

	return true, ""
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
