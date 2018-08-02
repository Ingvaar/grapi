package main

import (
	"database/sql"
	"log"
	"net/http"
)

var dbSQL *sql.DB
var cfg Config

func main() {
	opt := ParsCmdline()
	cfg = GetConfig(opt)
	dbSQL = openSQLDatabase(cfg)
	router := NewRouter(opt)

	defer dbSQL.Close()
	log.Fatal(http.ListenAndServe(":8080", router))
}
