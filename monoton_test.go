// Copyright 2020 Mustafa Turan. All rights reserved.
// Use of this source code is governed by a Apache License 2.0 license that can
// be found in the LICENSE file.

package monoton

import (
	"errors"
	"math"
	"strings"
	"testing"

	"github.com/mustafaturan/monoton/sequencer"
)

func TestNew(t *testing.T) {
	tests := []struct {
		s           sequencer.Sequencer
		node        uint64
		initialTime uint64
		wantErr     error
		wantNode    string
		wantTime    uint64
	}{
		{
			&validSequencer{},
			3843,
			uint64(1),
			nil,
			"zz",
			uint64(1),
		},
	}

	errorTests := []struct {
		s           sequencer.Sequencer
		node        uint64
		initialTime uint64
		wantErr     error
		wantNode    string
		wantTime    uint64
	}{
		{
			&validSequencer{},
			3844,
			uint64(2),
			errors.New("node can't be greater than 3843 (given 3844)"),
			"",
			uint64(2),
		},
		{
			&invalidSequencer{},
			1,
			uint64(0),
			errors.New("max byte size sum of sequence(8) and time sequence(8) can't be >= total byte size(16)"),
			"",
			uint64(0),
		},
	}

	configureMsg := "New(%v, %d, %d) want: %v, got: %v"
	nodeMsg := "New(%v, %d, _) want node: %s, got node: %s"
	timeMsg := "New(%v, _, %d) want time: %d, got time: %d"
	for _, test := range tests {
		test := test
		got, err := New(test.s, test.node, test.initialTime)

		t.Run("assigns node val correctly", func(t *testing.T) {
			if got.node != test.wantNode {
				t.Errorf(nodeMsg, test.s, test.node, test.wantNode, got.node)
			}
		})

		t.Run("assigns initialTime val correctly", func(t *testing.T) {
			if got.initialTime != test.wantTime {
				t.Errorf(timeMsg, test.s, test.initialTime, test.wantTime, got.initialTime)
			}
		})

		t.Run("must not have error", func(t *testing.T) {
			if err != test.wantErr {
				t.Errorf("want %+v but got %+v", test.wantErr, err)
			}
		})
	}

	for _, test := range errorTests {
		test := test
		got, err := New(test.s, test.node, test.initialTime)
		t.Run("must not initialize", func(t *testing.T) {
			if got != nil {
				t.Errorf("want nil but got %+v", got)
			}
		})
		t.Run("errors with correct message", func(t *testing.T) {
			if err != test.wantErr && err.Error() != test.wantErr.Error() {
				t.Errorf(configureMsg, test.s, test.node, test.initialTime, test.wantErr, got)
			}
		})
	}
}

func TestNext(t *testing.T) {
	m, _ := New(&validSequencer{}, 3843, 0)
	m1, m2 := m.Next(), m.Next()

	t.Run("generates greater sequences on each call", func(t *testing.T) {
		t.Parallel()
		if strings.Compare(m1, m2) >= 0 {
			t.Errorf("Next(): %s >= Next(): %s", m1, m2)
		}
	})

	t.Run("generates 16 bytes string sequences", func(t *testing.T) {
		t.Parallel()
		results := []string{m1, m2}
		for _, r := range results {
			if len(r) != 16 {
				t.Errorf("Next(): %s couldn't produce 16 bytes string", r)
			}
		}
	})
}

type validSequencer struct {
	counter uint64
}

type invalidSequencer struct {
	counter uint64
}

func (v *validSequencer) MaxTime() uint64 {
	return uint64(math.Pow(62, 8)) - 1
}

func (v *validSequencer) Max() uint64 {
	return uint64(math.Pow(62, 6)) - 1
}

func (v *validSequencer) Next() (uint64, uint64) {
	v.counter++
	return 1, v.counter
}

func (i *invalidSequencer) MaxTime() uint64 {
	return uint64(math.Pow(62, 8)) - 1
}

func (i *invalidSequencer) Max() uint64 {
	return uint64(math.Pow(62, 8)) - 1
}

func (i *invalidSequencer) Next() (uint64, uint64) {
	return 1, i.counter
}
