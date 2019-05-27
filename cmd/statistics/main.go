package main

import (
	"net/http"
	"testServerStats/pkg/logging"

	log "github.com/sirupsen/logrus"
	config "testServerStats/configs"
	"testServerStats/pkg/server"
)

func main() {
	logging.SetupLogger()
	r := server.GetRouter()
	err := http.ListenAndServe(config.Config.ServerPort, r)
	if err != nil && err != http.ErrServerClosed {
		log.WithError(err).Fatal("Failed to establish database connection")
	}
}
