package lazer

type Logger struct {
	IsLazy   bool
	Filepath string
}

func NewLogger(isLazy bool, filepath string) *Logger {
	return &Logger{
		IsLazy:   isLazy,
		Filepath: filepath,
	}
}

func Default() *Logger {
	return NewLogger(true, DefaultPath)
}

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
