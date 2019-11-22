// Package mtimer returns the current monotonic time in nanoseconds
package mtimer

import (
	"time"
)

var initialTime time.Time
var monotonicTime uint64

func init() {
	initialTime = time.Now()
	monotonicTime = uint64(initialTime.UnixNano())
}

// Now returns the current monotonic time in nanoseconds
func Now() uint64 {
	return monotonicTime + uint64(time.Since(initialTime).Nanoseconds())
}
