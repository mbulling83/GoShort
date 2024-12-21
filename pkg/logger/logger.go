package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

// Init initializes the logger
func Init() {
	log = logrus.New()
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

// Info logs an informational message
func Info(message string, fields map[string]interface{}) {
	if fields != nil {
		log.WithFields(fields).Info(message)
	} else {
		log.Info(message)
	}
}

// Error logs an error message
func Error(message string, fields map[string]interface{}) {
	if fields != nil {
		log.WithFields(fields).Error(message)
	} else {
		log.Error(message)
	}
}
