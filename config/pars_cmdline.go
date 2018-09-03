package config

import (
	"flag"
)

// Options : struct with path to config files as str
type Options struct {
	RoutesFile string
	ConfigFile string
}

// ParsCmdline : Pars the flags from the cmdline and returns an Options struct
func ParsCmdline() {

	flag.StringVar(&Cfg.Options.RoutesFile,
		"routes",
		"./routes.json",
		"Path to routes config file (json)")
	flag.StringVar(&Cfg.Options.ConfigFile,
		"config",
		"./config.json",
		"Path to api config file (json)")
	flag.Parse()
}
