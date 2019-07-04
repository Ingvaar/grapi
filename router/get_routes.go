package router

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"grapi/core"
)

// CreateRoutes :
func CreateRoutes(config *core.Config) []core.Route {
	handle, err := os.Open(config.Files.Routes)

	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()
	routes, err := parsRoutes(config.Files.Routes)
	if err != nil {
		log.Fatal(err)
	}
	return (routes)
}

func parsRoutes(path string) ([]core.Route, error) {
	raw, err := ioutil.ReadFile(path)
	var routes []core.Route

	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(raw), &routes)
	if err != nil {
		return nil, err
	}
	return routes, nil
}
