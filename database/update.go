package sql

import (
	"net/http"

	"github.com/gorilla/mux"

	"grapi/utils"
)

// Update : update value in row
func (db *SQL) Update(w http.ResponseWriter, r *http.Request) {
	multInsert := false
	r.ParseForm()

	statement := "UPDATE " + mux.Vars(r)["table"] + " SET "
	for key, vars := range r.Form {
		if multInsert {
			statement += ", "
		}
		for _, value := range vars {
			statement += key + "=" + ProcessStr(value)
		}
		multInsert = true
	}
	statement += " WHERE id=" + mux.Vars(r)["id"] + ";"
	_, err := db.DB.Exec(statement)
	if err != nil {
		utils.SendError(w, err, http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
