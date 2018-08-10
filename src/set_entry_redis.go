package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// setEntryRedis : use the hmset cmd from redis with a json passed as body
// and with {type}:{id} as the id of the entry
func setEntryRedis(w http.ResponseWriter, r *http.Request) {
	jsonmap := jsonToMap(w, r)
	pathVars := mux.Vars(r)
	id := fmt.Sprintf("%s:%s", pathVars["type"], pathVars["id"])

	for key, value := range jsonmap {
		err := redisCli.Cmd("HMSET", id, key, value).Err
		if err != nil {
			fmt.Fprintf(w, "%s\n", err)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusCreated)
		}
	}
}
