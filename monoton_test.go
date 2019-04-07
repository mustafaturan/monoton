package monoton

import (
	"errors"
	"math"
	"strings"
	"testing"

	"github.com/mustafaturan/monoton/sequencer"
)

func TestConfigure(t *testing.T) {
	tests := []struct {
		s            sequencer.Sequencer
		node         uint
		expectedErr  error
		expectedNode string
	}{
		{
			&validSequencer{},
			3843,
			nil,
			"zz",
		},
		{
			&validSequencer{},
			3844,
			errors.New("Node can't be greater than 3843 (given 3844)"),
			"",
		},
		{
			&invalidSequencer{},
			1,
			errors.New("Total byte size can't be >= to sum of s:8, t:8"),
			"",
		},
	}

	configureMsg := "Configure(%v, %d) expected: %v, resulted with: %v"
	nodeMsg := "Configure(%v, %d) expected node: %s, resulted with: %s"
	for _, test := range tests {
		result := Configure(test.s, test.node)

		t.Run("assigns node val correctly", func(t *testing.T) {
			if c.node != test.expectedNode {
				t.Errorf(nodeMsg, test.s, test.node, test.expectedNode, c.node)
			}
		})

		t.Run("errors with correct message", func(t *testing.T) {
			if result != test.expectedErr && result.Error() != test.expectedErr.Error() {
				t.Errorf(configureMsg, test.s, test.node, test.expectedErr, result)
			}
		})
	}
}

func TestNext(t *testing.T) {
	Configure(&validSequencer{}, 3843)
	m1, m2 := Next(), Next()

	t.Run("generates bigger sequences on each call", func(t *testing.T) {
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
	counter uint
}

type invalidSequencer struct {
	counter uint
}

func (v *validSequencer) MaxSequenceTime() uint {
	return uint(math.Pow(62, 8)) - 1
}

func (v *validSequencer) MaxSequence() uint {
	return uint(math.Pow(62, 6)) - 1
}

func (v *validSequencer) Next() (uint, uint) {
	v.counter++
	return 1, v.counter
}

func (i *invalidSequencer) MaxSequenceTime() uint {
	return uint(math.Pow(62, 8)) - 1
}

func (i *invalidSequencer) MaxSequence() uint {
	return uint(math.Pow(62, 8)) - 1
}

func (i *invalidSequencer) Next() (uint, uint) {
	return 1, i.counter
}
