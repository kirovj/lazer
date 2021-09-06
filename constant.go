package lazer

// State msg type
type Level uint8

const (
	INFO Level = iota
	WARN
	ERROR
	TRACE
	DEBUG
)

const (
	DefaultPath = "logs"
)
