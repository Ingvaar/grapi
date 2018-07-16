package main;

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"log"
)

func create_routes(opt Options) Routes {
	_, err := os.Stat(opt.RoutesConfig);

	if err == nil {
		return(pars_config(opt.RoutesConfig));
	} else {
		os.Exit(1);
	}
	return (nil);
}

func pars_config(path string) Routes {
	raw, err := ioutil.ReadFile(path);
	var routes Routes;

	if err != nil {
		log.Fatal("Error while reading routes config file\n");
		os.Exit(1);
	}
	json.Unmarshal([]byte(raw), &routes);
	return (routes);
}
