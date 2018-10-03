package utils

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
)

// SendResponse : print an error in json format to ResponseWriter
func SendResponse(w http.ResponseWriter, err error, status int) {
	w.WriteHeader(status)
	fmt.Fprintf(w, "{\"Error\": \"%v\"}", err)
}

// RowsToMap : convert a sql row to a map
func RowsToMap(rows *sql.Rows) (map[string]interface{}, error) {
	contentMap := make(map[string]interface{})
	cols, err := rows.Columns()

	if err != nil {
		return contentMap, err
	}
	for rows.Next() {
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}
		err = rows.Scan(columnPointers...)
		if err != nil {
			return contentMap, err
		}
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			contentMap[colName] = *val
		}
	}
	return contentMap, err
}

// BytesToString : convert bytes to string
func BytesToString(bs []uint8) string {
	b := make([]byte, len(bs))

	for i, v := range bs {
		b[i] = byte(v)
	}
	return string(b)
}

// RContToInt : convert an []uint8 to int
func RContToInt(bs []uint8) int {
	ret, err := strconv.Atoi(BytesToString(bs))

	if err != nil {
		return (-1)
	}
	return (ret)
}
