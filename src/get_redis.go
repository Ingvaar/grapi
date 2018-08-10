package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// getRedis : do a hget with a json array passed in the body
// on the id passed in the url an return a json array
func getRedis(w http.ResponseWriter, r *http.Request) {
	jsonmap := jsonToMap(w, r)
	pathVars := mux.Vars(r)
	id := fmt.Sprintf("%s:%s", pathVars["type"], pathVars["id"])
	mult := false

	fmt.Fprintf(w, "{")
	for key := range jsonmap {
		reply, err := redisCli.Cmd("HGET", id, key).Str()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		if mult {
			fmt.Fprintf(w, ", ")
		}
		mult = true
		fmt.Fprintf(w, "\"%s\":\"%s\"", key, reply)
	}
	fmt.Fprintf(w, "}")
}
