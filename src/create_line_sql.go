package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func createLineSQL(w http.ResponseWriter, r *http.Request) {
	jsonmap := json_to_map(w, r)
	pathVars := mux.Vars(r)
	tab_name := pathVars["table"]
	mult_insert := false

	statement := fmt.Sprintf("INSERT INTO %s (", tab_name)
	values := fmt.Sprintf("VALUES (")
	for key, value := range jsonmap {
		if mult_insert {
			statement = fmt.Sprintf("%s, ", statement)
			values = fmt.Sprintf("%s, ", values)
		}
		statement = fmt.Sprintf("%s%s", statement, key)
		values = fmt.Sprintf("%s%s", values, value)
		mult_insert = true
	}
	statement = fmt.Sprintf("%s) %s);", statement, values)
	_, err := dbSQL.Exec(statement)
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}
