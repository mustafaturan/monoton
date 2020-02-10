// Copyright 2020 Mustafa Turan. All rights reserved.
// Use of this source code is governed by a Apache License 2.0 license that can
// be found in the LICENSE file.

package sequencer

import (
	"testing"
	"time"

	"github.com/mustafaturan/monoton/mtimer"
)

func TestMax_Sequence(t *testing.T) {
	want := uint64(12345)
	s := &Sequence{max: want}

	if got := s.Max(); got != want {
		t.Errorf("Max() want: %d, got: %d", want, got)
	}
}

func TestMaxTime_Sequence(t *testing.T) {
	want := uint64(1<<64 - 1)
	s := &Sequence{maxTime: want}

	if got := s.MaxTime(); got != want {
		t.Errorf("MaxTime() want: %d, got: %d", want, got)
	}
}

func TestNext_Sequence(t *testing.T) {
	timer := mtimer.New()
	tests := []struct {
		want uint64
		now  func() uint64
	}{
		{want: uint64(2), now: func() uint64 { return 0 }},
		{want: uint64(0), now: func() uint64 { time.Sleep(time.Nanosecond); return timer.Now() }},
	}

	for _, test := range tests {
		test := test
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
