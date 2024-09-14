package main

import (
	"database/sql"
	"log"
	"pafaul/ritter/client-server"

	_ "github.com/mattn/go-sqlite3"
)

// @title			Swagger Example API
// @version			1.0
// @license.name	Apache 2.0
// @host			localhost:8000
// @BasePath		/api/v1
func main() {
	db, err := sql.Open("sqlite3", "file:db.sqlite")
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	s := client_server.NewServer("localhost:8000", db)
	err = s.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
