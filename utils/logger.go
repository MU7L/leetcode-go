package utils

import (
	"fmt"
	"os"
)

type logger struct {
	module string
}

func Logger(module string) *logger {
	return &logger{module}
}

const (
	errorColor   = "\033[31m"
	defaultColor = "\033[0m"
)

func (l *logger) Log(format string, args ...interface{}) {
	l.println(defaultColor, format, args...)
}

// Error 打印错误日志（导出，大写）
func (l *logger) Error(format string, args ...interface{}) {
	l.println(errorColor, format, args...)
}

// Fatal 打印错误日志并退出（导出，大写）
func (l *logger) Fatal(format string, args ...interface{}) {
	l.println(errorColor, format, args...)
	os.Exit(1)
}

// print 内部打印方法（私有，小写）
func (l *logger) println(color string, format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	fmt.Printf("%s[%s] %s%s\n",
		color,
		l.module,
		msg,
		defaultColor)
}
