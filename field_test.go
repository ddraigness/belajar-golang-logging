package belajargolanglogging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.WithField("username", "surya").Info("Hello Logging")

	logger.WithField("username", "yasin").WithField("name", "Yasheeennn").Info("Hello Logging")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "surya",
		"name": "Surya Paloh",
	}).Info("Hello Logging")
}