package dlog

import (
	"fmt"
	"io"
	"runtime"
	"sync"
	"time"
)

type Logger struct {
	mu     sync.Mutex
	prefix string
	out    io.Writer
}

func New(out io.Writer, prefix string) *Logger {
	return &Logger{out: out, prefix: prefix}
}

func time33(s string) int64 {
	var ret int64
	for _, c := range []byte(s) {
		ret *= 33
		ret += int64(c)
	}
	if ret > 0 {
		return ret
	}
	return -1 * ret
}

func (l *Logger) header(tm time.Time, file string, line int, s string) string {
	lid := 100000 + time33(s)%10000
	return fmt.Sprintf("%s %d %s file %s line %d ", tm.Format("2006-01-02 15:04:05"), lid, l.prefix, file, line)
}

func (l *Logger) Output(calldepth int, s string) error {
	now := time.Now() // get this early.
	var file string
	var line int
	l.mu.Lock()
	defer l.mu.Unlock()
	var ok bool
	_, file, line, ok = runtime.Caller(calldepth)
	if !ok {
		file = "???"
		line = 0
	}

	head := l.header(now, file, line, s)
	buf := make([]byte, 0, len(head))
	buf = append(buf, head...)
	for _, c := range []byte(s) {
		if c != '\n' {
			buf = append(buf, c)
		} else {
			buf = append(buf, '\n')
			_, err := l.out.Write(buf)
			if err != nil {
				return err
			}
			buf = buf[:0]
			buf = append(buf, head...)
		}
	}
	if len(buf) > len(head) {
		buf = append(buf, '\n')
		_, err := l.out.Write(buf)
		if err != nil {
			return err
		}
	}
	return nil
}
