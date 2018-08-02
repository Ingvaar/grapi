package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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
		_, err := dbSQL.Query(statement)
		if err != nil {
			fmt.Fprintf(w, "%s", err)
		} else {
			fmt.Fprintf(w, "Line %s deleted", id)
		}
	}
}
