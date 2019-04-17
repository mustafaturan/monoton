package sequencer

import (
	"testing"
	"time"

	"github.com/mustafaturan/monoton/mtimer"
)

func TestMax_Sequence(t *testing.T) {
	want := uint(12345)
	s := &Sequence{max: want}

	if got := s.Max(); got != want {
		t.Errorf("Max() want: %d, got: %d", want, got)
	}
}

func TestMaxTime_Sequence(t *testing.T) {
	want := uint(1<<64 - 1)
	s := &Sequence{maxTime: want}

	if got := s.MaxTime(); got != want {
		t.Errorf("MaxTime() want: %d, got: %d", want, got)
	}
}

func TestNext_Sequence(t *testing.T) {
	tests := []struct {
		want uint
		now  func() uint
	}{
		{want: uint(2), now: func() uint { return 0 }},
		{want: uint(0), now: func() uint { time.Sleep(time.Nanosecond); return mtimer.Now() }},
	}

	for _, test := range tests {
		t.Run("resets counter correctly when time changes", func(t *testing.T) {
			t.Parallel()
			s := &Sequence{now: test.now}
			s.Next()
			s.Next()
			if _, got := s.Next(); got != test.want {
				t.Errorf("Next().current want: %d, got: %d", test.want, got)
			}
		})
	}
}
