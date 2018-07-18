package main

import (
	"database/sql"
	"log"
	"os"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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
	var ping error

	db, err = sql.Open("mysql", connectionStr)
	ping = db.Ping()
	if err != nil && ping != nil {
		defer db.Close()
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println("Database connected")
	return (db)
}
