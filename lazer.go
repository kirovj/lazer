package lazer

import (
	"io"
	"os"
	"sync"
	"time"
)

type Flush func(msg string)

// Logger is a struct of Lazer.
type Logger struct {
	// Lazy is the flag of logger to define sync or async.
	Lazy bool
	// W is the Writer of logger, the default value is os.Stderr.
	W io.Writer
	// Ch is the channel to save log
	Ch    chan string
	ChLen int
	// FlushTime the time between flush
	FlushTime time.Duration
	FlushFunc Flush
}

func NewLogger(Lazy bool, writer io.Writer, chLen int, flushTime int64, flushFunc Flush) *Logger {
	l := &Logger{
		Lazy:      Lazy,
		W:         writer,
		Ch:        make(chan string, chLen),
		ChLen:     chLen,
		FlushTime: time.Duration(flushTime),
		FlushFunc: flushFunc,
	}
	go StartFlush(l)
	return l
}

var DefaultLogger *Logger
var once sync.Once

// Default to make a singleton instance of Logger
func Default() *Logger {
	once.Do(func() {
		DefaultLogger = &Logger{
			Lazy:      true,
			W:         os.Stderr,
			Ch:        make(chan string, 10),
			ChLen:     100,
			FlushTime: 2,
		}
		DefaultLogger.FlushFunc = func(msg string) {
			_, _ = DefaultLogger.W.Write([]byte(msg))
		}
		go StartFlush(DefaultLogger)
	})
	return DefaultLogger
}

// join makes one log with level, time, msg
func join(level, msg string) string {
	return "[" + level + "] " + msg + "\n"
}

func push(l *Logger, level, msg string) {
	s := join(level, msg)
	select {
	case l.Ch <- s:
		return
	default:
		l.FlushFunc(s)
	}
}

func (l *Logger) out(msg string) {
	_, _ = l.W.Write([]byte(msg))
}

func (l *Logger) Info(msg string) {
	go push(l, INFO, msg)
}

func (l *Logger) Warn(msg string) {
	go push(l, WARN, msg)
}

func (l *Logger) Trace(msg string) {
	go push(l, TRACE, msg)
}

func (l *Logger) Debug(msg string) {
	go push(l, DEBUG, msg)
}

func (l *Logger) Error(msg string) {
	go push(l, ERROR, msg)
}

func StartFlush(l *Logger) {
	for {
		select {
		case msg := <-l.Ch:
			l.FlushFunc(msg)
		default:
			time.Sleep(l.FlushTime * time.Second)
		}
	}
}
