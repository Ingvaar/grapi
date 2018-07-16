package main

import (
	"net/http"
	"log"
	"database/sql"
)

var db *sql.DB = nil

func main() {
	opt := Pars_cmdline();
	db = openDatabase(opt);
	router := NewRouter(opt);

	log.Fatal(http.ListenAndServe(":8080", router));
}
