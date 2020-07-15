package main

import (
	"database/sql"
	"net/http"

	"github.com/anand-kay/factly-crud/handlers"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	handlers.Db = initialiseDb()

	defer handlers.Db.Close()

	r := mux.NewRouter()

	regHandlers(r)

	http.ListenAndServe("localhost:3000", r)
}

func regHandlers(r *mux.Router) {
	r.HandleFunc("/create", handlers.CreateHandler).Methods(http.MethodPost)
	r.HandleFunc("/list", handlers.ListHandler).Methods(http.MethodGet)
	r.HandleFunc("/update/{id}", handlers.UpdateHandler).Methods(http.MethodPatch)
	r.HandleFunc("/delete/{id}", handlers.DeleteHandler).Methods(http.MethodDelete)
}

func initialiseDb() *sql.DB {
	var datname string

	connStr := `host=localhost
		port=` + handlers.DbPort +
		` user=` + handlers.DbUsername +
		` password=` + handlers.DbPassword +
		` sslmode=disable`
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}

	row := db.QueryRow("SELECT datname FROM pg_catalog.pg_database WHERE datname='factlycrud'")
	row.Scan(&datname)

	if datname == "" {
		db.Exec("CREATE DATABASE factlycrud;")
	}

	db.Close()

	connStr1 := `host=localhost
		port=` + handlers.DbPort +
		` user=` + handlers.DbUsername +
		` password=` + handlers.DbPassword +
		` dbname=factlycrud
		sslmode=disable`
	db1, err1 := sql.Open("postgres", connStr1)
	if err1 != nil {
		panic(err1.Error())
	}
	// defer db1.Close()

	db1.Exec(`
		CREATE TABLE IF NOT EXISTS users(
			id SERIAL PRIMARY KEY,
			username VARCHAR(30) UNIQUE NOT NULL,
			email VARCHAR(30) UNIQUE NOT NULL
		);`)

	return db1
}
