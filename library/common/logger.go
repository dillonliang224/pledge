package common

import "go.uber.org/zap"

func NewLogger(serviceName string, level string) *zap.Logger {
	config := zap.NewProductionConfig()
	config.Level.UnmarshalText([]byte(level))
	logger, _ := config.Build()
	logger = logger.Named(serviceName)
	return logger
}
