package common

import "go.uber.org/zap"

type zapLogger struct{ logger *zap.Logger }

type Logger interface {
	Info(msg string)
	Error(msg string)
}

func (l *zapLogger) Info(msg string) {
	l.logger.Info(msg)

}

func (l *zapLogger) Error(msg string) {
	l.logger.Error(msg)
}

func NewLogger() Logger {
	logger, err := zap.NewProduction()

	if err != nil {
		panic(err)
	}

	return &zapLogger{logger}
}
