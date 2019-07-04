package config

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"

	"grapi/core"
)

// GetConfig : returns the config struct from config file
// path in Options struct
func GetConfig(config *core.Config) {
	options := parsCmdline()
	if config.RoutesFile == "" {
		config.RoutesFile = options.RoutesFile
	}
	if config.ConfigFile == "" {
		config.ConfigFile = options.ConfigFile
	}
	handle, err := os.Open(config.ConfigFile)

	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()
	parsFile(handle, config)
}

func parsFile(file io.Reader, config *core.Config) {
	raw, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal([]byte(raw), &config)
	if err != nil {
		log.Fatal(err)
	}
}
