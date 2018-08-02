package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func updateLineSQL(w http.ResponseWriter, r *http.Request) {
	jsonmap := json_to_map(w, r)
	pathVars := mux.Vars(r)
	tab_name := pathVars["table"]
	id := pathVars["id"]
	mult_insert := false

	statement := fmt.Sprintf("UPDATE %s SET ", tab_name)
	for key, value := range jsonmap {
		if mult_insert {
			statement = fmt.Sprintf("%s, ", statement)
		}
		statement = fmt.Sprintf("%s%s = %s", statement, key, value)
		mult_insert = true
	}
	statement = fmt.Sprintf("%s WHERE id=%s;", statement, id)
	_, err := dbSQL.Exec(statement)
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
