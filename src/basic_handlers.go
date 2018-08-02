package main

import (
	"fmt"
	"net/http"
)

func status(w http.ResponseWriter, r *http.Request) {
	sqlErr := dbSQL.Ping()

	if sqlErr == nil {
		fmt.Fprintln(w, "SQL Database connected")
	} else if sqlErr != nil && cfg.SQL == 1 {
		fmt.Fprintln(w, sqlErr)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome !")
}
