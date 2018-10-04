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
	_, err := os.Stat(config.RoutesFile)

	if err == nil {
		return (parsRoutes(config.RoutesFile))
	}
	os.Exit(1)
	return (nil)
}

func parsRoutes(path string) []core.Route {
	raw, err := ioutil.ReadFile(path)
	var routes []core.Route

	if err != nil {
		log.Fatal("Error while reading routes file\n")
		os.Exit(1)
	}
	json.Unmarshal([]byte(raw), &routes)
	return (routes)
}
