package logger

import (
	"fmt"
	"os"
)

type Logger struct {
	module string
}

func (l *Logger) Log(format string, args ...any) {
	l.println(DefaultColor, format, args...)
}

// Error 打印错误日志（导出，大写）
func (l *Logger) Error(format string, args ...any) {
	l.println(ErrorColor, format, args...)
}

// Fatal 打印错误日志并退出（导出，大写）
func (l *Logger) Fatal(format string, args ...any) {
	l.println(ErrorColor, format, args...)
	os.Exit(1)
}

// print 内部打印方法（私有，小写）
func (l *Logger) println(color string, format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	fmt.Printf("%s[%s] %s%s\n",
		color,
		l.module,
		msg,
		DefaultColor)
}
