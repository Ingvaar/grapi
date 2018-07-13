package main

import (
	"net/http"
	"log"
)

func main() {
	opt := Pars_cmdline();
	router := NewRouter(opt);

	log.Fatal(http.ListenAndServe(":8080", router));
}
