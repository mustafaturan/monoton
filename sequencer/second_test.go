package sequencer

import (
	"reflect"
	"testing"
)

func TestNewSecond(t *testing.T) {
	expectedType := reflect.TypeOf(&Second{})
	gotType := reflect.TypeOf(NewSecond())

	if expectedType != gotType {
		t.Errorf(
			"NewSecond() call expected: %T, resulted with %T",
			expectedType,
			gotType,
		)
	}
}
