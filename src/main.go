package main

import (
	"net/http"
	"log"
)

func main() {
	opt := Pars_cmdline();
//	dblogin := get_login(opt);
	router := NewRouter(opt);

	log.Fatal(http.ListenAndServe(":8080", router));
}
