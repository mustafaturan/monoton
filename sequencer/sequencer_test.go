package sequencer

import (
	"math"
	"reflect"
	"testing"
	"time"
)

func TestMaxTime(t *testing.T) {
	tests := []struct {
		wantMaxTime uint64
		seq         Sequencer
	}{
		{uint64(math.Pow(62, 6) - 1), NewSecond()},
		{uint64(math.Pow(62, 8) - 1), NewMillisecond()},
		{uint64(1<<64 - 1), NewNanosecond()},
	}

	for _, test := range tests {
		gotMaxTime := test.seq.MaxTime()

		if test.wantMaxTime != gotMaxTime {
			t.Errorf(
				"%s.MaxTime(), want: %d, got: %d",
				reflect.TypeOf(test.seq).String(),
				test.wantMaxTime,
				gotMaxTime,
			)
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		wantMax uint64
		seq     Sequencer
	}{
		{uint64(math.Pow(62, 6) - 1), NewSecond()},
		{uint64(math.Pow(62, 4) - 1), NewMillisecond()},
		{uint64(math.Pow(62, 2) - 1), NewNanosecond()},
	}

	for _, test := range tests {
		gotMax := test.seq.Max()

		if test.wantMax != gotMax {
			t.Errorf(
				"%s.Max(), want: %d, got: %d",
				reflect.TypeOf(test.seq).String(),
				test.wantMax,
				gotMax,
			)
		}
	}
}

func TestNext(t *testing.T) {
	sameMomentTests := []struct {
		sequencer Sequencer
	}{
		{NewSecond()},
		{NewMillisecond()},
		// NOTE: For Nanosecond type it is possible to use the known time
		// freezing techniques to test
	}

	for _, test := range sameMomentTests {
		sequencer := test.sequencer

		t.Run("at the same time", func(t *testing.T) {
			sequenceTimeVals := [2]uint64{1, 2}
			sequenceVals := [2]uint64{0, 0}

			// Loops until we have the same time for two calls
			for sequenceTimeVals[0] != sequenceTimeVals[1] {
				sequenceTimeVals[0], sequenceVals[0] = sequencer.Next()
				sequenceTimeVals[1], sequenceVals[1] = sequencer.Next()
			}

			t.Run("sequence val should increment", func(t *testing.T) {
				t.Parallel()
				if sequenceVals[1] <= sequenceVals[0] {
					t.Errorf(
						"%s.Next() should increment the sequence, %v",
						reflect.TypeOf(sequencer).String(),
						sequenceVals,
					)
				}
			})
		})
	}

	futureTimeTests := []struct {
		sequencer     Sequencer
		sleepDuration time.Duration
	}{
		{NewSecond(), time.Second},
		{NewMillisecond(), time.Millisecond},
		{NewNanosecond(), time.Nanosecond},
	}

	for _, test := range futureTimeTests {
		sequencer := test.sequencer
		t.Run("in a future time", func(t *testing.T) {
			sequenceTimeVals := [2]uint64{0, 0}

			sequenceTimeVals[0], _ = sequencer.Next()
			time.Sleep(test.sleepDuration)
			sequenceTimeVals[1], _ = sequencer.Next()

			t.Run("time sequence val should increment", func(t *testing.T) {
				t.Parallel()
				if sequenceTimeVals[1] <= sequenceTimeVals[0] {
					t.Errorf(
						"%s.Next() should increment the time seq, %v",
						reflect.TypeOf(sequencer).String(),
						sequenceTimeVals,
					)
				}
			})
		})
	}
}
