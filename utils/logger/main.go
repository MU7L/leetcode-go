package logger

func NewLogger(module string) *Logger {
	return &Logger{module}
}
