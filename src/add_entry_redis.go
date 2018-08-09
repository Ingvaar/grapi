package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// addEntryRedis : use the hmset cmd from redis with a json passed as body
// and with {type}:{id} as the id of the entry
func addEntryRedis(w http.ResponseWriter, r *http.Request) {
	jsonmap := jsonToMap(w, r)
	pathVars := mux.Vars(r)
	id := fmt.Sprintf("%s:%s", pathVars["type"], pathVars["id"])
	var errStr string

	for key, value := range jsonmap {
		err := redisCli.Cmd("HMSET", id, key, value).Err
		errStr = fmt.Sprintf("%s\n%s", errStr, err)
	}
	if errStr != "" {
		fmt.Fprintf(w, "%s\n", errStr)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}
