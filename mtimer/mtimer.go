// Package mtimer returns the current monotonic time in nanoseconds
package mtimer

import (
	"time"
)

// Timer is a monotonic time ticker
type Timer struct {
	initialTime   time.Time
	monotonicTime uint64
}

// New inits the timer using current system time values
func New() Timer {
	initialTime := time.Now()
	monotonicTime := uint64(initialTime.UnixNano())
	return Timer{initialTime: initialTime, monotonicTime: monotonicTime}
}

// Now returns the current monotonic time in nanoseconds
func (t Timer) Now() uint64 {
	return t.monotonicTime + uint64(time.Since(t.initialTime).Nanoseconds())
}
