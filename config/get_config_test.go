package config

import (
	"io/ioutil"
	"log"
	"strings"
	"testing"

	"grapi/core"
)

func TestFileParsing(t *testing.T) {
	config := new(core.Config)
	log.SetOutput(ioutil.Discard)

	t.Run("parsJsonOK", func(t *testing.T) {
		err := parsFile(strings.NewReader("{\"cache\": 0}"), config)
		if err != nil {
			t.Errorf("JSON parsing error: %v", err)
		}
	})
	t.Run("parsJsonKO", func(t *testing.T) {
		err := parsFile(strings.NewReader("\"che\" 0}"), config)
		if err == nil {
			t.Errorf("JSON error detecting error")
		}
	})
	t.Run("readerKO", func(t *testing.T) {
		err := parsFile(strings.NewReader("{\"cache\": 0}"), config)
		if err == nil {
			t.Errorf("parsFile reader error")
		}
	})
}
