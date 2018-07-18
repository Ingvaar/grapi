package main

import (
	"flag"
)

type Options struct {
	RoutesConfig	string
	DBLogin		string
}

func Pars_cmdline() Options {
	var opt Options

	flag.StringVar(&opt.RoutesConfig,
				"config",
				"./config.json",
				"Path to routes config file (json)")
	flag.StringVar(&opt.DBLogin,
				"db",
				"./dbconfig.json",
				"Path to db config file (json)")
	flag.Parse()
	return opt
}
