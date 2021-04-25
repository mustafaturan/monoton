// Copyright 2021 Mustafa Turan. All rights reserved.
// Use of this source code is governed by a Apache License 2.0 license that can
// be found in the LICENSE file.

package sequencer

import (
	"sync/atomic"
)

// Sequence is an implementation of sequencer
type Sequence struct {
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
	now := s.now()
	time := atomic.LoadUint64(&s.time)
	var current uint64
	if time < now {
		time = now
		atomic.StoreUint64(&s.time, time)
		atomic.StoreUint64(&s.current, 0)
	} else {
		current = atomic.AddUint64(&s.current, 1)
	}
	return time, current
}
