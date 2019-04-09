// Copyright 2019 Mustafa Turan. All rights reserved.
// Use of this source code is governed by a Apache License 2.0 license that can
// be found in the LICENSE file.

package sequencer

import (
	"time"
)

// Nanosecond is monotonic nanosecond time based sequencer for monoton
type Nanosecond struct {
	sequence      uint
	sequenceTime  uint
	monotonicTime time.Time
}

var n *Nanosecond

const (
	// Maxiumum sequence time value for Nanosecond sequencer
	nanosecondMaxSequenceTime = 1<<64 - 1
	// Maxiumum sequence value for Nanosecond sequencer
	nanosecondMaxSequence = 62*62 - 1
)

func init() {
	monotonicTime := time.Now()
	n = &Nanosecond{
		sequence:      0,
		sequenceTime:  uint(monotonicTime.UnixNano()),
		monotonicTime: monotonicTime,
	}
}

// NewNanosecond returns the initialized nanosecond sequencer
func NewNanosecond() *Nanosecond {
	return n
}

// MaxSequenceTime returns the maximum possible time sequence value
func (n *Nanosecond) MaxSequenceTime() uint {
	return nanosecondMaxSequenceTime
}

// MaxSequence returns the maximum possible sequence value for a given time
func (n *Nanosecond) MaxSequence() uint {
	return nanosecondMaxSequence
}

// Next increments the time and related counter sequences
func (n *Nanosecond) Next() (uint, uint) {
	n.incrementSequences()
	return n.sequenceTime, n.sequence
}

func (n *Nanosecond) incrementSequences() {
	timeDiff := n.monotonicTimeDiff()
	if timeDiff > 0 {
		n.monotonicTime = n.monotonicTime.Add(time.Duration(timeDiff))
		n.sequenceTime += timeDiff
		n.sequence = 0
	} else {
		n.sequence++
	}
}

func (n *Nanosecond) monotonicTimeDiff() uint {
	timeDiff := time.Since(n.monotonicTime).Nanoseconds()
	return uint(timeDiff)
}
