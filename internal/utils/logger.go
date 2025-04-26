package utils

import (
	"log"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

var instance *logrus.Logger
var once sync.Once

func NewLogger(logLevel string) *logrus.Logger {

	once.Do(func() {
		instance = createLogger(logLevel)
	})
	return instance

}

func GetLogger() *logrus.Logger {
	if instance == nil {
		log.Fatal("Logger not initialized. Call NewLogger first.")
	}
	return instance
}

func createLogger(logLevel string) *logrus.Logger {
	logger := logrus.New()

	// Set the log level based on the provided logLevel string
	level, err := logrus.ParseLevel(strings.ToLower(logLevel))
	if err != nil {
		log.Fatalf("Invalid log level: %v. Falling back to 'info'", err)
		level = logrus.InfoLevel
	}
	logger.SetLevel(level)

	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	return logger
}
