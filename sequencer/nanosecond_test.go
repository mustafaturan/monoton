package sequencer

import (
	"reflect"
	"testing"
)

func TestNewNanosecond(t *testing.T) {
	want := reflect.TypeOf(&Nanosecond{})
	got := reflect.TypeOf(NewNanosecond())

	if want != got {
		t.Errorf(
			"NewNanosecond() call want type: %T, got type: %T",
			want,
			got,
		)
	}
}
