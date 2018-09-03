package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ingvaar/grapi/db"
)

// getAllRedis : do a hgetall on the id passed in the url an return a json array
func getAllRedis(w http.ResponseWriter, r *http.Request) {
	pathVars := mux.Vars(r)
	id := fmt.Sprintf("%s:%s", pathVars["type"], pathVars["id"])

	reply, err := db.Db.Redis.Cmd("HGETALL", id).Map()
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		printMapToJSON(w, reply)
		w.WriteHeader(http.StatusOK)
	}
}
