package main

import (
	"strconv"
	"fmt"
	"net/http"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
)

func getLineSQL(w http.ResponseWriter, r *http.Request) {
	pathVars := mux.Vars(r)
	tab_name := pathVars["table"]
	id := pathVars["id"]
	id_num, err_atoi := strconv.Atoi(id)

	statement := fmt.Sprintf("SELECT * FROM %s WHERE id=%d", tab_name, id_num)
	if err_atoi != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: invalid id '%s'\n", id)
	} else {
		rows, err := dbSQL.Query(statement)
		defer rows.Close()
		col_names, err_col := rows.Columns()
		if err != nil || err_col != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "%s", err)
		} else {
			Print_row(col_names, rows, w)
		}
	}
}

func Print_row(col_names []string, rows *sql.Rows,
		w http.ResponseWriter) {
	rows.Next()
	cols_map := Create_cols_map(col_names)
	cols_map.Update_col_map(rows)
	cols := cols_map.Get_cols_from_map()
	jsonStr, json_err := json.Marshal(cols)
	if json_err == nil {
		fmt.Fprintf(w, "%s", jsonStr)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, json_err)
	}
}
