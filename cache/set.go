package redis

import (
	"net/http"

	"github.com/gorilla/mux"

	"grapi/utils"
)

// Set : use the hmset cmd from redis with a json passed as body
// and with {type}:{id} as the id of the entry
func (rd *Redis) Set(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := mux.Vars(r)["type"] + ":" + mux.Vars(r)["id"]

	for key, value := range r.Form {
		err := rd.RC.Cmd("HMSET", id, key, value[0]).Err
		if err != nil {
			utils.SendError(w, err, http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusCreated)
		}
	}
}
