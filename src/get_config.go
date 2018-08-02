package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Config : struct of the config file
type Config struct {
	Redis         int    `json:"redis"`
	RedisAddress  string `json:"redis_address"`
	RedisPassword string `json:"redis_password"`
	SQL           int    `json:"sql"`
	SQLAddress    string `json:"sql_address"`
	SQLUsername   string `json:"sql_username"`
	SQLPassword   string `json:"sql_password"`
	SQLDatabase   string `json:"sql_database"`
}

// GetConfig : returns the config struct from config file
// path in Options struct
func GetConfig(opt Options) Config {
	var cfg Config
	_, err := os.Stat(opt.ConfigFile)

	if err == nil {
		cfg = parsCfg(opt.ConfigFile)
	} else {
		os.Exit(1)
	}
	return (cfg)
}

func parsCfg(path string) Config {
	raw, err := ioutil.ReadFile(path)
	var cfg Config

	if err != nil {
		log.Fatal("Error while reading config file\n")
		os.Exit(1)
	}
	json.Unmarshal([]byte(raw), &cfg)
	return (cfg)
}
