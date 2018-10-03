package nosql

import (
	"net/http"

	"github.com/gorilla/mux"

	"grapi/db"
	j "grapi/json"
	"grapi/utils"
)

// Set : use the hmset cmd from redis with a json passed as body
// and with {type}:{id} as the id of the entry
func Set(w http.ResponseWriter, r *http.Request) {
	jsonmap := j.ToMap(w, r)
	id := mux.Vars(r)["type"] + ":" + mux.Vars(r)["id"]

	for key, value := range jsonmap {
		err := db.Nosql.Cmd("HMSET", id, key, value).Err
		if err != nil {
			utils.SendResponse(w, err, http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusCreated)
		}
	}
}
