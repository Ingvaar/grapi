package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func openSQLDatabase(cfg Config) *sql.DB {
	connectionStr := fmt.Sprintf("%s:%s@%s/%s",
		cfg.SQLUsername,
		cfg.SQLPassword,
		cfg.SQLAddress,
		cfg.SQLDatabase)
	var db *sql.DB
	var err error
	var ping error

	if cfg.SQL == 0 {
		return (nil)
	}
	db, err = sql.Open("mysql", connectionStr)
	ping = db.Ping()
	if err != nil || ping != nil {
		defer db.Close()
		log.Fatal("Cannot connect to sql database")
		os.Exit(1)
	}
	fmt.Println("SQL Database connected")
	return (db)
}
