package handlers

import (
	"net/http"

	"grapi/auth"
	"grapi/nosql"
	"grapi/sql"
)

// HandlerFunc : map to get the handlers from the str in the struct
var HandlerFunc = map[string]http.HandlerFunc{
	"status": status,
	"index":  index,
	"show":   sql.Show,
	"select": sql.Select,
	"insert": sql.Insert,
	"delete": sql.Delete,
	"update": sql.Update,
	"set":    nosql.Set,
	"read":   nosql.Read,
	"remove": nosql.Delete,
	"login":  auth.Login,
}
