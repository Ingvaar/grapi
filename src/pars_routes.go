package main;

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"fmt"
)

func create_routes(opt Options) Routes {
	_, err := os.Stat(opt.RoutesConfig);

	if err == nil {
		fmt.Printf("Config file %s found !\n", opt.RoutesConfig);
		return(pars_config(opt.RoutesConfig));
	} else {
		fmt.Printf("Config file %s not found\n", opt.RoutesConfig);
		os.Exit(1);
	}
	return (nil);
}

func pars_config(path string) Routes {
	raw, err := ioutil.ReadFile(path);
	var routes Routes;

	if err != nil {
		fmt.Printf("Error while reading config file\n");
		os.Exit(1);
	}
	json.Unmarshal([]byte(raw), &routes);
	return (routes);
}
