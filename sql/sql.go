package sql

import (
	"database/sql"
	"fmt"
	"log"

	// Importing MySQL driver
	_ "github.com/go-sql-driver/mysql"

	"grapi/core"
)

// Database :
type Database struct {
	DB *sql.DB
	config core.Config
}

// Connect : connect the database
func (db *Database) Connect(config *core.Config) {
	db.config = *config
	connectionStr := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		db.config.UsernameSQL,
		db.config.PasswordSQL,
		db.config.AddressSQL,
		db.config.DatabaseSQL)

	if db.config.UseSQL == 0 {
		db.DB = nil
		return
	}
	var err error
	db.DB, err = sql.Open("mysql", connectionStr)
	ping := db.DB.Ping()
	if err != nil || ping != nil {
		defer db.DB.Close()
		log.Fatal("Cannot connect to SQL database")
	}
	log.Printf("SQL Database connected with address %s\n", db.config.AddressSQL)
}

// Register : register the functions to handler map
func (db *Database) Register(handlers *core.Handlers) {
	sql := Database{DB: db.DB, config: db.config}
	temp := core.Handlers{}

	for key, value := range *handlers {
		temp[key] = value
	}
	temp["show"] = sql.Show
	temp["select"] = sql.Select
	temp["insert"] = sql.Insert
	temp["delete"] = sql.Delete
	temp["update"] = sql.Update
	temp["login"] = sql.Login
	*handlers = temp
}
