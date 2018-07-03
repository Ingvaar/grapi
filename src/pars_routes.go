package main;

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"fmt"
)

func create_routes() Routes {
	args := os.Args;

	if len(args) > 1 {
		_, err := os.Stat(args[1]);
		if err == nil {
			fmt.Printf("Config file %s found !\n", args[1]);
			return(pars_config(args[1]));
		} else {
			fmt.Printf("Config file %s not found\n", args[1]);
			os.Exit(1);
		}
	} else {
		_, err := os.Stat("./config.json");
		if err == nil {
			fmt.Println("Config file found");
			return(pars_config("./config.json"));
		} else {
			fmt.Printf("Config file not found\n");
			os.Exit(1);
		}
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
