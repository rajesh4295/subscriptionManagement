package util

import "github.com/sirupsen/logrus"

func InitLogger(level string) (*logrus.Logger, error) {
	baseLogger := logrus.New()
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		return nil, err
	}
	baseLogger.SetLevel(logLevel)
	baseLogger.SetFormatter(&logrus.JSONFormatter{})
	return baseLogger, nil
}
