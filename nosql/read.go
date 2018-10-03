package nosql

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"grapi/db"
	j "grapi/json"
	"grapi/utils"
)

// Read : do a hget with a json array passed in the body
// on the id passed in the url an return a json array
func Read(w http.ResponseWriter, r *http.Request) {
	jsonmap := j.ToMap(w, r)
	id := mux.Vars(r)["type"] + ":" + mux.Vars(r)["id"]
	result := "{"
	mult := false

	for key := range jsonmap {
		reply, err := db.Nosql.Cmd("HGET", id, key).Str()
		if err != nil {
			utils.SendResponse(w, err, http.StatusBadRequest)
			return
		}
		if mult {
			result += ", "
		}
		mult = true
		result += "\"" + key + "\":\"" + reply + "\""
	}
	result += "}"
	fmt.Fprintln(w, result)
}

/*
// getAllRedis : do a hgetall on the id passed in the url an return a json array
func getAllRedis(w http.ResponseWriter, r *http.Request) {
	pathVars := mux.Vars(r)
	id := fmt.Sprintf("%s:%s", pathVars["type"], pathVars["id"])

	reply, err := db.Nosql.Cmd("HGETALL", id).Map()
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		j.PrintMapToJSON(w, reply)
		w.WriteHeader(http.StatusOK)
	}
}
*/
