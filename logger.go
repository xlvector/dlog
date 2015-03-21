package dlog

import (
	"fmt"
	"github.com/wsxiaoys/terminal/color"
	"io"
	"os"
)

const (
	FATAL = 0
	PANIC = 1
	ERROR = 2
	WARN  = 3
	INFO  = 4
	DEBUG = 5
)

var Level int64 = INFO

var lg *Logger

func init() {
	lg = New(os.Stdout, "")
}

func Info(format string, v ...interface{}) {
	if Level >= INFO {
		lg.Output(2, fmt.Sprintf("[INFO] "+format, v...))
	}
}

func Warn(format string, v ...interface{}) {
	if Level >= WARN {
		escapeCode := color.Colorize("y")
		io.WriteString(os.Stdout, escapeCode)
		line := color.Sprintf("[WARN] "+format, v...)
		lg.Output(2, line)
	}
}

func Error(format string, v ...interface{}) {
	if Level >= ERROR {
		escapeCode := color.Colorize("r")
		io.WriteString(os.Stdout, escapeCode)
		line := color.Sprintf("[ERROR] "+format, v...)
		lg.Output(2, line)
	}
}

func ErrorN(n int, format string, v ...interface{}) {
	if Level >= ERROR {
		lg.Output(2+n, fmt.Sprintf("[ERROR] "+format, v...))
	}
}

func Debug(format string, v ...interface{}) {
	if Level >= DEBUG {
		lg.Output(2, fmt.Sprintf("[DEBUG] "+format, v...))
	}
}

func Fatal(format string, v ...interface{}) {
	if Level >= FATAL {
		lg.Output(2, fmt.Sprintf("[FATAL] "+format, v...))
		os.Exit(1)
	}
}

func Panic(format string, v ...interface{}) {
	if Level >= PANIC {
		s := fmt.Sprintf("[PANIC] "+format, v...)
		lg.Output(2, s)
		panic(s)
	}
}
