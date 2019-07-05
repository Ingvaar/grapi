package config

import (
	"io"
	"testing"

	"grapi/core"
)

func testParsFileKO(reader io.Reader) func(*testing.T) {
	return func(t *testing.T) {
		var config core.Config

		err := parsFile(reader, &config)
		if err == nil {
			t.Errorf("Error nil")
		}
	}
}

func testParsFileOK(reader io.Reader) func(*testing.T) {
	return func(t *testing.T) {
		var config core.Config

		err := parsFile(reader, &config)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	}
}
