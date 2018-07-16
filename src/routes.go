package main;

import "net/http"

type Route struct {
	Name		string		`json:"name"`;
	Method		string		`json:"method"`;
	Pattern		string		`json:"pattern"`;
	HandlerFunc	string		`json:"handler"`;
}

var HandlerFunc = map[string]http.HandlerFunc{
	"status": status,
	"index": index,
	"getTable": getTable,
	"getLine": getLine,
	"createLine": createLine,
	"deleteLine": deleteLine,
}

type Routes []Route;
