package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// hmsetRedis : use the hmset cmd from redis with a json passed as body
// and with {type}:{id} as the id of the entry
func hmsetRedis(w http.ResponseWriter, r *http.Request) {
	jsonmap := jsonToMap(w, r)
	pathVars := mux.Vars(r)
	id := fmt.Sprintf("%s:%s", pathVars["type"], pathVars["id"])
	var err_str string

	for key, value := range jsonmap {
		err := redisCli.Cmd("HMSET", id, key, value).Err
		err_str = fmt.Sprintf("%s\n%s", err_str, err)
	}
	if err_str != "" {
		fmt.Fprintf(w, "%s\n", err_str)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}
