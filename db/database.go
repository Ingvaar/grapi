package db

import (
	"database/sql"

	"github.com/mediocregopher/radix.v2/redis"
)

// SQL : SQL database
var SQL *sql.DB

// Nosql : nosql database
var Nosql *redis.Client
