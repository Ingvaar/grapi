package main

import (
	"net/http"
	"log"
	"database/sql"
)

var dbSQL *sql.DB = nil
var cfg Config

func main() {
	opt := Pars_cmdline();
	cfg = Get_config(opt)
	dbSQL = openSQLDatabase(cfg);
	router := NewRouter(opt);

	defer dbSQL.Close()
	log.Fatal(http.ListenAndServe(":8080", router));
}
