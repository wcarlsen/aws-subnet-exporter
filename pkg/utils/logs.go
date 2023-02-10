package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

func SetupLogger(debug *bool) {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	if *debug {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
}
