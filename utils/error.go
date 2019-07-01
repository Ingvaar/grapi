package utils

import (
	"fmt"
	"net/http"
)

// SendError : print an error in json format to ResponseWriter
func SendError(w http.ResponseWriter, err error, status int) {
	w.WriteHeader(status)
	fmt.Fprintf(w, "{\"Error\": \"%v\"}", err)
}
