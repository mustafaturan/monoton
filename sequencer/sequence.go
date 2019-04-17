package sequencer

import "sync"

// Sequence is an implementation of sequencer
type Sequence struct {
	sync.Mutex

	current uint
	time    uint
	max     uint
	maxTime uint
	now     func() uint
}

// Max returns the maximum possible sequence value
func (s *Sequence) Max() uint {
	return s.max
}

// MaxTime returns the maximum possible time sequence value
func (s *Sequence) MaxTime() uint {
	return s.maxTime
}

// Next returns the next sequence
func (s *Sequence) Next() (uint, uint) {
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
