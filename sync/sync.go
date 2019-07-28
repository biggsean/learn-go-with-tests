package main

import (
	"sync"
)

// Counter is a counter
type Counter struct {
	mu    sync.Mutex
	value int
}

// Inc increments the counter
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.value++
}

// Value returns the value of the counter
func (c *Counter) Value() int {
	return c.value
}

// NewCounter constructs a Counter
func NewCounter() *Counter {
	return &Counter{}
}
