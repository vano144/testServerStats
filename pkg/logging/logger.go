package logging

import (
	"github.com/sirupsen/logrus"
	"os"
	config "testServerStats/configs"
)

func SetupLogger(){
	logLevel, err := logrus.ParseLevel(config.Config.LogLevel)
	if err != nil {
		panic(err)
	}
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetReportCaller(true)
}

