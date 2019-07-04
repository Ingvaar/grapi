package config

import (
	"flag"

	"grapi/core"
)

// getFiles : return the function struct filled via fillFiles function
func getFiles() core.Files {
	var files core.Files

	fillFilesStruct(&files)
	return files
}

// fillFiles : Pars the flags from the cmdline and returns an Files struct
func fillFilesStruct(files *core.Files) {
	flag.StringVar(&files.Routes,
		"routes",
		"./routes.json",
		"Path to routes config file (json)")
	flag.StringVar(&files.Config,
		"config",
		"./config.json",
		"Path to api config file (json)")
	flag.Parse()
}
