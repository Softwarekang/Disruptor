package disruptor

import "strconv"

const (
	BufferFullErr  = Err(0x1)
	BufferEmptyErr = Err(0x2)
)

type Err uintptr

func (e Err) Error() string {
	if 0 <= int(e) && int(e) < len(errors) {
		s := errors[e]
		if s != "" {
			return s
		}
	}
	return "errno " + strconv.Itoa(int(e))
}

// Error table
var errors = [...]string{
	1: "disruptor buffer is full, please try again later",
	2: "disruptor buffer is empty, please try again later",
}
