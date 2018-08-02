package main

import (
	"fmt"
	"net/http"
)

func status(w http.ResponseWriter, r *http.Request) {
	sql_err := dbSQL.Ping()

	if sql_err == nil {
		fmt.Fprintln(w, "SQL Database connected")
	} else if sql_err != nil && cfg.SQL == 1 {
		fmt.Fprintln(w, sql_err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome !")
}
