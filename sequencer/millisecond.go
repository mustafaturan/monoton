// Copyright 2019 Mustafa Turan. All rights reserved.
// Use of this source code is governed by a Apache License 2.0 license that can
// be found in the LICENSE file.

package sequencer

import (
	"time"
)

// Millisecond is monotonic millisecond time based sequencer for monoton
type Millisecond struct {
	sequence      uint
	sequenceTime  uint
	monotonicTime time.Time
}

var m *Millisecond

const (
	// Maxiumum sequence time value for Millisecond sequencer
	millisecondMaxSequenceTime = 62*62*62*62*62*62*62*62 - 1
	// Maxiumum sequence value for Millisecond sequencer
	millisecondMaxSequence = 62*62*62*62 - 1
	// One millisecond value in nanoseconds
	millisecondInNanoseconds = int64(time.Millisecond)
)

func init() {
	monotonicTime := time.Now()
	m = &Millisecond{
		sequence:      0,
		sequenceTime:  uint(monotonicTime.UnixNano() / millisecondInNanoseconds),
		monotonicTime: monotonicTime,
	}
}

// NewMillisecond returns the initialized millisecond sequencer
func NewMillisecond() *Millisecond {
	return m
}

// MaxSequenceTime returns the maximum possible time sequence value
func (m *Millisecond) MaxSequenceTime() uint {
	return millisecondMaxSequenceTime
}

// MaxSequence returns the maximum possible sequence value for a given time
func (m *Millisecond) MaxSequence() uint {
	return millisecondMaxSequence
}

// Next increments the time and related counter sequences
func (m *Millisecond) Next() (uint, uint) {
	m.incrementSequences()
	return m.sequenceTime, m.sequence
}

func (m *Millisecond) incrementSequences() {
	timeDiff, timeDiffInMilliseconds := m.monotonicTimeDiff()
	if timeDiffInMilliseconds > 0 {
		m.monotonicTime = m.monotonicTime.Add(time.Duration(timeDiff))
		m.sequenceTime += timeDiffInMilliseconds
		m.sequence = 0
	} else {
		m.sequence++
	}
}

func (m *Millisecond) monotonicTimeDiff() (uint, uint) {
	timeDiff := time.Since(m.monotonicTime).Nanoseconds()
	return uint(timeDiff), uint(timeDiff / millisecondInNanoseconds)
}
