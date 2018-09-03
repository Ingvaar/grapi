package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Config : struct of the config file
type Config struct {
	Options

	Redis	      int    `json:"redis"`
	RedisConnType string `json:"redis_conn_type"`
	RedisAddress  string `json:"redis_address"`
	RedisPort     string `json:"redis_port"`
	RedisPassword string `json:"redis_password"`
	UseSQL	      int    `json:"sql"`
	AddressSQL    string `json:"sql_address"`
	UsernameSQL   string `json:"sql_username"`
	PasswordSQL   string `json:"sql_password"`
	DatabaseSQL   string `json:"sql_database"`
	ServerPort    string `json:"server_port"`
}

// Cfg : global var of Config struct
var Cfg Config

// GetConfig : returns the config struct from config file
// path in Options struct
func GetConfig() {
	_, err := os.Stat(Cfg.Options.ConfigFile)

	if err == nil {
		raw, err := ioutil.ReadFile(Cfg.Options.ConfigFile)

		if err != nil {
			log.Fatal("Error while reading config file\n")
		}
		json.Unmarshal([]byte(raw), &Cfg)
	} else {
		log.Fatal(err)
	}
}
