package config

import (
	"testing"
)

func TestParsCmdlineNoArgs(t *testing.T) {
	opts := parsCmdline()
	routes := "./routes.json"
	config := "./config.json"

	if (opts.RoutesFile != routes) {
		t.Errorf("got %s want %s", opts.RoutesFile, routes)
	}
	if (opts.ConfigFile != config) {
		t.Errorf("got %s want %s", opts.ConfigFile, config)
	}
}
