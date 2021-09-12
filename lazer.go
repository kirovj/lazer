package lazer

import (
	"io"
	"os"
	"sync"
	"time"
)

// Logger is a struct of Lazer.
type Logger struct {
	// Writer is the Writer of logger, the default value is os.Stderr.
	Writer io.Writer
	// Pipe is the pipeline to transfer Msg
	Pipe *Pipe
	// Wait is the wait second between pulls
	Wait int
}

// Msg is the target to log.
type Msg struct {
	Level
	Content string
	Time    int64
}

func NewLogger(writer io.Writer, pipeSize int) *Logger {
	logger := &Logger{
		Writer: writer,
		Pipe:   NewPipe(pipeSize),
	}
	go pull(logger)
	return logger
}

var DefaultLogger *Logger
var once sync.Once

// Default to make a singleton instance of Logger
func Default() *Logger {
	once.Do(func() {
		DefaultLogger = NewLogger(os.Stderr, DefaultPipeSize)
	})
	return DefaultLogger
}

func push(l *Logger, level Level, content string) {
	m := Msg{
		Level:   level,
		Content: content,
		Time:    time.Now().Unix(),
	}
	l.Pipe.In <- format(m)
}

func format(m Msg) []byte {
	var level string
	switch m.Level {
	case INFO:
		level = "INFO"
	case WARN:
		level = "WARN"
	case TRACE:
		level = "TRACE"
	case ERROR:
		level = "ERROR"
	case DEBUG:
		level = "DEBUG"
	}

	return []byte("[" + level + "]: " + m.Content + "\n")
}

func pull(l *Logger) {
	for {
		select {
		case msg := <-l.Pipe.Out:
			if _, err := l.Writer.Write(msg); err != nil {
				_, _ = os.Stderr.Write([]byte("Write log error!"))
			}
		default:
			time.Sleep(time.Second * time.Duration(l.Wait))
		}
	}
}

func (l *Logger) out(msg string) {
	_, _ = l.Writer.Write([]byte(msg))
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
