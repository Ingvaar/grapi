package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// config : struct of the config file
type config struct {
	Options

	NoSQL	      int    `json:"nosql"`
	NoSQLConnType string `json:"nosql_conn_type"`
	NoSQLAddress  string `json:"nosql_address"`
	NoSQLPort     string `json:"nosql_port"`
	NoSQLPassword string `json:"nosql_password"`
	UseSQL	      int    `json:"sql"`
	AddressSQL    string `json:"sql_address"`
	UsernameSQL   string `json:"sql_username"`
	PasswordSQL   string `json:"sql_password"`
	DatabaseSQL   string `json:"sql_database"`
	ServerPort    string `json:"server_port"`
}

// Cfg : global var of Config struct
var Cfg config

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
