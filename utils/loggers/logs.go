package loggers

import "github.com/sirupsen/logrus"

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	//Setting the formt of logs to JSON
	Logger.SetFormatter(&logrus.JSONFormatter{})

	//Configuring for errors where application needs to get terminated
	Logger.ExitFunc = func(code int) {
		panic("Fatal error encountered")
	}
}
