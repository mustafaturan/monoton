// Copyright 2019 Mustafa Turan. All rights reserved.
// Use of this source code is governed by a Apache License 2.0 license that can
// be found in the LICENSE file.

package sequencer

import (
	"time"

	"github.com/mustafaturan/monoton/mtimer"
)

// Second is monotonic second time based sequencer for monoton
type Second struct {
	sequence     uint
	sequenceTime uint
}

var s *Second

const (
	// Maxiumum sequence time value for Second sequencer
	secondMaxSequenceTime = 62*62*62*62*62*62 - 1
	// Maxiumum sequence value for Second sequencer
	secondMaxSequence = 62*62*62*62*62*62 - 1
	// One second in nanoseconds
	secondInNanoseconds = uint(time.Second)
)

func init() {
	s = &Second{sequence: 0}
	s.sequenceTime = s.now()
}

// NewSecond returns initialized small genarator
func NewSecond() *Second {
	return s
}

// MaxSequenceTime returns the maximum possible time sequence value
func (s *Second) MaxSequenceTime() uint {
	return secondMaxSequenceTime
}

// MaxSequence returns the maximum possible sequence value for a given time
func (s *Second) MaxSequence() uint {
	return secondMaxSequence
}

// Next increments the time and related counter sequences
func (s *Second) Next() (uint, uint) {
	s.incrementSequences()
	return s.sequenceTime, s.sequence
}

func (s *Second) incrementSequences() {
	currentTime := s.now()
	if currentTime > s.sequenceTime {
		s.sequenceTime = currentTime
		s.sequence = 0
	} else {
		s.sequence++
	}
}

func (s *Second) now() uint {
	return mtimer.Now() / secondInNanoseconds
}
