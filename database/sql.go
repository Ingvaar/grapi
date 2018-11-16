package sql

import (
	"database/sql"
	"fmt"
	"log"

	// Importing MySQL driver
	_ "github.com/go-sql-driver/mysql"

	"grapi/core"
)

// SQL :
type SQL struct {
	DB *sql.DB
	config core.Config
}

// Connect : connect the database
func (db *SQL) Connect(config *core.Config) {
	db.config = *config
	connectionStr := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		db.config.DatabaseUsername,
		db.config.DatabasePassword,
		db.config.DatabaseAddress,
		db.config.DatabaseName)

	if db.config.Database == 0 {
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
	log.Printf("SQL Database connected with address %s\n", db.config.DatabaseAddress)
}

// Register : register the functions to handler map
func (db *SQL) Register(handlers *core.Handlers) {
	sql := SQL{DB: db.DB, config: db.config}
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
