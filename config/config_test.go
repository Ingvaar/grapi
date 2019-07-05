package config

import (
	"io/ioutil"
	"log"
	"strings"
	"testing"

	"grapi/core"
)

func TestConfig(t *testing.T) {
	log.SetOutput(ioutil.Discard)

	t.Run("parseFile: {\"cache\":0}",
		testParsFileOK(strings.NewReader("{\"cache\":0}")))
	t.Run("parseFile: 1234",
		testParsFileKO(strings.NewReader("1234")))
	t.Run("parseFile: ReaderError",
		testParsFileKO(strings.NewReader("")))
	t.Run("fillFileStruct",
		testFillFilesStructOK(core.Files{Routes: "./routes.json", Config: "./config.json"}))
}
