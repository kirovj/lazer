package lazer

import (
	"io"
	"os"
)

var DefaultLogger = NewLogger(true, os.Stderr)

// Logger is a struct of Lazer.
type Logger struct {
	// Lazy is the flag of logger to define sync or async.
	Lazy bool
	// W is the Writer of logger, the default value is os.Stderr.
	W  io.Writer
	Ch chan string
}

func NewLogger(Lazy bool, writer io.Writer) *Logger {
	return &Logger{
		Lazy: Lazy,
		W:    writer,
		Ch:   make(chan string),
	}
}

func Default() *Logger {
	return DefaultLogger
}

// func (l *Logger) out(msg string) {
// 	var buf []byte
// 	l.W.Write(msg)
// }

func (l *Logger) Info(msg string) {

}

func (l *Logger) Warn(msg string) {

}

func (l *Logger) Trace(msg string) {

}

func (l *Logger) Debug(msg string) {

}

func (l *Logger) Error(msg string) {

}
