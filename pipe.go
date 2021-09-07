package lazer

// Pipe is an infinite chan.
// In is used to write without blocking, which supports multiple writers.
// and Out is used to read, which supports multiple readers.
// You can close the in channel if you want.
type Pipe struct {
	bufCount int64
	In       chan<- Msg // channel for write
	Out      <-chan Msg // channel for read
	buffer   []Msg      // buffer
}

func DefaultPipe() *Pipe {
	return NewPipe()
}

func NewPipe() *Pipe {
	return &Pipe{
		In:     nil,
		Out:    nil,
		buffer: nil,
	}
}
