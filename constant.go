package lazer

// Level is to define the type of log
type Level uint8

const (
	INFO Level = iota
	WARN
	ERROR
	TRACE
	DEBUG
)

const (
	DefaultPath     = "logs"
	DefaultPipeSize = 3
)
