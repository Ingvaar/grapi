package sql

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"grapi/db"
	"grapi/utils"
)

// Insert :
func Insert(w http.ResponseWriter, r *http.Request) {
	multInsert := false
	tabName := mux.Vars(r)["table"]
	r.ParseForm()

	statement := "INSERT INTO " + tabName + " ("
	values := "VALUES ("
	for key, vars := range r.Form {
		if multInsert {
			statement += ", "
			values += ", "
		}
		statement += key
		for _, value := range vars {
			values += ProcessStr(value)
		}
		multInsert = true
	}
	statement = fmt.Sprintf("%s) %s);", statement, values)
	_, err := db.SQL.Exec(statement)
	if err != nil {
		utils.ErrorToJSON(w, err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}
