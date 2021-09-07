package lazer

import (
	"io"
	"os"
	"sync"
)

// Logger is a struct of Lazer.
type Logger struct {
	// Writer is the Writer of logger, the default value is os.Stderr.
	Writer io.Writer
	// Pipe is the pipeline to transfer Msg
	Pipe *Pipe
}

// Msg is the target to log.
type Msg struct {
	Level
	Content []byte
	Time    int64
}

func NewLogger(writer io.Writer) *Logger {
	logger := &Logger{
		Writer: writer,
		Pipe:   DefaultPipe(),
	}
	go pull(logger)
	return logger
}

var DefaultLogger *Logger
var once sync.Once

// Default to make a singleton instance of Logger
func Default() *Logger {
	once.Do(func() {
		DefaultLogger = NewLogger(os.Stderr)
	})
	return DefaultLogger
}

func push(logger *Logger, level Level, msg string) {
	//m := &Msg{
	//	Level:   level,
	//	Content: []byte(msg),
	//	Time:    time.Now().Unix(),
	//}

	//select {
	//case logger.Pipline <- m:
	//	return
	//default:
	//}
}

func pull(logger *Logger) {

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
