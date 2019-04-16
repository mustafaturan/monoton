// Package mtimer returns the current monotonic time in nanoseconds
package mtimer

import (
	"time"
)

var initialTime time.Time
var monotonicTime uint

func init() {
	initialTime = time.Now()
	monotonicTime = uint(initialTime.UnixNano())
}

// Now returns the current monotonic time in nanoseconds
func Now() uint {
	return monotonicTime + uint(time.Since(initialTime).Nanoseconds())
}
