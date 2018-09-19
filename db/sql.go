package db

import (
	"database/sql"
	"fmt"
	"log"

	// Importing MySQL driver
	_ "github.com/go-sql-driver/mysql"

	c "grapi/config"
)

// OpenSQL : open a connection to the SQL database
func OpenSQL() {
	connectionStr := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		c.Cfg.UsernameSQL,
		c.Cfg.PasswordSQL,
		c.Cfg.AddressSQL,
		c.Cfg.DatabaseSQL)
	var err error
	var ping error

	if c.Cfg.UseSQL == 0 {
		SQL = nil
		return
	}
	SQL, err = sql.Open("mysql", connectionStr)
	ping = SQL.Ping()
	if err != nil || ping != nil {
		defer SQL.Close()
		log.Fatal("Cannot connect to SQL database")
	}
	log.Printf("SQL Database connected with address %s\n", c.Cfg.AddressSQL)
}
