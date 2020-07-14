package main

import (
	"database/sql"
	"net/http"

	"github.com/anand-kay/factly-crud/handlers"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	connStr := `host=localhost 
				port=5433 
				user=anand 
				password=mypostgres 
				dbname=factlycrud 
				sslmode=disable`
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	handlers.Db = db

	r := mux.NewRouter()

	regHandlers(r)

	http.ListenAndServe("localhost:3000", r)
}

func regHandlers(r *mux.Router) {
	r.HandleFunc("/create", handlers.CreateHandler).Methods(http.MethodPost)
	r.HandleFunc("/list", handlers.ListHandler).Methods(http.MethodGet)
}
