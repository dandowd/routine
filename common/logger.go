package common

type Logger struct{}

func (l *Logger) Info(args ...interface{}) {
}

func (l *Logger) Error(args ...interface{}) {
}

func NewLogger() *Logger {
	return &Logger{}
}
