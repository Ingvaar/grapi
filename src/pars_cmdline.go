package main

import (
	"flag"
)

// Options : struct with path to config files as str
type Options struct {
	RoutesFile string
	ConfigFile string
}

// ParsCmdline : Pars the flags from the cmdline and returns an Options struct
func ParsCmdline() Options {
	var opt Options

	flag.StringVar(&opt.RoutesFile,
		"routes",
		"./routes.json",
		"Path to routes config file (json)")
	flag.StringVar(&opt.ConfigFile,
		"config",
		"./config.json",
		"Path to api config file (json)")
	flag.Parse()
	return opt
}
