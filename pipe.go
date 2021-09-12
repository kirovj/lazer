package lazer

import (
	"sync/atomic"
)

// T defines interface{}, and will be used for generic type after go 1.18 is released.
type T interface{}

// Pipe is an infinite chan.
// In is used to write without blocking, which supports multiple writers.
// and Out is used to read, which supports multiple readers.
// You can close the in channel if you want.
type Pipe struct {
	bufCount int64
	In       chan<- T // channel for write
	Out      <-chan T // channel for read
	buffer   *Ring    // buffer
}

// Len returns len of In plus len of Out plus len of buffer.
// It is not accurate and only for your evaluating approximate number of elements in this chan,
// see https://github.com/smallnest/chanx/issues/7.
func (c Pipe) Len() int {
	return len(c.In) + c.BufLen() + len(c.Out)
}

// BufLen returns len of the buffer.
// It is not accurate and only for your evaluating approximate number of elements in this chan,
// see https://github.com/smallnest/chanx/issues/7.
func (c Pipe) BufLen() int {
	return int(atomic.LoadInt64(&c.bufCount))
}

// NewPipe creates the unbounded chan.
// in is used to write without blocking, which supports multiple writers.
// and out is used to read, which supports multiple readers.
// You can close the in channel if you want.
func NewPipe(initCapacity int) *Pipe {
	return NewPipeWithSize(initCapacity, initCapacity, initCapacity)
}

// NewPipeWithSize is like NewPipe, but you can set initial capacity for In, Out, Buffer.
func NewPipeWithSize(initInCapacity, initOutCapacity, initBufCapacity int) *Pipe {
	in := make(chan T, initInCapacity)
	out := make(chan T, initOutCapacity)
	ch := Pipe{In: in, Out: out, buffer: NewRingBuffer(initBufCapacity)}

	go process(in, out, &ch)

	return &ch
}

func process(in, out chan T, ch *Pipe) {
	defer close(out)
loop:
	for {
		val, ok := <-in
		if !ok { // in is closed
			break loop
		}

		// make sure values' order
		// buffer has some values
		if atomic.LoadInt64(&ch.bufCount) > 0 {
			ch.buffer.Write(val)
			atomic.AddInt64(&ch.bufCount, 1)
		} else {
			// out is not full
			select {
			case out <- val:
				continue
			default:
			}

			// out is full
			ch.buffer.Write(val)
			atomic.AddInt64(&ch.bufCount, 1)
		}

		for !ch.buffer.IsEmpty() {
			select {
			case val, ok := <-in:
				if !ok { // in is closed
					break loop
				}
				ch.buffer.Write(val)
				atomic.AddInt64(&ch.bufCount, 1)

			case out <- ch.buffer.Peek():
				ch.buffer.Pop()
				atomic.AddInt64(&ch.bufCount, -1)
				if ch.buffer.IsEmpty() && ch.buffer.size > ch.buffer.initialSize { // after burst
					ch.buffer.Reset()
					atomic.StoreInt64(&ch.bufCount, 0)
				}
			}
		}
	}

	// drain
	for !ch.buffer.IsEmpty() {
		out <- ch.buffer.Pop()
		atomic.AddInt64(&ch.bufCount, -1)
	}

	ch.buffer.Reset()
	atomic.StoreInt64(&ch.bufCount, 0)
}
