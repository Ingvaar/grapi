package db

import (
	"database/sql"

	"github.com/mediocregopher/radix.v2/redis"
)

// Database : struct containing the SQL and Redis database
type Database struct {
	SQL   *sql.DB
	Redis *redis.Client
}

// Db : global var of Database struct
var Db Database
