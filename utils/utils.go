package utils

import (
	"fmt"
	"net/http"
	"strconv"
)

// SendError : print an error in json format to ResponseWriter
func SendError(w http.ResponseWriter, err error, status int) {
	w.WriteHeader(status)
	fmt.Fprintf(w, "{\"Error\": \"%v\"}", err)
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
