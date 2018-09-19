package utils

import (
	"fmt"
	"net/http"
)

// ErrorToJSON : print an error in json format to ResponseWriter
func ErrorToJSON(w http.ResponseWriter, err error) {
	fmt.Fprintf(w, "{\"Error\": \"%v\"}", err)
}
