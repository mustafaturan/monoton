package sequencer

import (
	"reflect"
	"testing"
)

func TestNewSecond(t *testing.T) {
	want := reflect.TypeOf(&Sequence{})
	got := reflect.TypeOf(NewSecond())

	if want != got {
		t.Errorf("NewSecond() call want type: %T, got type: %T", want, got)
	}
}
