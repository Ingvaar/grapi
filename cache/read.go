package redis

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"grapi/utils"
)

// Read : do a hget with a json array passed in the body
// on the id passed in the url an return a json array
func (rd *Redis) Read(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := mux.Vars(r)["type"] + ":" + mux.Vars(r)["id"]
	result := "{"
	mult := false

	for key := range r.Form {
		reply, err := rd.RC.Cmd("HGET", id, key).Str()
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
