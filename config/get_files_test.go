package config

import (
	"testing"

	"grapi/core"
)

func testFillFilesStructOK(expected core.Files) func(*testing.T) {
	return func(t *testing.T) {
		var files core.Files

		fillFilesStruct(&files)
		if files != expected {
			t.Errorf("Got %v but expected %v", files, expected)
		}
	}
}
