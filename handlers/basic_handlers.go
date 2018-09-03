package handlers

import (
	"fmt"
	"net/http"

	"github.com/ingvaar/grapi/config"
	"github.com/ingvaar/grapi/db"
)

func status(w http.ResponseWriter, r *http.Request) {
	sqlErr := db.Db.SQL.Ping()

	if sqlErr == nil {
		fmt.Fprintln(w, "SQL Database connected")
	} else if sqlErr != nil && config.Cfg.UseSQL == 1 {
		fmt.Fprintln(w, sqlErr)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome !")
}
