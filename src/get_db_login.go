package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"os"
)

func get_login(opt Options) db_login {
	var db db_login
	_, err := os.Stat(opt.DBLogin)

	if err == nil {
		fmt.Printf("Config file %s found !\n", opt.DBLogin)
		db = pars_db(opt.DBLogin)
	} else {
		fmt.Printf("Config file %s not found\n", opt.DBLogin)
		os.Exit(1)
	}
	return (db)
}

func pars_db(path string) db_login {
	raw, err := ioutil.ReadFile(path)
	var db db_login;

	if err != nil {
		fmt.Printf("Error while reading config file\n")
		os.Exit(1)
	}
	json.Unmarshal([]byte(raw), &db)
	return (db);
}
