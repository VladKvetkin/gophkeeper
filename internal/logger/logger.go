package logger

import "go.uber.org/zap"

// Logger - структура логгера
type Logger struct {
	Log *zap.SugaredLogger
}

func NewLogger(logLevel string) (*Logger, error) {
	lvl, err := zap.ParseAtomicLevel(logLevel)
	if err != nil {
		return nil, err
	}

	loggerConfig := zap.NewProductionConfig()
	loggerConfig.Level = lvl

	zapLogger, err := loggerConfig.Build()
	if err != nil {
		return nil, err
	}

	return &Logger{Log: zapLogger.Sugar()}, nil
}
