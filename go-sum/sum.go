package sum

import "sync/atomic"

type Adder struct {
	*int64
}

func NewAdder() Adder {
	return Adder{int64: new(int64)}
}

func (a Adder) Add(val int64) {
	atomic.AddInt64(a.int64, val)
}
func (a Adder) Sum() int64 {
	return atomic.LoadInt64(a.int64)
}
