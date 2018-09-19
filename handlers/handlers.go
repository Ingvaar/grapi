package handlers

import (
	"net/http"

	"grapi/sql"
	"grapi/nosql"
)

// HandlerFunc : map to get the handlers from the str in the struct
var HandlerFunc = map[string]http.HandlerFunc{
	"status":	status,
	"index":	index,
	"show":		sql.Show,
	"select":	sql.Select,
	"insert":	sql.Insert,
	"delete":	sql.Delete,
	"update":	sql.Update,
	"create":	nosql.Create,
	"read":		nosql.Read,
	"updatenosql":	nosql.Update,
	"deletenosql":	nosql.Delete,
}
