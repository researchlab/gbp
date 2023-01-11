package logs

import (
	"os"

	"github.com/sirupsen/logrus"
)

var serviceLog = logrus.New()

func Init(logFile *os.File) {
	serviceLog.SetLevel(logrus.DebugLevel)
	serviceLog.Formatter = &logrus.JSONFormatter{}
	serviceLog.SetOutput(logFile)
}

func Logger() *logrus.Logger {
	return serviceLog
}

func Entry() *logrus.Entry {
	return serviceLog.WithField("service", "scope")
}
