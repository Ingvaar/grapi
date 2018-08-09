package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// jsonToMap : convert a json array to a map
func jsonToMap(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	jsonMap := make(map[string]interface{})
	body, err := ioutil.ReadAll(r.Body)

	if err == nil {
		array := strings.Split(string(body), "\n")
		json.Unmarshal([]byte(array[3]), &jsonMap)
	}
	return (jsonMap)
}

// printMapToJson : print a json array from a map
// to the specified http response writer
func printMapToJson(w http.ResponseWriter, redisMap map[string]string) {
	multInsert := false

	fmt.Fprintf(w, "{")
	for key, value := range redisMap {
		if multInsert {
			fmt.Fprintf(w, ", ")
		}
		fmt.Fprintf(w, "\"%s\":\"%s\"", key, value)
		multInsert = true
	}
	fmt.Fprintf(w, "}")
}
