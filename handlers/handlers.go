package handlers

import (
	"net/http"
)

// HandlerFunc : map to get the handlers from the str in the struct
var HandlerFunc = map[string]http.HandlerFunc{
	"status":	 status,
	"index":	 index,
	"getTableSQL":	 getTableSQL,
	"getLineSQL":	 getLineSQL,
	"createLineSQL": createLineSQL,
	"deleteLineSQL": deleteLineSQL,
	"updateLineSQL": updateLineSQL,
	"setEntryRedis": setEntryRedis,
	"getAllRedis":	 getAllRedis,
	"getRedis":	 getRedis,
}
