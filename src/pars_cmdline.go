package main

import (
	"flag"
)

type Options struct {
	RoutesFile	string
	ConfigFile	string
}

func Pars_cmdline() Options {
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
