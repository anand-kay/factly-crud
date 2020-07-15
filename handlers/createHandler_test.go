package handlers

import (
	"testing"
)

func TestCheckData(t *testing.T) {
	var userInfo userInfo

	userInfo.UserName = "ferguson"
	userInfo.Email = "ferg@gmail.com"
	isValid1, _ := userInfo.checkData()
	if !isValid1 {
		t.Error(`"username":"ferguson", "email":"ferg@gmail.com" - INVALID`)
	}

	userInfo.UserName = "45ferguson"
	userInfo.Email = "ferg@gmail.com"
	isValid2, _ := userInfo.checkData()
	if isValid2 {
		t.Error(`"username":"45ferguson", "email":"ferg@gmail.com" - VALID`)
	}

	userInfo.UserName = "ferguson"
	userInfo.Email = "ferggmail.com"
	isValid3, _ := userInfo.checkData()
	if isValid3 {
		t.Error(`"username":"ferguson", "email":"ferggmail.com" - VALID`)
	}
}

func TestEscapeHTML(t *testing.T) {
	var userInfo userInfo

	userInfo.UserName = "<script>jordy"
	userInfo.Email = "jordy@gmail.com"
	userInfo.escapeHTML()
	if userInfo.UserName != "&lt;script&gt;jordy" || userInfo.Email != "jordy@gmail.com" {
		t.Error(`"username":"<script>jordy", "email":"jordy@gmail.com" - Error escaping HTML`)
	}

	userInfo.UserName = "jordy"
	userInfo.Email = "<script>jordy@gmail.com"
	userInfo.escapeHTML()
	if userInfo.UserName != "jordy" || userInfo.Email != "&lt;script&gt;jordy@gmail.com" {
		t.Error(`"username":"jordy", "email":"<script>jordy@gmail.com" - Error escaping HTML`)
	}
}
