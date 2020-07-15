package handlers

import (
	"database/sql"
	"net/http"
	"testing"
)

var db *sql.DB
var err error

func init() {
	connStr := `host=localhost
		port=` + DbPort +
		` user=` + DbUsername +
		` password=` + DbPassword +
		` sslmode=disable`
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}

	db.Exec("DROP DATABASE IF EXISTS factlycrudtest;")
	db.Exec("CREATE DATABASE factlycrudtest;")

	db.Close()

	connStr1 := `host=localhost
		port=` + DbPort +
		` user=` + DbUsername +
		` password=` + DbPassword +
		` dbname=factlycrudtest
		sslmode=disable`
	db, err = sql.Open("postgres", connStr1)
	if err != nil {
		panic(err.Error())
	}

	db.Exec(`
		CREATE TABLE users(
			id SERIAL PRIMARY KEY,
			username VARCHAR(30) UNIQUE NOT NULL,
			email VARCHAR(30) UNIQUE NOT NULL
		);`)

	Db = db
}

func TestDatabase(t *testing.T) {
	defer Db.Close()

	tInsertDb(t)

	tReadDb(t)

	tUpdateDb(t)

	tDeleteDb(t)
}

func tInsertDb(t *testing.T) {
	var userInfo userInfo

	userInfo.UserName = "rohit"
	userInfo.Email = "rohit@gmail.com"
	statusCode, err := userInfo.insertToDb()
	if statusCode != http.StatusCreated {
		t.Error(err.Error())
	}

	userInfo.UserName = "pranay"
	userInfo.Email = "pranay@gmail.com"
	statusCode1, err1 := userInfo.insertToDb()
	if statusCode1 != http.StatusCreated {
		t.Error(err1.Error())
	}

	userInfo.UserName = "amit"
	userInfo.Email = "amit@gmail.com"
	statusCode2, err2 := userInfo.insertToDb()
	if statusCode2 != http.StatusCreated {
		t.Error(err2.Error())
	}

	userInfo.UserName = "rahul"
	userInfo.Email = "rahul@gmail.com"
	statusCode3, err3 := userInfo.insertToDb()
	if statusCode3 != http.StatusCreated {
		t.Error(err3.Error())
	}

	userInfo.UserName = "hrishi"
	userInfo.Email = "hrishi@gmail.com"
	statusCode4, err4 := userInfo.insertToDb()
	if statusCode4 != http.StatusCreated {
		t.Error(err4.Error())
	}

	userInfo.UserName = "shikar"
	userInfo.Email = "shikar@gmail.com"
	statusCode5, err5 := userInfo.insertToDb()
	if statusCode5 != http.StatusCreated {
		t.Error(err5.Error())
	}

	userInfo.UserName = "amit"
	userInfo.Email = "amitdeol@gmail.com"
	statusCode6, err6 := userInfo.insertToDb()
	if statusCode6 != http.StatusConflict {
		t.Error(err6.Error())
	}

	userInfo.UserName = "rakshan"
	userInfo.Email = "rakshan@gmail.com"
	statusCode7, err7 := userInfo.insertToDb()
	if statusCode7 != http.StatusCreated {
		t.Error(err7.Error())
	}

	userInfo.UserName = "karthik"
	userInfo.Email = "karthik@gmail.com"
	statusCode8, err8 := userInfo.insertToDb()
	if statusCode8 != http.StatusCreated {
		t.Error(err8.Error())
	}

	userInfo.UserName = "sonu"
	userInfo.Email = "sonu@gmail.com"
	statusCode9, err9 := userInfo.insertToDb()
	if statusCode9 != http.StatusCreated {
		t.Error(err9.Error())
	}

	userInfo.UserName = "suresh"
	userInfo.Email = "suresh@gmail.com"
	statusCode10, err10 := userInfo.insertToDb()
	if statusCode10 != http.StatusCreated {
		t.Error(err10.Error())
	}

	userInfo.UserName = "vijay"
	userInfo.Email = "vijay@gmail.com"
	statusCode11, err11 := userInfo.insertToDb()
	if statusCode11 != http.StatusCreated {
		t.Error(err11.Error())
	}

	userInfo.UserName = "sanjay"
	userInfo.Email = "sonu@gmail.com"
	statusCode12, err12 := userInfo.insertToDb()
	if statusCode12 != http.StatusConflict {
		t.Error(err12.Error())
	}

	userInfo.UserName = "vishal"
	userInfo.Email = "vishal@gmail.com"
	statusCode13, err13 := userInfo.insertToDb()
	if statusCode13 != http.StatusCreated {
		t.Error(err13.Error())
	}
}

func tReadDb(t *testing.T) {
	userInfos, err := getAllUsers()
	if err != nil {
		t.Error(err.Error())
	}
	if len(userInfos) != 12 {
		t.Error("Error while fetching all users")
	}

	userInfos1, err1 := getUsers(0)
	if err1 != nil {
		t.Error(err1.Error())
	}
	if len(userInfos1) != 10 {
		t.Error("Error while fetching users by page number")
	}

	userInfos2, err2 := getUsers(1)
	if err2 != nil {
		t.Error(err2.Error())
	}
	if len(userInfos2) != 2 {
		t.Error("Error while fetching users by page number")
	}

	userInfos3, err3 := getUsers(2)
	if err3 != nil {
		t.Error(err3.Error())
	}
	if len(userInfos3) != 0 {
		t.Error("Error while fetching users by page number")
	}

	userInfos4, err4 := getUsers(35)
	if err4 != nil {
		t.Error(err4.Error())
	}
	if len(userInfos4) != 0 {
		t.Error("Error while fetching users by page number")
	}
}

func tUpdateDb(t *testing.T) {
	var userInfo, userInfo1 userInfo

	userInfo.UserName = "hari"
	statusCode, err := userInfo.updateInDb("username", "5")
	if statusCode != http.StatusOK {
		t.Error(err.Error())
	}

	userInfo1.Email = "ajay@gmail.com"
	statusCode1, err1 := userInfo1.updateInDb("email", "2")
	if statusCode1 != http.StatusOK {
		t.Error(err1.Error())
	}

	userInfo2, err2 := getUser("9")
	if err2 != nil {
		t.Error(err2.Error())
	}
	if userInfo2.UserName != "karthik" {
		t.Error("Error while fetching user info")
	}
	if userInfo2.Email != "karthik@gmail.com" {
		t.Error("Error while fetching user info")
	}

	_, err3 := getUser("26")
	if err3 == nil {
		t.Error("Error while fetching non-existing user")
	}
}

func tDeleteDb(t *testing.T) {
	statusCode, err := deleteInDb("8")
	if statusCode != http.StatusOK {
		t.Error(err.Error())
	}

	statusCode1, _ := deleteInDb("57")
	if statusCode1 != http.StatusNotFound {
		t.Error("Error while deleting non-existing user")
	}
}
