package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func getLineSQL(w http.ResponseWriter, r *http.Request) {
	pathVars := mux.Vars(r)
	tabName := pathVars["table"]
	id := pathVars["id"]
	idNum, errAtoi := strconv.Atoi(id)

	statement := fmt.Sprintf("SELECT * FROM %s WHERE id=%d", tabName, idNum)
	if errAtoi != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: invalid id '%s'\n", id)
	} else {
		rows, err := dbSQL.Query(statement)
		defer rows.Close()
		colNames, errCol := rows.Columns()
		if err != nil || errCol != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "%s", err)
		} else {
			printRow(colNames, rows, w)
		}
	}
}

// printRow : function to print one row from a table
func printRow(colNames []string, rows *sql.Rows,
	w http.ResponseWriter) {
	rows.Next()
	colsMap := CreateColsMap(colNames)
	colsMap.UpdateColMap(rows)
	cols := colsMap.GetColsFromMap()
	jsonStr, jsonErr := json.Marshal(cols)
	if jsonErr == nil {
		fmt.Fprintf(w, "%s", jsonStr)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, jsonErr)
	}
}
