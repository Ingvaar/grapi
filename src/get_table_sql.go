package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"database/sql"
	"fmt"
)

type colStruct struct {
	colPtr		[]interface{}
	colCount	int
	colNames	[]string
	rowContent	map[string]string
}

func getTableSQL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r);
	tab_name := vars["table"]
	statement := fmt.Sprintf("SELECT * FROM %s", tab_name)

	rows, err := dbSQL.Query(statement)
	defer rows.Close()
	col_names, err_col := rows.Columns()
	if err != nil || err_col != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", err)
	} else {
		Print_rows(col_names, rows, w)
	}
}

func Print_rows(col_names []string, rows *sql.Rows,
		w http.ResponseWriter) {
	cols_map := Create_cols_map(col_names)
	fmt.Fprintf(w, "[")
	mult_rows := false
	for rows.Next() {
		if mult_rows {
			fmt.Fprintf(w, ",")
		}
		cols_map.Update_col_map(rows)
		cols := cols_map.Get_cols_from_map()
		jsonStr, json_err := json.Marshal(cols)
		if json_err == nil {
			fmt.Fprintf(w, "%s", jsonStr)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, json_err)
		}
		mult_rows = true
	}
	fmt.Fprintf(w, "]")
}

func Create_cols_map(columns []string) *colStruct {
	col_len := len(columns)
	colStruct := &colStruct {
		colPtr:		make([]interface{}, col_len),
		colCount:	col_len,
		colNames:	columns,
		rowContent:	make(map[string]string, col_len),
	}

	for i := 0; i < col_len; i++ {
		colStruct.colPtr[i] = new(sql.RawBytes)
	}
	return (colStruct)
}

func (colStruct *colStruct) Update_col_map(rows *sql.Rows) error {
	err := rows.Scan(colStruct.colPtr...)

	if err != nil {
		return (err)
	}
	for i := 0; i < colStruct.colCount; i++ {
		rb, ok := colStruct.colPtr[i].(*sql.RawBytes)
		if ok {
			colStruct.rowContent[colStruct.colNames[i]] = string(*rb)
			*rb = nil
		} else {
			err_conv := fmt.Errorf("Cannot convert index %d column %s",
						i, colStruct.colNames[i])
			return (err_conv)
		}
	}
	return (nil)
}

func (colStruct *colStruct) Get_cols_from_map() map[string]string {
	return (colStruct.rowContent)
}
