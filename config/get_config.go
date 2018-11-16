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
func GetConfig() *core.Config {
	config := new(core.Config)
	options := parsCmdline()
	config.RoutesFile = options.RoutesFile
	_, err := os.Stat(options.ConfigFile)

	if err == nil {
		raw, err := ioutil.ReadFile(options.ConfigFile)
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
	return (config)
}
