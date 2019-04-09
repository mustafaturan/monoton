// Copyright 2019 Mustafa Turan. All rights reserved.
// Use of this source code is governed by a Apache License 2.0 license that can
// be found in the LICENSE file.

package sequencer

import (
	"time"
)

// Second is monotonic second time based sequencer for monoton
type Second struct {
	sequence      uint
	sequenceTime  uint
	monotonicTime time.Time
}

var s *Second

const (
	// Maxiumum sequence time value for Second sequencer
	secondMaxSequenceTime = 62*62*62*62*62*62 - 1
	// Maxiumum sequence value for Second sequencer
	secondMaxSequence = 62*62*62*62*62*62 - 1
	// One second value in nanoseconds
	secondInNanoseconds = int64(time.Second)
)

func init() {
	monotonicTime := time.Now()
	s = &Second{
		sequence:      1,
		sequenceTime:  uint(monotonicTime.Unix()),
		monotonicTime: monotonicTime,
	}
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
	timeDiff, timeDiffInSeconds := s.monotonicTimeDiff()
	if timeDiffInSeconds > 0 {
		s.monotonicTime = s.monotonicTime.Add(time.Duration(timeDiff))
		s.sequenceTime += timeDiffInSeconds
		s.sequence = 0
	} else {
		s.sequence++
	}
}

func (s *Second) monotonicTimeDiff() (uint, uint) {
	timeDiff := time.Since(s.monotonicTime).Nanoseconds()
	return uint(timeDiff), uint(timeDiff / secondInNanoseconds)
}
