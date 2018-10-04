package redis

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	j "grapi/json"
	"grapi/utils"
)

// Read : do a hget with a json array passed in the body
// on the id passed in the url an return a json array
func (db *Database) Read(w http.ResponseWriter, r *http.Request) {
	jsonmap := j.ToMap(w, r)
	id := mux.Vars(r)["type"] + ":" + mux.Vars(r)["id"]
	result := "{"
	mult := false

	for key := range jsonmap {
		reply, err := db.DB.Cmd("HGET", id, key).Str()
		if err != nil {
			utils.SendError(w, err, http.StatusBadRequest)
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
