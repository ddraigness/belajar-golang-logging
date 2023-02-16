package belajargolanglogging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSingleton(t *testing.T) {
	logrus.Info("Hello Info")
	logrus.Error("Hello Error")

	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Info("Info Hello")
	logrus.Error("Error Hello")
}
