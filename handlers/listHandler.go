package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// ListHandler - Handles list route
func ListHandler(w http.ResponseWriter, req *http.Request) {
	if len(req.URL.Query()) == 0 {
		userInfos, err := getAllUsers()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Database error"))

			return
		}

		pages := paginate(userInfos)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(pages)
	} else {
		pageNo, err := strconv.Atoi(req.URL.Query().Get("page"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid page number"))

			return
		}

		userInfos, err := getUsers(pageNo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Database error"))

			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(userInfos)
	}
}

func getUsers(pageNo int) ([]userInfo, error) {
	var userInfos []userInfo
	var userInfo userInfo

	rows, err := Db.Query("SELECT * FROM users OFFSET $1 LIMIT 10", (pageNo * 10))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&userInfo.ID, &userInfo.UserName, &userInfo.Email); err != nil {
			return nil, err
		}

		userInfos = append(userInfos, userInfo)
	}

	return userInfos, nil
}

func getAllUsers() ([]userInfo, error) {
	var userInfos []userInfo
	var userInfo userInfo

	rows, err := Db.Query("SELECT * FROM users;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&userInfo.ID, &userInfo.UserName, &userInfo.Email); err != nil {
			return nil, err
		}

		userInfos = append(userInfos, userInfo)
	}

	return userInfos, nil
}

func paginate(userInfos []userInfo) map[int][]userInfo {
	pages := make(map[int][]userInfo)

	userInfosLen := len(userInfos)
	var noOfPages int

	if userInfosLen%10 == 0 {
		noOfPages = userInfosLen / 10
	} else {
		noOfPages = (userInfosLen / 10) + 1
	}

	for i := 0; i < noOfPages; i++ {
		if i == noOfPages-1 && userInfosLen%10 != 0 {
			pages[i+1] = userInfos[(10 * i):((10 * i) + userInfosLen%10)]
		} else {
			pages[i+1] = userInfos[(10 * i):((10 * i) + 10)]
		}
	}

	return pages
}
