package main

import "net/http"

// Route : struct for the routes config file
type Route struct {
	Name        string `json:"name"`
	Method      string `json:"method"`
	Pattern     string `json:"pattern"`
	HandlerFunc string `json:"handler"`
}

// HandlerFunc : map to get the handlers from the str in the struct
var HandlerFunc = map[string]http.HandlerFunc{
	"status":        status,
	"index":         index,
	"getTableSQL":   getTableSQL,
	"getLineSQL":    getLineSQL,
	"createLineSQL": createLineSQL,
	"deleteLineSQL": deleteLineSQL,
	"updateLineSQL": updateLineSQL,
}

// Routes : declares the type of an array of Route
type Routes []Route
