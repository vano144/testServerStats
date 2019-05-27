package logging

import (
	"bytes"
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	config "testServerStats/configs"
	"testing"
)

func TestSetupLogger(t *testing.T) {
	config.Config.LogLevel = "debug"
	buf := &bytes.Buffer{}

	SetupLogger()
	log.SetOutput(buf)

	if log.GetLevel() != log.DebugLevel {
		t.Fatalf("Log level is not set and equal %s", log.GetLevel())
	}

	log.WithField("error", errors.New("wild walrus")).Debug("test")
	var data map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &data); err != nil {
		t.Fatal("Error Unmarshal logger json")
	}

}


func TestSetErrorLevel(t *testing.T) {
	config.Config.LogLevel = "testError"

	defer (func() {
		if err := recover(); err == nil {
			t.Fatal("Not panic if logger config with error loglevel as testError")
		}
	})()

	SetupLogger()
}
