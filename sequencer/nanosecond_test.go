package sequencer

import (
	"reflect"
	"testing"
)

func TestNewNanosecond(t *testing.T) {
	expectedType := reflect.TypeOf(&Nanosecond{})
	gotType := reflect.TypeOf(NewNanosecond())

	if expectedType != gotType {
		t.Errorf(
			"NewNanosecond() call expected: %T, resulted with %T",
			expectedType,
			gotType,
		)
	}
}
