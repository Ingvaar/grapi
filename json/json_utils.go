package json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// ToMap : convert a json array to a map
func ToMap(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	jsonMap := make(map[string]interface{})
	body, err := ioutil.ReadAll(r.Body)

	if err == nil {
		array := strings.Split(string(body), "\n")
		json.Unmarshal([]byte(array[3]), &jsonMap)
	}
	return (jsonMap)
}

// PrintMapToJSON : print a json array from a map
// to the specified http response writer
func PrintMapToJSON(w http.ResponseWriter, redisMap map[string]interface{}) {
	multInsert := false

	fmt.Fprintf(w, "{")
	for key, value := range redisMap {
		if multInsert {
			fmt.Fprintf(w, ", ")
		}
		fmt.Fprintf(w, "\"%s\":\"%v\"", key, value)
		multInsert = true
	}
	fmt.Fprintf(w, "}")
}
