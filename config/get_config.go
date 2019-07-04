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
	config.Files = getFiles()
	handle, err := os.Open(config.Files.Config)

	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()
	err = parsFile(handle, config)
	if err != nil {
		log.Fatal(err)
	}
}

func parsFile(file io.Reader, config *core.Config) error {
	raw, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(raw), &config)
	if err != nil {
		return err
	}
	return nil
}
