package sql

import (
	"net/http"

	"github.com/gorilla/mux"

	"grapi/utils"
)

// Show : describe a table
func (db *SQL) Show(w http.ResponseWriter, r *http.Request) {
	tabName := mux.Vars(r)["table"]
	statement := "DESCRIBE " + tabName

	rows, err := db.DB.Query(statement)
	if err != nil {
		utils.SendError(w, err, http.StatusBadRequest)
	} else {
		defer rows.Close()
		colNames, errCol := rows.Columns()
		if errCol == nil {
			printMult(colNames, rows, w)
		}
	}
}
