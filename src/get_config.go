package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Redis		int		`json:"redis"`
	Redis_Adress	string		`json:"redis_adress"`
	Redis_Password	string		`json:"redis_password"`
	SQL		int		`json:"sql"`
	SQL_Adress	string		`json:"sql_address"`
	SQL_Username	string		`json:"sql_username"`
	SQL_Password	string		`json:"sql_password"`
	SQL_Database	string		`json:"sql_database"`
}

func Get_config(opt Options) Config {
	var cfg Config
	_, err := os.Stat(opt.ConfigFile)

	if err == nil {
		cfg = pars_cfg(opt.ConfigFile)
	} else {
		os.Exit(1)
	}
	return (cfg)
}

func pars_cfg(path string) Config {
	raw, err := ioutil.ReadFile(path)
	var cfg Config;

	if err != nil {
		log.Fatal("Error while reading config file\n")
		os.Exit(1)
	}
	json.Unmarshal([]byte(raw), &cfg)
	return (cfg);
}
