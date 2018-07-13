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
				"routes config")
	flag.StringVar(&opt.DBLogin,
				"db",
				"./dbconfig.json",
				"db config")
	flag.Parse()
	return opt
}
