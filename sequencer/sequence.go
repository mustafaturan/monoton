// Copyright 2020 Mustafa Turan. All rights reserved.
// Use of this source code is governed by a Apache License 2.0 license that can
// be found in the LICENSE file.

package sequencer

import "sync"

// Sequence is an implementation of sequencer
type Sequence struct {
	sync.Mutex

	current uint64
	time    uint64
	max     uint64
	maxTime uint64
	now     func() uint64
}

// Max returns the maximum possible sequence value
func (s *Sequence) Max() uint64 {
	return s.max
}

// MaxTime returns the maximum possible time sequence value
func (s *Sequence) MaxTime() uint64 {
	return s.maxTime
}

// Next returns the next sequence
func (s *Sequence) Next() (uint64, uint64) {
	s.Lock()
	defer s.Unlock()

	s.increment()
	return s.time, s.current
}

func (s *Sequence) increment() {
	now := s.now()
	if s.time < now {
		s.time = now
		s.current = 0
	} else {
		s.current++
	}
}
