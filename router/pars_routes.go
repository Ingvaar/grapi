package router

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"grapi/config"
)

func createRoutes() Routes {
	_, err := os.Stat(config.Cfg.Options.RoutesFile)

	if err == nil {
		return (parsRoutes(config.Cfg.Options.RoutesFile))
	}
	os.Exit(1)
	return (nil)
}

func parsRoutes(path string) Routes {
	raw, err := ioutil.ReadFile(path)
	var routes Routes

	if err != nil {
		log.Fatal("Error while reading routes file\n")
		os.Exit(1)
	}
	json.Unmarshal([]byte(raw), &routes)
	return (routes)
}
