package config

import (
	"flag"
)

// Options : struct with path to config files as str
type Options struct {
	RoutesFile string
	ConfigFile string
}

// parsCmdline : Pars the flags from the cmdline and returns an Options struct
func parsCmdline() Options {
	options := *new(Options)

	flag.StringVar(&options.RoutesFile,
		"routes",
		"./routes.json",
		"Path to routes config file (json)")
	flag.StringVar(&options.ConfigFile,
		"config",
		"./config.json",
		"Path to api config file (json)")
	flag.Parse()
	return (options)
}
