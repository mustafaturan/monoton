// Copyright 2019 Mustafa Turan. All rights reserved.
// Use of this source code is governed by a Apache License 2.0 license that can
// be found in the LICENSE file.

package sequencer

import (
	"time"

	"github.com/mustafaturan/monoton/mtimer"
)

// Millisecond is monotonic millisecond time based sequencer for monoton
type Millisecond struct {
	sequence     uint
	sequenceTime uint
}

var m *Millisecond

const (
	// Maxiumum sequence time value for Millisecond sequencer
	millisecondMaxSequenceTime = 62*62*62*62*62*62*62*62 - 1
	// Maxiumum sequence value for Millisecond sequencer
	millisecondMaxSequence = 62*62*62*62 - 1
	// One millisecond in nanoseconds
	millisecondInNanoseconds = uint(time.Millisecond)
)

func init() {
	m = &Millisecond{sequence: 0}
	m.sequenceTime = m.now()
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
	currentTime := m.now()
	if currentTime > m.sequenceTime {
		m.sequenceTime = currentTime
		m.sequence = 0
	} else {
		m.sequence++
	}
}

func (m *Millisecond) now() uint {
	return mtimer.Now() / millisecondInNanoseconds
}
