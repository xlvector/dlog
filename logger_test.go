package dlog

import (
	"testing"
)

func TestLogger(t *testing.T) {
	Warn("hello %s", "world")
	Info("hello %s", "world")
	Error("hello %s", "world")
	Info("hello %s", "world")
}
