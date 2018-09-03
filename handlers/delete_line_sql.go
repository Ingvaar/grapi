package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/ingvaar/grapi/db"
)

func deleteLineSQL(w http.ResponseWriter, r *http.Request) {
	pathVars := mux.Vars(r)
	tabName := pathVars["table"]
	id := pathVars["id"]
	idNum, errAtoi := strconv.Atoi(id)

	statement := fmt.Sprintf("DELETE FROM %s WHERE id=%d", tabName, idNum)
	if errAtoi != nil {
		fmt.Fprintf(w, "Error: invalid id '%s'\n", id)
	} else {
		_, err := db.Db.SQL.Query(statement)
		if err != nil {
			fmt.Fprintf(w, "%s", err)
		} else {
			fmt.Fprintf(w, "Line %s deleted", id)
		}
	}
}
