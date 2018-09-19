package sql

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"grapi/utils"
	"net/http"
	"strconv"
	"strings"
)

type colStruct struct {
	colPtr	   []interface{}
	colCount   int
	colNames   []string
	rowContent map[string]string
}

// PrintOne : function to print one row from a table
func PrintOne(colNames []string, rows *sql.Rows,
	w http.ResponseWriter) {
	rows.Next()
	colsMap := createMap(colNames)
	colsMap.updateMap(rows)
	cols := colsMap.fillFromMap()
	jsonStr, jsonErr := json.Marshal(cols)
	if jsonErr == nil {
		fmt.Fprintf(w, "%s", jsonStr)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ErrorToJSON(w, jsonErr)
	}
}

// PrintMult : print multiple rows from a table
func printMult(colNames []string, rows *sql.Rows,
	w http.ResponseWriter) {
	colsMap := createMap(colNames)
	fmt.Fprintf(w, "[")
	multRows := false
	for rows.Next() {
		if multRows {
			fmt.Fprintf(w, ",")
		}
		colsMap.updateMap(rows)
		cols := colsMap.fillFromMap()
		jsonStr, jsonErr := json.Marshal(cols)
		if jsonErr == nil {
			fmt.Fprintf(w, "%s", jsonStr)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			utils.ErrorToJSON(w, jsonErr)
		}
		multRows = true
	}
	fmt.Fprintf(w, "]")
}

// CreateMap : creates maps of the columns
func createMap(columns []string) *colStruct {
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

func (colStruct *colStruct) updateMap(rows *sql.Rows) error {
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

func (colStruct *colStruct) fillFromMap() map[string]string {
	return (colStruct.rowContent)
}

// ProcessStr : format the str to match sql request
func ProcessStr(str string) string {
	_, err := strconv.Atoi(str)
	if err != nil {
		str = strings.Replace(str, `\`, `\\`, -1)
		str = strings.Replace(str, `'`, `\'`, -1)
		str = strings.Replace(str, `"`, `\"`, -1)
		str = `"` + str + `"`
	}
	return (str)
}
