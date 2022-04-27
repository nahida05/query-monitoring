package logger

import "github.com/sirupsen/logrus"

func New() *logrus.Logger {
	return logrus.New()
}
