package logger

import (
	"go.uber.org/zap"
)

func SetupLogger(logLevel string, isProduction bool) *zap.Logger {
	var config zap.Config
	if isProduction {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
	}

	level, err := zap.ParseLevel(logLevel)
	if err != nil {
		level = zap.DebugLevel
	}
	config.Level.SetLevel(level)

	logger, err := config.Build()
	if err != nil {
		panic("failed to initialize logger: " + err.Error())
	}
	return logger
}
