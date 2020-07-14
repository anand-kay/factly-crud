package handlers

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

// DeleteHandler - Handles delete route
func DeleteHandler(w http.ResponseWriter, req *http.Request) {
	statusCode, err := deleteInDb(mux.Vars(req)["id"])
	if err != nil {
		w.WriteHeader(statusCode)
		w.Write([]byte(err.Error()))

		return
	}

	w.WriteHeader(statusCode)
	w.Write([]byte("User deleted successfully"))
}

func deleteInDb(id string) (int, error) {
	result, err := Db.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return http.StatusInternalServerError, errors.New("Error while deleting user")
	}

	noOfRows, _ := result.RowsAffected()

	if noOfRows == 0 {
		return http.StatusNotFound, errors.New("User not found")
	}

	return http.StatusOK, nil
}
