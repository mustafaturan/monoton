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
		s           sequencer.Sequencer
		node        uint
		initialTime uint
		wantErr     error
		wantNode    string
		wantTime    uint
	}{
		{
			&validSequencer{},
			3843,
			uint(1),
			nil,
			"zz",
			uint(1),
		},
		{
			&validSequencer{},
			3844,
			uint(2),
			errors.New("Node can't be greater than 3843 (given 3844)"),
			"",
			uint(2),
		},
		{
			&invalidSequencer{},
			1,
			uint(0),
			errors.New("Sum of s:8, t:8 bytes can't be >= total byte size"),
			"",
			uint(0),
		},
	}

	configureMsg := "Configure(%v, %d, %d) want: %v, got: %v"
	nodeMsg := "Configure(%v, %d, _) want node: %s, got node: %s"
	timeMsg := "Configure(%v, _, %d) want time: %d, got time: %d"
	for _, test := range tests {
		got := Configure(test.s, test.node, test.initialTime)

		t.Run("assigns node val correctly", func(t *testing.T) {
			if c.node != test.wantNode {
				t.Errorf(nodeMsg, test.s, test.node, test.wantNode, c.node)
			}
		})

		t.Run("assigns initialTime val correctly", func(t *testing.T) {
			if c.initialTime != test.wantTime {
				t.Errorf(timeMsg, test.s, test.initialTime, test.wantTime, c.initialTime)
			}
		})

		t.Run("errors with correct message", func(t *testing.T) {
			if got != test.wantErr && got.Error() != test.wantErr.Error() {
				t.Errorf(configureMsg, test.s, test.node, test.initialTime, test.wantErr, got)
			}
		})
	}
}

func TestNext(t *testing.T) {
	Configure(&validSequencer{}, 3843, 0)
	m1, m2 := Next(), Next()

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
