package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type colStruct struct {
	colPtr     []interface{}
	colCount   int
	colNames   []string
	rowContent map[string]string
}

func getTableSQL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tabName := vars["table"]
	statement := fmt.Sprintf("SELECT * FROM %s", tabName)

	rows, err := dbSQL.Query(statement)
	defer rows.Close()
	colNames, errCol := rows.Columns()
	if err != nil || errCol != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", err)
	} else {
		PrintRows(colNames, rows, w)
	}
}

// PrintRows : print multiple rows from a table
func PrintRows(colNames []string, rows *sql.Rows,
	w http.ResponseWriter) {
	colsMap := CreateColsMap(colNames)
	fmt.Fprintf(w, "[")
	multRows := false
	for rows.Next() {
		if multRows {
			fmt.Fprintf(w, ",")
		}
		colsMap.UpdateColMap(rows)
		cols := colsMap.GetColsFromMap()
		jsonStr, jsonErr := json.Marshal(cols)
		if jsonErr == nil {
			fmt.Fprintf(w, "%s", jsonStr)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, jsonErr)
		}
		multRows = true
	}
	fmt.Fprintf(w, "]")
}

// CreateColsMap : creates maps of the columns
func CreateColsMap(columns []string) *colStruct {
	colLen := len(columns)
	colStruct := &colStruct{
		colPtr:     make([]interface{}, colLen),
		colCount:   colLen,
		colNames:   columns,
		rowContent: make(map[string]string, colLen),
	}

	for i := 0; i < colLen; i++ {
		colStruct.colPtr[i] = new(sql.RawBytes)
	}
	return (colStruct)
}

func (colStruct *colStruct) UpdateColMap(rows *sql.Rows) error {
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
			errConv := fmt.Errorf("Cannot convert index %d column %s",
				i, colStruct.colNames[i])
			return (errConv)
		}
	}
	return (nil)
}

func (colStruct *colStruct) GetColsFromMap() map[string]string {
	return (colStruct.rowContent)
}
