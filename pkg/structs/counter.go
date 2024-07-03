package structs

import "sync/atomic"

type Counter struct {
	count int32
}

func NewCounter() *Counter {
	return &Counter{
		count: 0,
	}
}

func (c *Counter) Increment() int32 {
	return atomic.AddInt32(&c.count, 1)
}

func (c *Counter) Decrement() int32 {
	return atomic.AddInt32(&c.count, -1)
}

func (c *Counter) Get() int32 {
	return atomic.LoadInt32(&c.count)
}
