package sql

import (
	"net/http"

	"github.com/gorilla/mux"

	"grapi/db"
	"grapi/utils"
)

// Show :
func Show(w http.ResponseWriter, r *http.Request) {
	tabName := mux.Vars(r)["table"]
	statement := "DESCRIBE " + tabName

	rows, err := db.SQL.Query(statement)
	if err != nil {
		utils.SendResponse(w, err, http.StatusBadRequest)
	} else {
		defer rows.Close()
		colNames, errCol := rows.Columns()
		if errCol == nil {
			printMult(colNames, rows, w)
		}
	}
}
