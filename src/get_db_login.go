package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func get_login(opt Options) db_login {
	var db db_login
	_, err := os.Stat(opt.DBLogin)

	if err == nil {
		db = pars_db(opt.DBLogin)
	} else {
		os.Exit(1)
	}
	return (db)
}

func pars_db(path string) db_login {
	raw, err := ioutil.ReadFile(path)
	var db db_login;

	if err != nil {
		log.Fatal("Error while reading db config file\n")
		os.Exit(1)
	}
	json.Unmarshal([]byte(raw), &db)
	return (db);
}
