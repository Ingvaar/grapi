package main

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"strings"
)

func json_to_map(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	jsonMap := make(map[string]interface{})
	body, err := ioutil.ReadAll(r.Body)

	if err == nil {
		array := strings.Split(string(body), "\n")
		json.Unmarshal([]byte(array[3]), &jsonMap)
	}
	return (jsonMap)
}
