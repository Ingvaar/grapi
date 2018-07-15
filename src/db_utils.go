package main

import (
	"database/sql"
	"log"
	"fmt"
)

type db_login struct {
	Adress		string		`json:"address"`
	Username	string		`json:"username"`
	Password	string		`json:"password"`
	Database	string		`json:"database"`
}

func openDatabase(opt Options) *sql.DB {
	dblogin := get_login(opt)
	connectionStr := fmt.Sprintf("%s:%s@%s/%s",
			dblogin.Username,
			dblogin.Password,
			dblogin.Adress,
			dblogin.Database)
	var db *sql.DB
	var err error

	db, err = sql.Open("mysql", connectionStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
