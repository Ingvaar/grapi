package sql

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"grapi/utils"
)

// Insert : insert value in table
func (db *Database) Insert(w http.ResponseWriter, r *http.Request) {
	multInsert := false
	tabName := mux.Vars(r)["table"]
	r.ParseForm()

	fmt.Printf("DEBUG : %v\n", r.Form)
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
	_, err := db.DB.Query(statement)
	if err != nil {
		utils.SendError(w, err, http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}
