package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SetupLogger(logLevel string, isProduction bool) *zap.Logger {

	// Zap has two logging configs out the box, but you can create custom ones too

	// There is a lot more you can do with this logger, adding it into the middleware to get request context logs added automatically.

	var config zap.Config
	if isProduction {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
	}

	// Convert string to zap level
	var level zapcore.Level
	if err := level.UnmarshalText([]byte(logLevel)); err != nil {
		// Handle invalid level string
		level = zapcore.InfoLevel
	}
	config.Level.SetLevel(level)

	logger, err := config.Build()
	if err != nil {
		panic("failed to initialize logger: " + err.Error())
	}
	return logger
}
