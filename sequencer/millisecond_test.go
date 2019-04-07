package sequencer

import (
	"reflect"
	"testing"
)

func TestNewMillisecond(t *testing.T) {
	expectedType := reflect.TypeOf(&Millisecond{})
	gotType := reflect.TypeOf(NewMillisecond())

	if expectedType != gotType {
		t.Errorf(
			"NewMillisecond() call expected: %T, resulted with %T",
			expectedType,
			gotType,
		)
	}
}
