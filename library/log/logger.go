package log

import "go.uber.org/zap"

var Logger *zap.Logger
var sugar *zap.SugaredLogger

func init() {
	Logger = zap.NewNop()
	sugar = Logger.Sugar()
}

func SetLogger(name string, level string) *zap.Logger {
	config := zap.NewProductionConfig()
	_ = config.Level.UnmarshalText([]byte(level))

	Logger, _ = config.Build()
	Logger = Logger.Named(name)

	sugar = Logger.Sugar()

	return Logger
}

func Info(args ...interface{}) {
	sugar.Info(args...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	sugar.Infow(msg, keysAndValues...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	sugar.Warnw(msg, keysAndValues...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	sugar.Errorw(msg, keysAndValues)
}
