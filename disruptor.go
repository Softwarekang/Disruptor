// Package disruptor implements a lock free RingBuffer.
package disruptor

import (
	"sync/atomic"
)

// Disruptor is a lock-free data structure for high performance concurrent programming.
type Disruptor struct {
	w     int64
	r     int64
	cap   int64
	event []Event
}

// Event is the data structure stored in the ring buffer.
type Event struct {
	data []byte
}

// NewDisruptor returns a new Disruptor.
func NewDisruptor(cap int64) *Disruptor {
	return &Disruptor{
		w:     0,
		r:     0,
		cap:   cap,
		event: make([]Event, cap),
	}
}

// Put an event into disruptor.
// Put is safe for concurrent use by multiple goroutines.
func (d *Disruptor) Put(e Event) error {
	if d.full() {
		return BufferFullErr
	}
	for {
		w := atomic.LoadInt64(&d.w)
		nextSeq := w + 1
		if !atomic.CompareAndSwapInt64(&d.w, w, nextSeq) {
			continue
		}

		d.event[w] = e
		break
	}

	return nil
}

// Get an event from disruptor.
// Get is safe for concurrent use by multiple goroutines.
func (d *Disruptor) Get() (Event, error) {
	return Event{}, nil
}

func (d *Disruptor) full() bool {
	return true
}
