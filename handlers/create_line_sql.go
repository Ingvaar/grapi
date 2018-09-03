package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"

	"grapi/db"
)

func createLineSQL(w http.ResponseWriter, r *http.Request) {
	var db db.Database
	jsonmap := jsonToMap(w, r)
	pathVars := mux.Vars(r)
	tabName := pathVars["table"]
	multInsert := false

	statement := fmt.Sprintf("INSERT INTO %s (", tabName)
	values := fmt.Sprintf("VALUES (")
	for key, value := range jsonmap {
		if multInsert {
			statement = fmt.Sprintf("%s, ", statement)
			values = fmt.Sprintf("%s, ", values)
		}
		statement = fmt.Sprintf("%s%s", statement, key)
		values = fmt.Sprintf("%s%s", values, value)
		multInsert = true
	}
	statement = fmt.Sprintf("%s) %s);", statement, values)
	_, err := db.SQL.Exec(statement)
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}
