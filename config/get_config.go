package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"grapi/core"
)

// GetConfig : returns the config struct from config file
// path in Options struct
func GetConfig(config *core.Config) {
	options := parsCmdline()
	if (config.RoutesFile == "") {
		config.RoutesFile = options.RoutesFile
	}
	if (config.ConfigFile == "") {
		config.ConfigFile = options.ConfigFile
	}
	_, err := os.Stat(config.ConfigFile)

	if err == nil {
		raw, err := ioutil.ReadFile(config.ConfigFile)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal([]byte(raw), &config)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
}
